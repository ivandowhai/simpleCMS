<script>
    import Http from '../services/http'

    let posts = getPosts()

    async function getPosts() {
        return await Http.get('post');
    }

</script>

<main>
    {#await posts}
        <div>...waiting</div>
    {:then posts}
        {#each posts as post}
            <div>
                <h2>
                    <a href="/posts/view/{post.ID}">
                        {post.Title}
                    </a>
                </h2>
                <p>  {post.Content}</p>
            </div>
        {/each}
    {/await}
</main>