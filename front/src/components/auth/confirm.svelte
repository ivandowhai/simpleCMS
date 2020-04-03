<script>
    import {navigateTo} from 'svelte-router-spa'
    import Http from '../../services/http'

    let error
    let code = ''
    let email = ''

    async function confirm(e) {
        e.preventDefault()

        await new Http().post('/confirm', {code: code, email: email})
                .then(response => {
                    alert(response.Result)
                    navigateTo('/')
                })
                .catch(e => {error = e})
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
                    <label for="name">Email</label>
                    <input class="form-control" type="email" name="email" bind:value="{email}"/>
                </div>

                <div class="col-md-12">
                    <label for="name">Code</label>
                    <input class="form-control" type="text" name="code" bind:value="{code}"/>
                </div>

                <div class="col-md-12">
                    <label></label>
                </div>

                <div class="col-md-12">
                    <button class="btn btn-primary" on:click={confirm}>Confirm</button>
                </div>
            </div>
        </div>
    </form>
</main>