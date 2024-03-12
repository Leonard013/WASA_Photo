

<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			username: null,
			userId: null,

			some_data: null,
		}
	},
	methods: {
		async refresh() {
			this.loading = true;
			this.errormsg = null;

			this.$axios.get("/").then((response) => {
				this.some_data = response.data;
				console.log(response.data);
			}).catch(
				(error) => {
					console.log(error);
				}
			);
			this.loading = false;
		},

		async login(usern) {
			this.loading = true;
			this.errormsg = null;

			sessionStorage.setItem('LoggedIn', true);
			console.log(usern + " logged | Login_2");
			this.$axios.post("/session", {username: usern}).then((response) => {
					sessionStorage.setItem('User', JSON.stringify(response.data));
					sessionStorage.setItem('Profile',JSON.stringify(response.data)) 
					console.log(response.data);
					this.$router.push('/account');
					window.location.reload();
				}).catch(
					(error) => {
						console.log(error);
					}
				);

			this.loading = false;

		}
	},
	mounted() {
		if (sessionStorage.getItem('LoggedIn')) {
			this.$router.push('/account');
		}
		console.log("mounted_Login_2");
	}
}
</script>

<template>
	<div class="container">

		<div class="header">
			<h1>Welcome to WASA Photo</h1>
		</div>

		<div class="login-container">
			<div class="form-group">
				<input v-model="username" placeholder="Enter username" /> 
			</div>
			<button @click="login(username)" class="login-button"> Login </button>
		</div>
	</div>
</template>



<style scoped>
.box{
	background-color: purple;
	height: 100px;
	width: 100px;
}

.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
}

.header h1 {
	align-items: center;
	margin-bottom: 20px;
}

.login-container {
  width: 300px;
  align-items: center;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 5px;
}

.form-group {
  margin-bottom: 15px;
  align-items: center;
}

.form-group label {
  display: block;
  align-items: center;
  margin-bottom: 5px;
}

.login-button {
  width: 100%;
  padding: 10px;
  align-items: center;
  border: none;
  border-radius: 5px;
  background-color: #4CAF50;
  color: white;
  cursor: pointer;
  transition: background-color 0.3s;
}

.login-button:hover {
  background-color: #45a049;
}

.error-message {
  color: red;
  font-size: 0.9em;
}

.signup-link {
  text-align: center;
  margin-top: 15px;
}

.signup-link a {
  color: #007bff;
  text-decoration: none;
}
</style>