<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			regalo: "ghepardo",
			username: null,
			name: null,
			userId: null,
			some_data: "ghepardo",
			count: 0,
		}
	},
	methods: {
		async login() {
			this.loading = true;
			this.errormsg = null;
			this.name = document.getElementById("username").value;
			try {
				let response = await this.$axios.post("/session", {username: this.name});
				this.some_data = response.data;
				this.$router.push('/account');
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.name = null;
			this.loading = false;
			
		},
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/");
				this.some_data = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			console.log(response.data);
			this.loading = false;
		},
	},
	mounted() {
		this.refresh()
	}
}
</script>

<template>
	<div class="container">
	  <div class="header">
		<h1>Welcome to Leonardo</h1>
	  </div>
	  <div class="login-container">
		<div class="form-group">
			<label for="username">username</label>
			<input type="text" id="username" v-model="username" placeholder="Enter username">
			<button type="submit" class="login-button" @click="login()">Login</button>
		</div>
	  </div>
	  <span> {{ name }}</span>
	  <span> {{ some_data }}</span>
	</div>
  </template>



<style scoped>
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