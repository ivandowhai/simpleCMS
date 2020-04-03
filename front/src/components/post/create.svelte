<script>
    import Http from '../../services/http'
    import { navigateTo } from 'svelte-router-spa'

    let error
    let title = ''
    let content = ''

    async function create(e) {
        e.preventDefault()

        await new Http().post('/post/store', {title: title, content: content})
            .then(() => {navigateTo('/')})
            .catch(err => {error = err})
    }

</script>

<main>
    <form method="post">
        {#if error}
            <div class="alert alert-danger">
                {error}
            </div>
        {/if}

        <div class="container">
            <div class="row">
                <div class="col-md-12">
                    <label for="name">Title</label>
                    <input class="form-control" type="text" name="title" bind:value="{title}"/>
                </div>

                <div class="col-md-12">
                    <label for="content">Content</label>
                    <textarea class="form-control" name="content" bind:value="{content}"></textarea>
                </div>

                <div class="col-md-12">
                    <label></label>
                </div>

                <div class="col-md-12">
                    <button class="btn btn-primary" on:click={create}>Create</button>
                </div>
            </div>
        </div>
    </form>
</main>