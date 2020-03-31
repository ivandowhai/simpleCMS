<script>
    import {Router, Route} from 'svelte-router-spa'
    import {Navigate} from 'svelte-router-spa'

    export let currentRoute
    export let name;

    const params = {}

    let isLogged = localStorage.getItem('token') !== null && localStorage.getItem('token') !== ''
    console.log(localStorage.getItem('token'))
</script>

<main>
    <nav class="navbar navbar-expand-lg navbar-light bg-light">


        <Navigate class="navbar-brand" to="/">Blog</Navigate>

        {#if isLogged}
            <Navigate class="btn btn-outline-success" to="/profile">Profile</Navigate>
            <Navigate class="btn btn-outline-success" to="/logout">Logout</Navigate>
        {:else}
            <Navigate class="btn btn-outline-success" to="/login">Login</Navigate>
            <Navigate class="btn btn-outline-success" to="/register">Register</Navigate>
        {/if}

    </nav>

    <section class="section">
        <Route {currentRoute} {params}/>
    </section>

</main>