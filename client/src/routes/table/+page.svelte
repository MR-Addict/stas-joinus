<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';

	import auth from '$stores/auth';
	import applicants from '$stores/applicants';
	import fetchApplicants from '$lib/applicant/fetchApplicants';
	import type { TableFilter } from '$types/tableFilter';

	import Table from './components/Table/Table.svelte';
	import Header from './components/Header/Header.svelte';
	import Pagination from './components/Pagination/Pagination.svelte';

	let tableFilter: TableFilter = {
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

	onMount(async () => {
		if ($auth || (await auth.ping())) fetchApplicants();
		else goto('/', { replaceState: true });
	});
</script>

<svelte:head>
	<title>校大学生科协 • 后台管理</title>
</svelte:head>

<main>
	{#if $applicants}
		{#if $applicants.length > 0}
			<Header bind:tableFilter />
			<Table applicants={$applicants} {tableFilter} />
			<Pagination />
		{:else}
			<p>Woops！还没有任何数据哦</p>
		{/if}
	{:else}
		<p>数据加载中，请稍后...</p>
	{/if}
</main>

<style lang="postcss">
	main {
		@apply flex-1 py-5 sm:py-10 px-4 sm:px-20 flex flex-col gap-2;
	}
	p {
		@apply w-full flex-1 flex items-center justify-center;
	}
</style>
