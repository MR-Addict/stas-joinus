<script lang="ts">
	import { onMount } from 'svelte';
	import stats from '$stores/stats';

	import ToggleChoice from './components/ToggleChoice/ToggleChoice.svelte';
	import DoughnutCharts from './components/DoughnutCharts/DoughnutCharts.svelte';

	let toggleChoice = true;

	onMount(stats.refersh);
</script>

<svelte:head>
	<title>报名统计 • 校大学生科协</title>
</svelte:head>

<main>
	{#if $stats}
		<ToggleChoice bind:toggleChoice />
		<DoughnutCharts data={$stats} {toggleChoice} />
	{:else}
		<p>数据加载中，请稍后...</p>
	{/if}
</main>

<style>
	main {
		@apply px-4 py-5 lg:px-20 sm:py-10 flex flex-col items-center gap-3;
	}
</style>
