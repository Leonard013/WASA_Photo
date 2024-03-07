

<script>

export default {
	data: function() {
		return {
			user: JSON.parse(sessionStorage.getItem('User')),
			errormsg: null,
			loading: false,
			photos: "Non ci sono foto",

			title: null,
			photo: null,

			


		}
	},
	methods: {

		async getPhotos() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/photos/"+user.username, {
						username: this.name
					}, {
						Headers: {
							'Authorization': this.user.userId,
							'userId': this.user.userId
						}
					}
				);
				this.photos = response.data.Message;
			} catch (e) {
				this.errormsg = e.toString();
			}
			
			this.name = null;
			this.loading = false;
		},


		async uploadPhoto() {
			this.loading = true;
			this.errormsg = null;

			const formData = new FormData();
			const currentPhoto = this.$refs.photo.files[0];

			formData.append('title', this.title);
			formData.append('userId', this.user.userId); // Corrected user reference
			formData.append('image', currentPhoto);

			this.$axios.post("/photos/", formData, {
					Headers: {
						'Authorization': this.user.userId,
						'userId': this.user.userId
					}
				}
			).then(() => {
						this.loading = false
						console.log("Photo uploaded")
						this.success = true
						this.error = false
					}).catch(
						(error) => {
							console.log("Photo not uploaded")
							console.log(error)
							this.success = false
							this.error = true
							this.photo = null
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
		// this.getPhotos()
		// this.refresh()
	}
}
</script>

<template>
	<div>
		<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">

			<h1 class="h2">
				Profile of {{ user.username }} 
			</h1>
			
			<div class="btn-toolbar mb-2 mb-md-0">

				<div class="btn-group me-2">
					<input v-model="title" placeholder="Enter title photo"> 

					<input type="file" ref="photo" accept="image/png">
					<button  type="submit" class="btn btn-sm btn-outline-primary" @click="uploadPhoto">Add Photo</button>
				</div>
			</div>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<div >
			Account ID: {{ user.userId }}
		</div>

		<div >
			le foto sono: {{ photo }}
		</div>
	</div>
</template>

<style>
</style>
