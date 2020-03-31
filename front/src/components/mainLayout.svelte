<script>
    import {Router, Route} from 'svelte-router-spa'
    import {Navigate} from 'svelte-router-spa'

    export let currentRoute
    export let name;

    const params = {}

    let isLogged = localStorage.getItem('token') !== null && localStorage.getItem('token') !== ''

    function logout() {
        localStorage.setItem('token', '')
        isLogged = false
    }

</script>

<main>
    <nav class="navbar navbar-expand-lg navbar-light bg-light">


        <Navigate class="navbar-brand" to="/">Blog</Navigate>

        {#if isLogged}
            <Navigate class="btn btn-outline-success" to="/profile">Profile</Navigate>
            <span class="btn btn-outline-success" on:click={logout}>Logout</span>
        {:else}
            <Navigate class="btn btn-outline-success" to="/login">Login</Navigate>
            <Navigate class="btn btn-outline-success" to="/register">Register</Navigate>
        {/if}

    </nav>

    <section class="section">
        <Route {currentRoute} {params}/>
    </section>

</main>