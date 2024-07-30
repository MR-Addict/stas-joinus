<script lang="ts">
	import mapGender from '$lib/utils/mapGender';
	import type { ApplicantChoiceStatsType, ApplicantStatsType, GenderType } from '$types/applicant';

	import Doughnut from '$components/Charts/Doughnut.svelte';

	export let data: ApplicantStatsType[];

	let firstChoice = true;

	let choice:
		| {
				name: '第一志愿';
				label: 'first_choice';
		  }
		| {
				name: '第二志愿';
				label: 'second_choice';
		  } = {
		name: '第一志愿',
		label: 'first_choice'
	};

	$: choice = firstChoice
		? {
				name: '第一志愿',
				label: 'first_choice'
			}
		: {
				name: '第二志愿',
				label: 'second_choice'
			};

	function mapData(data: ApplicantChoiceStatsType) {
		return Object.entries(data).map(([label, value]) => ({ label: mapGender(label as GenderType), value }));
	}
</script>

<div class="wrapper">
	<h1 class="text-center text-lg font-semibold">各个部门的总人数及男女分布</h1>

	<button type="button" on:click={() => (firstChoice = !firstChoice)}>
		<span class="bubble" class:active={firstChoice} />
		<span class="option" class:active={firstChoice}>第一志愿</span>
		<span class="option" class:active={!firstChoice}>第二志愿</span>
	</button>

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
	button {
		@apply relative flex flex-row items-center gap-2 rounded-md py-2 px-2 bg-black/10 isolate;

		& .option {
			@apply py-1 px-2 rounded-md duration-300;
		}

		& .bubble {
			--gap: 0.25rem;
			--py: 0.45rem;
			--px: 0.5rem;
			top: var(--py);
			right: var(--px);
			height: calc(100% - var(--py) * 2);
			width: calc(50% - var(--px) - var(--gap));
			@apply absolute bg-black/20 -z-10 rounded-md duration-500 ease-in-out;

			&.active {
				transform: translateX(calc(-100% - var(--px)));
			}
		}
	}

	.doughnuts {
		@apply w-full grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 2xl:grid-cols-4 place-items-center gap-10;
	}
</style>
