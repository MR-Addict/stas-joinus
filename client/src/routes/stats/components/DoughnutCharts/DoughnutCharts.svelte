<script lang="ts">
	import Doughnut from '$components/Charts/Doughnut.svelte';

	import mapGender from '$lib/utils/mapGender';
	import type { ApplicantStatsType, ApplicantChoiceStatsType, ChoiceType } from '$types/applicant';

	export let choice: ChoiceType;
	export let data: ApplicantStatsType[];

	function mapData(data: ApplicantChoiceStatsType) {
		return Object.entries(data).map(([label, value]) => ({ label: mapGender(label), value }));
	}
</script>

<div class="wrapper">
	<h1 class="text-center text-lg font-semibold">各部门总人数及男女分布</h1>

	<div class="doughnuts">
		{#each data as d (d.name)}
			<Doughnut data={mapData(d[choice.label])} title={`${d.name}（${choice.name}）`} />
		{/each}
	</div>
</div>

<style>
	.wrapper {
		@apply w-full flex flex-col items-center justify-center gap-5;
	}

	.doughnuts {
		@apply w-full grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4 place-items-center gap-10;
	}
</style>
