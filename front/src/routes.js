import postsList from './components/post/index.svelte'
import postsView from './components/post/view.svelte'
import login from './components/auth/login.svelte'
import register from './components/auth/register.svelte'
import mainLayout from './components/mainLayout.svelte'

const routes = [
    {
        name: '/',
        component: postsList,
        layout: mainLayout,
    },
    {
        name: '/post/:id',
        component: postsView,
        layout: mainLayout,
    },
    {
        name: '/login',
        component: login,
        layout: mainLayout,
    },
    {
        name: '/register',
        component: register,
        layout: mainLayout,
    }
];

export default routes