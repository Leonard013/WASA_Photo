import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: HomeView},
		{path: '/link1', component: PlayView},
		// {path: '/link2', component: HomeView},
		{path: '/some/:id/link', component: HomeView},
	]
})

export default router