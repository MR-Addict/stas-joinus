<script lang="ts">
	import DoughnutChart from './DoughnutChart.svelte';
	import type { ApplicantStatsType } from '$types/applicant';

	export let toggleChoice: boolean;
	export let data: ApplicantStatsType[];

	const firstChoice = { name: '第一志愿', label: 'first_choice' } as const;
	const secondChoice = { name: '第二志愿', label: 'second_choice' } as const;

	type ChoiceType = typeof firstChoice | typeof secondChoice;

	let choice: ChoiceType = firstChoice;

	$: choice = toggleChoice ? firstChoice : secondChoice;
</script>

<div class="wrapper">
	<h1 class="text-center text-lg font-semibold">各个部门的总人数及男女分布</h1>

	<div class="doughnuts">
		{#each data as d (d.name)}
			<DoughnutChart data={d[choice.label]} title={`${d.name}（${choice.name}）`} />
		{/each}
	</div>
</div>

<style>
	.wrapper {
		@apply w-full flex flex-col items-center justify-center gap-5 isolate;
	}

	.doughnuts {
		@apply w-full grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4 place-items-center gap-10;
	}
</style>
