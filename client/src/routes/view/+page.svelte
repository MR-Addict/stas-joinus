<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { browser } from '$app/environment';

	import view from '$stores/view';
	import auth from '$stores/auth';
	import filter from './lib/filter';

	import Table from './components/Table/Table.svelte';
	import Header from './components/Header/Header.svelte';
	import Pagination from './components/Pagination/Pagination.svelte';

	let loaded = false;
	let tableFilter = filter.default;

	$: if (browser && loaded) filter.save(tableFilter);

	onMount(async () => {
		if ($auth || (await auth.ping()).success) view.refersh();
		else goto('/', { replaceState: true });

		tableFilter = filter.load();
		loaded = true;
	});
</script>

<svelte:head>
	<title>报名数据 • 校大学生科协</title>
</svelte:head>

<main>
	{#if $view}
		{#if $view.applicants.length > 0}
			<Header bind:tableFilter />
			<Table applicants={$view.applicants} pagination={$view.pagination} {tableFilter} />
			<Pagination pagination={$view.pagination} />
		{:else}
			<p>糟糕！还没有人提交</p>
		{/if}
	{:else}
		<p>数据加载中，请稍后...</p>
	{/if}
</main>

<style>
	main {
		@apply py-5 sm:py-10 px-4 sm:px-10 flex flex-col gap-3;
	}

	p {
		@apply w-full flex-1 flex items-center justify-center;
	}
</style>
