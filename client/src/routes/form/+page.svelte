<script lang="ts">
	import Form from './components/Form/Form.svelte';
	import Success from './components/Success/Success.svelte';

	let isFormOpen = true;
	const endDate = new Date('2023-09-16T00:00:00.000Z');
	const openForm = (value: boolean) => (isFormOpen = value);
</script>

<svelte:head>
	<title>校大学生科协 • 参与报名</title>
</svelte:head>

<main>
	{#if new Date() > endDate}
		<p class="text-center">报名时间已截至，如果你还想报名，可以前往首页加群联系我们</p>
	{:else if isFormOpen}
		<Form {openForm} />
	{:else}
		<Success />
	{/if}
</main>

<style lang="postcss">
	main {
		min-height: calc(100vh - 60px);
		min-height: calc(100svh - 60px);
		@apply bg-gray-100 flex flex-col items-center justify-center px-4 py-5 md:py-10;
	}
</style>
