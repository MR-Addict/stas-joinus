<script lang="ts">
	import auth from '$stores/auth';
	import { page } from '$app/stores';

	import Menu from './components/Menu/Menu.svelte';

	let pathname = '/';

	$: pathname = $page.url.pathname.split('/').slice(0, 2).join('/');
</script>

<nav>
	<a href="/" class="flex flex-row items-center gap-1">
		<img src="/favicon.png" alt="logo" class="w-7" />
		<h1 class="font-semibold text-gray-800">校大学生科协</h1>
	</a>

	<div data-sveltekit-preload-data class="ml-auto space-x-2">
		{#if $auth}
			<a href="/table" class="link" class:active={pathname === '/table'}>表格</a>
			<a href="/analytics" class="link" class:active={pathname === '/analytics'}>统计</a>
		{/if}
	</div>

	<Menu />
</nav>

<style lang="postcss">
	nav {
		height: 60px;
		@apply flex flex-row items-center gap-4 px-4 md:px-20 border-b border-b-gray-300;
	}
	.link {
		@apply text-sm text-gray-500;
		&:hover,
		&.active {
			@apply text-black;
		}
	}
</style>
