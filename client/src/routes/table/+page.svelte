<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';

	import auth from '$stores/auth';
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
		created_at: true,
		first_choice: true,
		second_choice: true,
		introduction: true
	};

	onMount(async () => {
		const success = await auth.ping();
		if (!success) goto('/', { replaceState: true });
		else fetchApplicants();
	});
</script>

<svelte:head>
	<title>欢迎加入校科协 • 后台管理</title>
</svelte:head>

<main>
	<Header bind:tableFilter />
	<Table {tableFilter} />
	<Pagination />
</main>

<style lang="postcss">
	main {
		min-height: calc(100vh - 60px);
		min-height: calc(100svh - 60px);
		@apply py-5 md:py-10 px-4 md:px-20 flex flex-col gap-2;
	}
</style>
