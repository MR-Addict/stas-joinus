<script lang="ts">
	import type { ApplicantStatsType } from '$types/applicant';

	import DoughnutChart from './DoughnutChart.svelte';

	export let data: ApplicantStatsType[];

	const firstChoice = { name: '第一志愿', label: 'first_choice' } as const;
	const secondChoice = { name: '第二志愿', label: 'second_choice' } as const;

	type ChoiceType = typeof firstChoice | typeof secondChoice;

	let toggleChoice = true;
	let choice: ChoiceType = firstChoice;

	$: choice = toggleChoice ? firstChoice : secondChoice;
</script>

<div class="wrapper">
	<h1 class="text-center text-lg font-semibold">各个部门的总人数及男女分布</h1>

	<button type="button" on:click={() => (toggleChoice = !toggleChoice)}>
		<span class="bubble" class:active={toggleChoice} />
		<span class="option" class:active={toggleChoice}>第一志愿</span>
		<span class="option" class:active={!toggleChoice}>第二志愿</span>
	</button>

	<div class="doughnuts">
		{#each data as d (d.name)}
			<DoughnutChart data={d[choice.label]} title={`${d.name}（${choice.name}）`} />
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
</style>
