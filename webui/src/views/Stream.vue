

<script>

export default {
	data: function() {
		return {
			user: JSON.parse(sessionStorage.getItem('User')),
			profile: JSON.parse(sessionStorage.getItem('Profile')),
			errormsg: null,
			loading: false,
			photos: [],

			title: "",
			photoUploaded: null,
			search_username: "",

			isOwner: false,

			hasBanned: null,
			isBanned: null,
			isFollowing: null,

			comment_text: {},

			new_username: "",
			

		}
	},
	methods: {

		async getProfile(reset = false, test = true) {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/users/"+this.search_username, {
						headers: {
							'Authorization': this.user.userId,
							'userId': this.user.userId
						}
					}
				);
				if (response.data.userId == this.user.userId) {
					sessionStorage.setItem('User', JSON.stringify(response.data));
					this.user = response.data;
					if (reset) {
						sessionStorage.setItem('Profile', JSON.stringify(response.data));
						this.profile = response.data;
					}
				} else {
					sessionStorage.setItem('Profile', JSON.stringify(response.data));
					this.profile = response.data;
				}
				this.loading = false;
				if (test) {
					window.location.reload();
				}

			} catch (e) {
				if (e.response.status == 403) {
					this.errormsg = "Banned from the user.";
					console.log(this.errormsg)
					
				} else {
					this.errormsg = e.toString();
					console.log(this.errormsg)
				}
			}
		},

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
					this.search_username = this.user.username
					this.comment_text[id] = null
					this.getProfile()

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
				this.search_username = this.user.username
				this.getProfile()

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
				this.search_username = this.user.username
				this.getProfile()

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
				this.search_username = this.user.username
				this.getProfile()

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
				if (response.data == "There are no photos in the stream") {
					console.log(response.data)
				} else {
					this.photos = this.photos.concat(response.data);
					this.photos = response.data.reverse();
					for (let i = 0; i < this.photos.length; i++) {
						this.photos[i].File = 'data:image/*;base64,' + this.photos[i].File
					}
					console.log("Stream loaded")
					console.log(response.data)
				}
			} catch (e) {
				this.errormsg = e.toString();
				console.log(this.errormsg)
			} 
		},

	},
	mounted() {
		console.log("mounted_Stream", new Date().toLocaleTimeString());
		// this.getMyStream()
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Stream</h1>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<div v-if="photos" v-for="photo in photos" :key="photo.photoId" class="photo-entry">
			<tr>
				<td>
					<h3>{{ photo.title }}</h3>
					<img :src="photo.File" alt="Photo" class="photo-container">
					<p>PhotoId: {{ photo.photoId }}</p>
					<p>Date: {{ new Date(photo.date).toLocaleDateString() }}</p>
				</td>
				<td>
					<tr>
						<td>
							<button v-if="isOwner" type="submit" class="btn btn-sm btn-outline-primary" @click="deletePhoto(photo.photoId)">Delete Photo</button>
						</td>
					</tr>
					<tr>
						<td>
							<button v-if="photo.likeAuthors ? !photo.likeAuthors.includes(user.userId) : true" type="submit" class="btn btn-sm btn-outline-primary" @click="likePhoto(photo.photoId)">Like</button>
							<button v-if="photo.likeAuthors ? photo.likeAuthors.includes(user.userId) : false" type="submit" class="btn btn-sm btn-outline-primary" @click="unlikePhoto(photo.photoId)">Unlike</button>
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
										<button v-if="photo.commentAuthors[key] == user.username" type="submit" class="btn btn-sm btn-outline-primary" @click="uncommentPhoto(photo.photoId, photo.commentIds[key])"> Delete</button>
									</td>
								</tr>
							</div>
						</td>
					</tr>
				</td>
			</tr>
		</div>
		<div v-else>
			<p>No photos</p>
		</div>
		<div>
			<button type="button" class="btn btn-sm btn-outline-primary" @click="getMyStream">New</button>
		</div>
	</div>
</template>

<style>
</style>
