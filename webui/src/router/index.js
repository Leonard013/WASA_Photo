import {createRouter, createWebHashHistory} from 'vue-router'
import Original from '../views/Original.vue'
import Account from '../views/Account.vue'
import Login from '../views/Login.vue'
import Login_2 from '../views/Login_2.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: Login_2},
		{path: '/link1', component: Original},
		{path: '/stream', component: Original},
		{path: '/account', component: Account},
		{path: '/profile/:username', component: Account},

		{path: '/some/:id/link', component: Original},
	]
})

export default router
