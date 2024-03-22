import {createRouter, createWebHashHistory} from 'vue-router'
import Account from '../views/Account.vue'
import Login from '../views/Login.vue'
import Stream from '../views/Stream.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: Login},
		{path: '/stream', component: Stream},
		{path: '/account', component: Account},
	]
})

export default router
