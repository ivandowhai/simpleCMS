<script>
    import { Navigate } from 'svelte-router-spa'
    import Http from '../../services/http'

    let posts = getPosts()

    async function getPosts() {
        return await new Http().get('/post');
    }

</script>

<main>
    {#await posts}
        <div>...waiting</div>
    {:then posts}
        {#each posts as post}
            <div>
                <h2>
                    <Navigate to="/post/{post.ID}">
                        {post.Title}
                    </Navigate>
                </h2>
                <p>  {post.Content}</p>
            </div>
        {/each}
    {/await}
</main>