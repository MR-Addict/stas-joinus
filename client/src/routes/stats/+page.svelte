<script lang="ts">
	import { onMount } from 'svelte';

	import stats from '$stores/stats';
	import { firstChoice } from '$data/choice';
	import type { ChoiceType } from '$types/applicant';
	
	import BarCharts from './components/BarCharts/BarCharts.svelte';
	import ToggleChoice from './components/ToggleChoice/ToggleChoice.svelte';
	import DoughnutCharts from './components/DoughnutCharts/DoughnutCharts.svelte';

	let choice: ChoiceType = firstChoice;

	onMount(stats.refersh);
</script>

<svelte:head>
	<title>报名统计 • 校大学生科协</title>
</svelte:head>

<main>
	{#if $stats}
		<ToggleChoice bind:choice />
		<BarCharts data={$stats} {choice} />
		<hr />
		<DoughnutCharts data={$stats} {choice} />
	{:else}
		<p>数据加载中，请稍后...</p>
	{/if}
</main>

<style>
	main {
		@apply px-4 py-5 lg:px-20 sm:py-10 flex flex-col items-center justify-center gap-5;
	}
	hr {
		@apply w-full border-t border-gray-300 my-5;
	}
</style>
