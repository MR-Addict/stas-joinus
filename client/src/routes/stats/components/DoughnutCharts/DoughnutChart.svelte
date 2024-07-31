<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import { ArrowDownToLine } from 'lucide-svelte';
	import type { ApplicantChoiceStatsType, GenderType } from '$types/applicant';

	import mapGender from '$lib/utils/mapGender';
	import downloadSvg from '$lib/utils/downloadSvg';
	import Doughnut from '$components/Charts/Doughnut.svelte';

	export let title: string;
	export let data: ApplicantChoiceStatsType;

	let svgRef: SVGElement;
	let wrapperRef: HTMLDivElement;
	let dimensions: { width: number; height: number };

	function mapData(data: ApplicantChoiceStatsType) {
		return Object.entries(data).map(([label, value]) => ({ label: mapGender(label as GenderType), value }));
	}

	const observer = new ResizeObserver((entries) => {
		for (let entry of entries) {
			const { width, height } = entry.contentRect;
			dimensions = { width, height };
		}
	});

	onMount(() => observer.observe(wrapperRef));
	onDestroy(() => observer.disconnect());
</script>

<div bind:this={wrapperRef}>
	<Doughnut data={mapData(data)} {title} bind:ref={svgRef} size={dimensions} />

	<button type="button" on:click={() => downloadSvg(svgRef, title)} aria-label="download">
		<ArrowDownToLine size={16} />
	</button>
</div>

<style>
	div {
		height: 420px;
		@apply w-full flex items-center justify-center relative shadow-md rounded-xl p-5 border border-gray-300;
	}

	button {
		@apply absolute right-5 bottom-5 p-1 text-gray-700 rounded-full bg-black/10 duration-300;

		&:hover {
			@apply bg-black/30;
		}
	}
</style>
