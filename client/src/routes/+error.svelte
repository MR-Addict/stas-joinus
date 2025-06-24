<script lang="ts">
	import { page } from '$app/state';
</script>

<svelte:head>
	<title>校大学生科协 • 页面错误</title>
</svelte:head>

<main>{page.status} | {page.error?.message || 'Error'}</main>

<style>
	main {
		min-height: calc(100vh - 60px);
		min-height: calc(100svh - 60px);
		@apply w-full flex-1 flex items-center justify-center text-gray-700;
	}
</style>
