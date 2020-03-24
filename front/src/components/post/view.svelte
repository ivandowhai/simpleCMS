<script>
    import Http from '../../services/http'
    export let currentRoute

    let post = getPost()

    async function getPost() {
        return await new Http().get('/post/view/' + currentRoute.namedParams.id);
    }

</script>

<main>
    {#await post}
        <div>...waiting</div>
    {:then post}
        <h1>{post.Title}</h1>
        <div>
            {post.Content}
        </div>
    {/await}
</main>