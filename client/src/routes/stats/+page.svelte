<script lang="ts">
	import { onMount } from 'svelte';
	import stats from '$stores/stats';

	import DoughnutChart from './components/DoughnutChart.svelte';

	onMount(stats.refersh);
</script>

<svelte:head>
	<title>报名统计 • 校大学生科协</title>
</svelte:head>

<main>
	{#if $stats}
		<DoughnutChart data={$stats} />
	{:else}
		<p>数据加载中，请稍后...</p>
	{/if}
</main>

<style>
	main {
		@apply px-4 py-5 lg:px-20 sm:py-10 flex flex-col items-center justify-center;
	}
</style>
