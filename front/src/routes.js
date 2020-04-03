import postsList from './components/post/index.svelte'
import postsView from './components/post/view.svelte'
import postCreate from './components/post/create.svelte'
import login from './components/auth/login.svelte'
import register from './components/auth/register.svelte'
import confirm from './components/auth/confirm.svelte'
import profile from './components/profile.svelte'
import mainLayout from './components/mainLayout.svelte'

const routes = [
    {
        name: '/',
        component: postsList,
        layout: mainLayout,
    },
    {
        name: '/post/view/:id',
        component: postsView,
        layout: mainLayout,
    },
    {
        name: '/post/create',
        component: postCreate,
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
    },
    {
        name: '/confirm',
        component: confirm,
        layout: mainLayout,
    },
    {
        name: '/profile',
        component: profile,
        layout: mainLayout,
    }
];

export default routes