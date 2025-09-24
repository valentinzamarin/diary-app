import { createRouter, createWebHashHistory } from 'vue-router'
import ArticleEditor from '../components/article/ArticleEditor.vue'
import ArticleViewer from '../components/article/ArticleViewer.vue'

const routes = [
    { path: '/', name: 'home', component: ArticleEditor },
    { path: '/entry/:id', name: 'article', component: ArticleViewer, props: true },
    { path: '/edit/:id', name: 'edit', component: ArticleEditor, props: true }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router