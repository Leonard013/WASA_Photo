

<script>

export default {
	data: function() {
		return {
			user: JSON.parse(sessionStorage.getItem('User')),
			profile: JSON.parse(sessionStorage.getItem('Profile')),
			errormsg: null,
			loading: false,
			photos: null,

			title: null,
			photoUploaded: null,
			search_username: null,

			isOwner: false,

			hasBanned: null,
			isBanned: null,
			isFollowing: null,


		}
	},
	methods: {

		backToProfile() {
			sessionStorage.setItem('Profile', JSON.stringify(this.user));
			window.location.reload();
		},

		async deletePhoto(Id) {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.delete("/photos/"+Id, {
						headers: {
							'Authorization': this.user.userId,
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
			console.log(this.user)
			console.log(this.profile)

			if ((this.user.banned == null) || !(this.user.banned.includes(this.profile.userId)) ) { // Check if user is banned from the user_profile page
				this.hasBanned = false;
			} else {
				this.hasBanned = true;
			}

			if ((this.user.isBanned == null) || !(this.user.isBanned.includes(this.profile.userId)) ) { // Check if user is banned from the user_profile page
				this.isBanned = false;
			} else {
				this.isBanned = true;
			}

			if ((this.user.following == null) || !(this.user.following.includes(this.profile.userId)) ) { // Check if user is following the user_profile page
				this.isFollowing = false;
			} else {
				this.isFollowing = true;
			}
			console.log("hasBanned: " + this.hasBanned)
			console.log("isBanned: " + this.isBanned)
			console.log("isFollowing: " + this.isFollowing)
		},

		async getPhotos() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/photos/"+this.profile.username, {
						headers: {
							'Authorization': this.user.userId,
						}
					}
				);
				if (!response.data.message) {
					this.photos = response.data;
					for (let i = 0; i < this.photos.length; i++) {
					this.photos[i].File = 'data:image/*;base64,' + this.photos[i].File
					}
					console.log(this.photos);
				}
				this.loading = false;
				
			} catch (e) {
				this.errormsg = e.toString();
				console.log(this.errormsg)
			}
			
			this.loading = false;
		},

		async getProfile() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/users/"+this.search_username, {
						headers: {
							'Authorization': this.user.userId,
						}
					}
				);
				if (response.data.userId == this.user.userId) {
					sessionStorage.setItem('User', JSON.stringify(response.data));
					this.user = response.data;
				} else {
					sessionStorage.setItem('Profile', JSON.stringify(response.data));
					this.profile = response.data;
				}
				this.loading = false;
				window.location.reload();
			} catch (e) {
				if (e.response.status == 403) {
					this.errormsg = "Banned from the user";
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
							this.loading = false
							console.log("Photo not uploaded")
							console.log(error)
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
	},
	mounted() {
		if (this.user.userId == this.profile.userId) {
			this.isOwner = true;
		} else {
			this.getInfo()
		}
		console.log("mounted_Account")
		this.getPhotos()
	}
}
</script>

<template>
	<div>
		<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">

			<h1 class="h2">
				Profile of {{ profile.username }} 
			</h1>
			
			<div v-if="isOwner" class="btn-toolbar mb-2 mb-md-0">

				<div class="btn-group me-2">
					<tr>
						<td>
							<input v-model="title" placeholder="Enter title photo"> 
						</td>
						<td>
							<input type="file" ref="photoUploaded" accept="image/png">
						</td>
						<button  type="submit" class="btn btn-sm btn-outline-primary" @click="uploadPhoto">Add Photo</button>
					</tr>
				</div>
			</div>

			<div v-if="!isBanned && !isOwner">
				<button v-if="!isFollowing" type="submit" class="btn btn-sm btn-outline-primary" @click="followUser">Follow</button>
				<button v-if="isFollowing" type="submit" class="btn btn-sm btn-outline-primary" @click="unfollowUser">Unfollow</button>
				<button v-if="!hasBanned" type="submit" class="btn btn-sm btn-outline-primary" @click="banUser">Ban</button>
				<button v-if="hasBanned" type="submit" class="btn btn-sm btn-outline-primary" @click="unbanUser">Unban</button>
			</div>

			<div>
				<input v-model="search_username" placeholder="Enter username you are looking for">
				<button type="submit" class="btn btn-sm btn-outline-primary" @click="getProfile">Search</button>
				<button v-if="!isOwner" type="submit" class="btn btn-sm btn-outline-primary" @click="backToProfile">Back to Profile</button>
			</div>

		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<div v-if="!isBanned">
			<div v-if="isOwner" >
				Account ID: {{ this.profile.userId }}
				<br>
				<p v-for = "(value, key) in this.profile" :key="key">
					{{ key }}: {{ value }}
				</p>
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
						<button v-if="isOwner" type="submit" class="btn btn-sm btn-outline-primary" @click="deletePhoto(photo.photoId)">Delete</button>
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

</style>
