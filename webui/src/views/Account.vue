

<script>

export default {
	data: function() {
		return {
			user: JSON.parse(sessionStorage.getItem('User')),
			profile: JSON.parse(sessionStorage.getItem('Profile')),
			errormsg: null,
			loading: false,
			photos: null,

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

		backToProfile() {
			// sessionStorage.setItem('Profile', JSON.stringify(this.user));
			// window.location.reload();
			this.search_username = this.user.username
			this.getProfile(true)
		},

		async deletePhoto(Id) {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.delete("/photos/"+Id, {
						headers: {
							'Authorization': this.user.userId,
							'userId': this.user.userId
						}
					}
				);
				this.loading = false;
				console.log("Photo deleted")
				window.location.reload()
			} catch (e) {
				this.errormsg = e.toString();
				console.log(this.errormsg)
			}
		},

		getInfo() {
			if (this.user.banned == null) {
				this.hasBanned = false;
				this.user.banned = [];
			} else if (!this.user.banned.includes(this.profile.userId))  { // Check if user is banned from the user_profile page
				this.hasBanned = false;
			} else {
				this.hasBanned = true;
			}

			if (this.user.isBanned == null){
				this.isBanned = false;
				this.user.isBanned = [];
			} else if (!this.user.isBanned.includes(this.profile.userId)) { // Check if user is banned from the user_profile page
				this.isBanned = false;
			} else {
				this.isBanned = true;
			}

			if (this.user.following == null) {
				this.isFollowing = false;
				this.user.following = [];
			} else if (!this.user.following.includes(this.profile.userId)) { // Check if user is following the user_profile page
				this.isFollowing = false;
			} else {
				this.isFollowing = true;
			}

			if (this.user.followers == null) {
				this.user.followers = [];
			}

			console.log("hasBanned: " + this.hasBanned)
			console.log("isBanned: " + this.isBanned)
			console.log("isFollowing: " + this.isFollowing)
			console.log(this.user)
		},

		async getPhotos() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/photos/"+this.profile.username, {
						headers: {
							'Authorization': this.user.userId,
							'userId': this.user.userId,
						}
					}
				);
				if (!response.data.message) {
					this.photos = response.data.reverse();
					for (let i = 0; i < this.photos.length; i++) {
						this.photos[i].File = 'data:image/*;base64,' + this.photos[i].File
					}
					
					console.log(this.photos);
				}
				this.loading = false;
				
			} catch (e) {
				if (e.response.status == 403) {
					this.errormsg = "Banned from the user";
					console.log(this.errormsg)
					this.search_username = this.user.username
					this.getProfile(true)
				} else {
					this.errormsg = e.toString();
					console.log(this.errormsg)
				}
			}
			
			this.loading = false;
		},

		async getProfile(reset = false, test = true) {
			this.loading = true;
			this.errormsg = null;
			console.log("getProfile executed")
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

		async uploadPhoto() {
			this.loading = true;
			this.errormsg = null;

			const formData = new FormData();
			const currentPhoto = this.$refs.photoUploaded.files[0];

			formData.append('title', this.title);
			formData.append('userId', this.user.userId); // Corrected user reference
			formData.append('image', currentPhoto);

			this.$axios.post("/photos/", formData, {
					headers: {
						'Authorization': this.user.userId,
					}
				}
			).then(() => {
						this.loading = false
						window.location.reload()
						console.log("Photo uploaded")
					}).catch(
						(error) => {
							if (error.response.status == 400 ) {
								this.errormsg = "No photo uploaded"
								console.log(this.errormsg)
							} else {
								this.errormsg = error.toString();
								console.log(this.errormsg)
							}
							this.loading = false
							this.photoUploaded = null
						}
					)
		},

		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/");
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},

		async followUser() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.post("/follow/", {
					username: this.profile.username,
					userId: this.user.userId
				}, {
						headers: {
							'Authorization': this.user.userId,
						}
					}
				);
				this.loading = false;
				console.log("User followed")
				this.search_username = this.user.username
				this.getProfile()
				
			} catch (e) {
				this.errormsg = e.toString();
				console.log(this.errormsg)
			}
		},

		async unfollowUser() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.delete("/follow/"+this.profile.username, {
						headers: {
							'Authorization': this.user.userId,
							'userId': this.user.userId,
						}
					}
				);
				this.loading = false;
				console.log("User unfollowed")
				this.search_username = this.user.username
				this.getProfile()

			} catch (e) {
				this.errormsg = e.toString();
				console.log(this.errormsg)
			}
		},

		async banUser() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.post("/ban/", {
					username: this.profile.username,
					userId: this.user.userId
				}, {
						headers: {
							'Authorization': this.user.userId,

						}
					}
				);
				this.loading = false;
				console.log("User banned")
				this.search_username = this.user.username
				this.getProfile()

			} catch (e) {
				this.errormsg = e.toString();
				console.log(this.errormsg)
			}
		},

		async unbanUser() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.delete("/ban/"+this.profile.username, {
						headers: {
							'Authorization': this.user.userId,
							'userId': this.user.userId,
						}
					}
				);
				this.loading = false;
				console.log("User unbanned")
				this.search_username = this.user.username
				this.getProfile()

			} catch (e) {
				this.errormsg = e.toString();
				console.log(this.errormsg)
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
		
		async changeUsername() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.put("/users/"+this.user.username, {
					username: this.new_username
				}, {
						headers: {
							'Authorization': this.user.userId,
							'userId': this.user.userId,
						}
					}
				);
				this.loading = false;
				console.log("Username changed to ", this.new_username)
				this.search_username = this.new_username
				this.getProfile(true)
			} catch (e) {
				if (e.response.status == 403) {
					this.errormsg = "Username already taken";
					console.log(this.errormsg)
				} else {
					this.errormsg = e.toString();
					console.log(this.errormsg)
				}
			}
		}


	},
	mounted() {
		if (this.user.userId == this.profile.userId) {
			this.isOwner = true;
			this.search_username = this.user.username
			this.getProfile(true,false)
			this.search_username = ""
		} else {
			this.search_username = this.profile.username
			this.getProfile(false,false)
			this.search_username = ""
			this.getInfo()
		}
		console.log("mounted_Account", new Date().toLocaleTimeString());
		console.log(this.profile)
		this.getPhotos()
	}
}
</script>

