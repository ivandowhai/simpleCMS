<script>

	// TODO: https://linguinecode.com/post/how-to-add-routes-to-svelte-spa-with-routify use instead  'svelte-router-spa'
	import { Router } from 'svelte-router-spa'
	import routes from './routes'

	export let name;

</script>

<main>
	<nav class="navbar navbar-expand-lg navbar-light bg-light">

		<a class="navbar-brand" href="/">Blog</a>
<!--		<a class="btn btn-outline-success" href="/profile">Profile</a>-->
<!--		<a class="btn btn-outline-success" href="/logout">Logout</a>-->
<!--		<a class="btn btn-outline-success" href="/login">Login</a>-->
<!--		<a class="btn btn-outline-success" href="/register">Register</a>-->
<!--		{{end}}-->
	</nav>
	<Router {routes} />
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>