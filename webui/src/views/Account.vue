

<script>

export default {
	data: function() {
		return {
			user: JSON.parse(sessionStorage.getItem('User')),
			profile: JSON.parse(sessionStorage.getItem('Profile')),
			errormsg: null,
			loading: false,
			photos: "Non ci sono foto",

			title: null,
			photoUploaded: null,

			isOwner: false,


		}
	},
	methods: {

		async getPhotos() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/photos/"+this.profile.username, {
						headers: {
							'Authorization': this.profile.userId,
						}
					}
				);
				// this.profile = response.data.photos
				// this.followers = response.data.followers
				// this.following = response.data.following

				// for (let i = 0; i < this.profile.length; i++) {
				// 	this.profile[i].image = 'data:image/*;base64,' + this.profile[i].image
				// }
				this.photos = response.data;
				
				for (let i = 0; i < this.photos.length; i++) {
					this.photos[i].File = 'data:image/*;base64,' + this.photos[i].File
				}




				console.log(this.photos);
			} catch (e) {
				this.errormsg = e.toString();
				console.log(this.errormsg)
			}
			
			this.loading = false;
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
	},
	mounted() {
		if (this.user.userId == this.profile.userId) {
			this.isOwner = true;
		}
		console.log("mounted_Account")
		this.getPhotos()
		// this.refresh()


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
					<input v-model="title" placeholder="Enter title photo"> 

					<input type="file" ref="photoUploaded" accept="image/png">
					<button  type="submit" class="btn btn-sm btn-outline-primary" @click="uploadPhoto">Add Photo</button>
				</div>
			</div>

		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<div v-if="isOwner" >
			Account ID: {{ profile.userId }}
		</div>

		<div v-for="photo in photos" :key="photo.photoId" class="photo-entry">
			<h3>{{ photo.title }}</h3>
			<img :src="photo.File" alt="Photo" class="photo-container">
			<p>Date: {{ new Date(photo.date).toLocaleDateString() }}</p>
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
