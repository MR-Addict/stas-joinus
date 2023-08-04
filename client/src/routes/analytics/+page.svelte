<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';

	import auth from '$stores/auth';

	onMount(async () => {
		const success = await auth.ping();
		if (!success) goto('/', { replaceState: true });
	});
</script>

<svelte:head>
	<title>欢迎加入校科协 • 后台管理</title>
</svelte:head>

<main>
	<p>统计页面</p>
</main>

<style lang="postcss">
	main {
		min-height: calc(100vh - 60px);
		min-height: calc(100svh - 60px);
		@apply py-5 md:py-10 px-4 md:px-20 flex flex-col items-center justify-center gap-2;
	}
</style>