<template>
	<div v-cloak>
		<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">

			<h2 class="h2">
				Profile of {{ profile.username }} 
			</h2>
			
			<div v-if="isOwner" class="btn-toolbar mb-2 mb-md-0">
				<button type="submit" class="btn btn-sm btn-outline-primary" @click="backToProfile">Refresh</button>
				<div class="btn-group me-2">
					<tr>
						<td>
							<input v-model="title" placeholder="Enter title photo"> 
						</td>
						<td>
							<input type="file" ref="photoUploaded" accept="image/png">
						</td>
						<div v-if="title.length <= 30 && title.length > 0">
							<button  type="submit" class="btn btn-sm btn-outline-primary" @click="uploadPhoto">Add Photo</button>
						</div>

					</tr>
				</div>

			</div>

			<div v-if="!isBanned && !isOwner">
				<div v-if="!isFollowing">
					<button  type="submit" class="btn btn-sm btn-outline-primary" @click="followUser">Follow</button>
				</div>
				<div v-if="isFollowing">
					<button type="submit" class="btn btn-sm btn-outline-primary" @click="unfollowUser">Unfollow</button>
				</div>
				<div v-if="!hasBanned">
					<button type="submit" class="btn btn-sm btn-outline-primary" @click="banUser">Ban</button>
				</div>
				<div v-if="hasBanned">
					<button  type="submit" class="btn btn-sm btn-outline-primary" @click="unbanUser">Unban</button>
				</div>
			</div>

			<div>
				<input v-model="search_username" placeholder="Profile you are looking for">
				<div v-if="search_username.length >= 3 && search_username.length <= 20">
					<button type="submit" class="btn btn-sm btn-outline-primary" @click="getProfile">Search</button>
				</div>
				<div v-if="!isOwner">
					<button type="submit" class="btn btn-sm btn-outline-primary" @click="backToProfile">Back to Profile</button>
				</div>
			</div>

		</div>
		<div v-if="errormsg">
			<ErrorMsg :msg="errormsg"></ErrorMsg>
		</div>
		<div v-if="!isBanned">
			<div>
				<tr>
					<td>
						<tr>
							<td>
								<div v-if="isOwner">
									<input v-model="new_username" placeholder="Enter new username">
								</div>
							</td>
						</tr>
						<tr>
							<td>
								<div v-if="isOwner && (new_username.length >= 3 && new_username.length <= 20)">
									<button type="submit" class="btn btn-sm btn-outline-primary" @click="changeUsername">Change Username</button>
								</div>
							</td>
						</tr>
					</td>
					<td>
						<p>
							Followers: {{ profile.followers ? profile.followers.length : 0 }}
							<br>
							Following: {{ profile.following ? profile.following.length : 0 }}
						</p>
					</td>
				</tr>
			</div>
			
			<div v-if="isOwner" >
				Account ID: {{ this.profile.userId }}
				<!-- <br>
				<p v-for = "(value, key) in this.profile" :key="key">
					{{ key }}: {{ value }}
				</p> -->
			</div>

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
								<div v-if="isOwner">
									<button type="submit" class="btn btn-sm btn-outline-primary" @click="deletePhoto(photo.photoId)">Delete Photo</button>
								</div>
							</td>
						</tr>
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
											<div  v-if="photo.commentAuthors[key] == user.username">
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
			<div v-else>
				<p>No photos</p>
			</div>
		</div>

		<div v-if="isBanned && !isOwner">
			you are banned from seeing this profile
		</div>

	</div>
</template>

<style>
.photo-entry {
	margin-bottom: 20px;
}

.photo-container {
	width: 100%;
	max-width: 300px;
	height: auto;
}
.comments-box {
  border: 1px solid #ccc;
  margin-top: 10px;
  padding: 10px;
  max-height: 200px;
  overflow-y: auto;
}
[v-clock] {
	display: none;
}
</style>
