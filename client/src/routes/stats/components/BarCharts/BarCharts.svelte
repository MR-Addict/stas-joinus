<script lang="ts">
	import Bar from '$components/Charts/Bar.svelte';
	import type { ApplicantStatsType, ChoiceType } from '$types/applicant';
	import { onDestroy, onMount } from 'svelte';

	export let choice: ChoiceType;
	export let data: ApplicantStatsType[];

	let width: number = 400;
	let wrapperRef: HTMLDivElement;

	const resizeObserver = new ResizeObserver((entries) => (width = entries[0].contentRect.width));

	function mapData(data: ApplicantStatsType[], choice: ChoiceType) {
		return data.map((d) => ({ label: d.name, value: d[choice.label].boy + d[choice.label].girl }));
	}

	onMount(() => resizeObserver.observe(wrapperRef));
	onDestroy(() => resizeObserver.disconnect());
</script>

<div bind:this={wrapperRef} class="wrapper">
	<Bar
		data={mapData(data, choice)}
		title={`各个部门的人数对比（${choice.name}）`}
		size={{ width: Math.min(800, width), height: 400 }}
	/>
</div>

<style>
	.wrapper {
		@apply w-full flex flex-col items-center justify-center;
	}
</style>
