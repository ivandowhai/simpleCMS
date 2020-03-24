<script>
    import Http from '../../services/http'

    let error
    let email
    let password

    async function login() {
        await new Http().post('/login', {email: email, password: password})
                .then(response => {
                    console.log(response)
                })
                .catch(error => {
                    console.log(error)
                })
    }

</script>

<main>
    <form action="/login-submit" method="post">
        {#if error}
            <div class="alert alert-danger">
                {error}
            </div>
        {/if}}
        <div class="container">
            <div class="row">
                <div class="col-md-12">
                    <label for="email">Email</label>
                    <input class="form-control" type="email" name="email" value="{email}"/>
                </div>

                <div class="col-md-12">
                    <label for="password">Password</label>
                    <input class="form-control" type="password" name="password" value="password"/>
                </div>

                <div class="col-md-12">
                    <label></label>
                </div>

                <div class="col-md-12">
                    <button class="btn btn-primary" on:click={login}>Login</button>
                </div>
            </div>
        </div>
    </form>
</main>