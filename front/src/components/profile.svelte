<script>
    import {Navigate} from 'svelte-router-spa'
    import Http from '../services/http'


</script>

<main>
    <div class="jumbotron">
<!--        <h1>Hello, {{ .User.Name }}!</h1>-->
        <Navigate to="/post/create" class="btn btn-primary">Create post</Navigate>
    </div>

    <table class="table table-striped">
        <thead>
        <tr>
            <td>Title</td>
            <td>Actions</td>
        </tr>
        </thead>

<!--        <tbody>-->
<!--        {{ range $post := .Posts }}-->
<!--        <tr>-->
<!--            <td> {{ $post.Title }}</td>-->
<!--            <td>-->
<!--                <a href="/posts/view/{{ $post.ID }}" class="btn btn-sm btn-success">View</a>-->
<!--                <a href="/posts/edit/{{ $post.ID }}" class="btn btn-sm btn-info">Edit</a>-->
<!--                <a href="/posts/delete/{{ $post.ID }}" class="btn btn-sm btn-danger">Delete</a>-->
<!--            </td>-->
<!--        </tr>-->
<!--        {{ end }}-->
<!--        </tbody>-->
    </table>
</main>