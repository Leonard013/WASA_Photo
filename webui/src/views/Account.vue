

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
			try {
				let response = await this.$axios.post("/photos");
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;


			this.$axios.post("/photos", {
				title: this.title,
				userId: user.userId,
				image: this.image,

			}).then(() => {
						this.success = true
						this.error = false
					}).catch(
						() => {
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
					<input type="file" ref="photo" accept="image/png">
					<button type="submit" class="btn btn-sm btn-outline-primary">Add Photo</button>

					<!-- <button type="button" class="btn btn-sm btn-outline-primary" @click="uploadPhoto">
						Post a New Photo
					</button> -->
				</div>
			</div>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<div >
			Account ID: {{ user.userId }}
		</div>

		<div >
			le foto sono: {{ photos }}
		</div>
	</div>
</template>

<style>
</style>
