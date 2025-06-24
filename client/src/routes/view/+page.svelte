<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';

	import view from '$stores/view';
	import auth from '$stores/auth';
	import persistantStore from '$lib/utils/persistantStore';
	import type { TableFilterType } from '$types/tableFilter';

	import Table from './components/Table/Table.svelte';
	import Header from './components/Header/Header.svelte';
	import Pagination from './components/Pagination/Pagination.svelte';

	const defautlFilter: TableFilterType = {
		name: true,
		gender: true,
		phone: false,
		qq: false,
		email: false,
		student_id: true,
		college: false,
		major: false,
		created_at: false,
		first_choice: true,
		second_choice: true,
		introduction: true
	};

	let tableFilter = persistantStore('tableFilter', defautlFilter);

	onMount(async () => {
		if ($auth || (await auth.ping()).success) view.refersh();
		else goto('/login/?callback=/view', { replaceState: true });
	});
</script>

<svelte:head>
	<title>报名数据 • 校大学生科协</title>
</svelte:head>

<main>
	{#if $view}
		<Header {tableFilter} />
		{#if $view.applicants.length > 0}
			<Table applicants={$view.applicants} pagination={$view.pagination} tableFilter={$tableFilter} />
			<Pagination pagination={$view.pagination} />
		{:else}
			<p>这里什么都没有😢</p>
		{/if}
	{:else}
		<p>数据加载中，请稍后...</p>
	{/if}
</main>

<style>
	main {
		@apply px-4 sm:px-10 pb-20 pt-5 sm:pt-10 flex flex-col gap-3;
	}

	p {
		@apply w-full flex-1 flex items-center justify-center;
	}
</style>
