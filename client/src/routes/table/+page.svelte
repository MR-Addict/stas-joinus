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
		college: true,
		major: true,
		created_at: false,
		first_choice: true,
		second_choice: true,
		introduction: true
	};

	let loading = true;

	onMount(async () => {
		if ($auth || (await auth.ping())) fetchApplicants().then(() => (loading = false));
		else goto('/', { replaceState: true });
	});
</script>

<svelte:head>
	<title>欢迎加入校科协 • 后台管理</title>
</svelte:head>

<main>
	{#if $applicants.length > 0}
		<Header bind:tableFilter />
		<Table {tableFilter} />
		<Pagination />
	{:else if !loading}
		<p>Woops！还没有任何数据哦</p>
	{/if}
</main>

<style lang="postcss">
	main {
		min-height: calc(100vh - 60px);
		min-height: calc(100svh - 60px);
		@apply py-5 md:py-10 px-4 md:px-20 flex flex-col gap-2;
	}
	p {
		@apply w-full flex-1 flex items-center justify-center;
	}
</style>
