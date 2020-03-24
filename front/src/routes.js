import postsList from './components/post/index.svelte'
import postsView from './components/post/view.svelte'
import login from './components/auth/login.svelte'

const routes = [
    {
        name: '/',
        component: postsList
    },
    {
        name: 'post/:id',
        component: postsView
    },
    {
        name: 'login',
        component: login
    }
];

export default routes