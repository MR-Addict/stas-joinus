<script lang="ts">
	import Login from '$components/Login/Login.svelte';
</script>

<svelte:head>
	<title>管理员登录 • 校大学生科协</title>
</svelte:head>

<main>
	<Login />
</main>

<style>
	main {
		@apply flex flex-col items-center justify-center px-4 pb-20 pt-5 sm:pt-10;
	}
</style>
