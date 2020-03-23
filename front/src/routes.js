import postsList from './components/post/index.svelte'
import postsView from './components/post/view.svelte'

const routes = [
    {
        name: '/',
        component: postsList
    },
    {
        name: 'post/:id',
        component: postsView
    }
];

export default routes