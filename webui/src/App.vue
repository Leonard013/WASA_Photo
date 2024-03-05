<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {
	data: function() {
		return {
			isLogged: false,
		}
	},
	mounted() {
		if (sessionStorage.getItem('LoggedIn')) {
			this.isLogged = true;
			console.log("The user is logged in |	App")
		} else {
			console.log("The user is not logged in |	App")
		}
	}
}

</script>

<template>

	<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
		<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="#/">Example App</a>
		<button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
			<span class="navbar-toggler-icon"></span>
		</button>
	</header>

	<div class="container-fluid">
		<div class="row">
			<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
				<div v-if="isLogged" class="position-sticky pt-3 sidebar-sticky">

					<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
						<span>General</span>
					</h6>

					<ul class="nav flex-column">

						<li class="nav-item">
							<RouterLink to="/stream" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
								Stream
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to="/link1" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#layout"/></svg>
								Menu item 1
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to="/account" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#key"/></svg>
								Profile
							</RouterLink>
						</li>

					</ul>

					<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
						<span>Secondary menu</span>
					</h6>
					
					<!-- <ul class="nav flex-column">
						<li class="nav-item">
							<RouterLink :to="'/some/' + 'variable_here' + '/path'" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#file-text"/></svg>
								Item 1
							</RouterLink>
						</li>
					</ul> -->
				</div>
			</nav>

			<main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
				<RouterView />
			</main>
		</div>
	</div>
</template>

<style>
</style>
