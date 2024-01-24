import {createRouter, createWebHashHistory} from 'vue-router'
import Original from '../views/Original.vue'
import Account from '../views/Account.vue'
import Login from '../views/Login.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: Login},
		{path: '/link1', component: Original},
		{path: '/stream', component: Original},
		{path: '/account', component: Account},

		{path: '/some/:id/link', component: Original},
	]
})

export default router
