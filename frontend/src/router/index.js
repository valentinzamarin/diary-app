import { createRouter, createWebHashHistory } from 'vue-router'
import ArticleList from '../components/article/ArticleList.vue'
import ArticleEditor from '../components/article/ArticleEditor.vue'
import ArticleViewer from '../components/article/ArticleViewer.vue'

const routes = [
    { path: '/', name: 'home', component: ArticleEditor },
    { path: '/entry/:id', name: 'article', component: ArticleViewer, props: true }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router