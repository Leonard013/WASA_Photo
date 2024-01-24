

<script>

import { sharedData } from '../main.js';


export default {
	data: function() {
		return {
			...sharedData,
			errormsg: null,
			loading: false,
			mario: "marione",
		}
	},
	methods: {
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
		this.refresh()
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Profile of {{ some_data.username }}</h1>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<span> Il counter è: {{ count }}</span><br>
		<span> Account ID: {{ some_data.userId }}</span>
	</div>
</template>

<style>
</style>
