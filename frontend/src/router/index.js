import { createMemoryHistory, createRouter } from 'vue-router'
import Frontpage from '../components/Frontpage.vue';

const routes = [
    { path: '/', component: Frontpage },
]

const router = createRouter({
    history: createMemoryHistory(),
    routes,
})

export default router;