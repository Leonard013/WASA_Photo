

<script>

export default {
	data: function() {
		return {
			user: JSON.parse(sessionStorage.getItem('User')),
			stream_over: false,
			errormsg: null,
			loading: false,
			photos: [],
			stream: [],
			comment_text: {},
		}
	},
	methods: {

		async commentPhoto(id) {
			if (this.comment_text[id].length > 0 && this.comment_text[id].length <= 300) {
				this.loading = true;
				this.errormsg = null;
				const formData = new FormData();
				formData.append('text', this.comment_text[id]);
				formData.append('authorId', this.user.userId);
				formData.append('photoId', id);

				try {
					let response = await this.$axios.post("/comments/", formData, {
							headers: {
								'Authorization': this.user.userId,
							}
						}
					);
					this.loading = false;
					console.log("Comment added")
					this.comment_text[id] = null
					window.location.reload();

				} catch (e) {
					this.errormsg = e.toString();
					console.log(this.errormsg)
				}
			} else {
				this.errormsg = "The length of the comment must be between 1 and 300 characters"
			}
		},

		async uncommentPhoto(photo_id,comment_id) {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.delete("/comments/"+comment_id, {
						headers: {
							'Authorization': this.user.userId,
							'userId': this.user.userId,
							'photoId': photo_id
						}
					}
				);
				this.loading = false;
				console.log("Comment deleted")
				window.location.reload();

			} catch (e) {
				this.errormsg = e.toString();
				console.log(this.errormsg)
			}
		},

		async likePhoto(id) {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.post("/likes/", {
					photoId: id,
					userId: this.user.userId
				}, {
						headers: {
							'Authorization': this.user.userId,
						}
					}
				);
				this.loading = false;
				console.log("Photo liked")
				window.location.reload();

			} catch (e) {
				this.errormsg = e.toString();
				console.log(this.errormsg)
			}
		},

		async unlikePhoto(id) {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.delete("/likes/"+id, {
						headers: {
							'Authorization': this.user.userId,
							'userId': this.user.userId,
						}
					}
				);
				this.loading = false;
				console.log("Photo unliked")
				window.location.reload();

			} catch (e) {
				this.errormsg = e.toString();
				console.log(this.errormsg)
			}
		},
		
		async getMyStream() {
			this.loading = true;
			this.errormsg = null;

			try {
				let response = await this.$axios.get("/streams/" + this.user.userId, {
						headers: {
							'Authorization': this.user.userId,
						}
					}
				);
				this.loading = false;
				this.photos = response.data.reverse();
				for (let i = 0; i < this.photos.length; i++) {
					this.photos[i].File = 'data:image/*;base64,' + this.photos[i].File
				}
				console.log(this.photos)
				this.refreshStream();
			} catch (e) {
				this.errormsg = e.toString();
				console.log(this.errormsg)
			} 
		},

		async refreshStream() {

			if (this.photos != null) {
				if (this.photos.length > 20) { // take the first 20 elements and remove them from the slice
					var slice = this.photos.slice(0, 20);
					this.photos = this.photos.slice(20, this.photos.length);
					this.stream = this.stream.concat(slice);
					console.log(this.stream)
				} else {
					this.stream = this.stream.concat(this.photos);
					this.photos = null;
					this.stream_over = true;
				}
			}
		},

	},
	mounted() {
		console.log("mounted_Stream", new Date().toLocaleTimeString());
		this.getMyStream();
		
	}
}
</script>

<template>
	<div v-cloak>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Stream</h1>
		</div>
		
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		
		<div v-if="stream" v-for="photo in stream" :key="photo.photoId" class="photo-entry">
			<tr>
				<td>
					<h3>{{ photo.title }}</h3>
					<img :src="photo.File" alt="Photo" class="photo-container">
					<p>Author: {{ photo.username }}</p>
					<p>Date: {{ new Date(photo.date).toLocaleDateString() }}</p>
				</td>
				<td>
					<tr>
						<td>
							<div v-if="photo.likeAuthors ? !photo.likeAuthors.includes(user.userId) : true">
								<button type="submit" class="btn btn-sm btn-outline-primary" @click="likePhoto(photo.photoId)">Like</button>
							</div>
							<div v-if="photo.likeAuthors ? photo.likeAuthors.includes(user.userId) : false">
								<button type="submit" class="btn btn-sm btn-outline-primary" @click="unlikePhoto(photo.photoId)">Unlike</button>
							</div>
						</td>
						<td>
							<p>Likes: {{ photo.likeIds ? photo.likeIds.length : 0 }}</p>
						</td>
					</tr>
					<tr>
						<td>
							<button type="submit" class="btn btn-sm btn-outline-primary" @click="commentPhoto(photo.photoId)">Comment</button>
						</td>
						<td>
							<input v-model="comment_text[photo.photoId]" placeholder="Enter comment Max 300 char">
						</td>
					</tr>
					<tr>
						<td colspan="2">
							<div class="comments-box">
								<tr v-for="(comment, key) in photo.commentTexts" :key="comment.commentId">
									<td>
										<p>
											{{photo.commentAuthors[key]}}:  {{ comment }} 
										</p>
									</td>
									<td>
										<div v-if="photo.commentAuthors[key] == user.username">
											<button type="submit" class="btn btn-sm btn-outline-primary" @click="uncommentPhoto(photo.photoId, photo.commentIds[key])"> Delete</button>
										</div>
									</td>
								</tr>
							</div>
						</td>
					</tr>
				</td>
			</tr>
		</div>
		<div v-if="!stream_over">
			<button type="button" class="btn btn-sm btn-outline-primary" @click="refreshStream">Refresh</button>
		</div>
		<div v-if="stream_over">
			<p> No more photos in the stream </p>
		</div>
	</div>
</template>

<style>
[v-clock] {
	display: none;
}
</style>
