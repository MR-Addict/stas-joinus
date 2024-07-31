<script lang="ts">
	import mapGender from '$lib/utils/mapGender';
	import { ArrowDownToLine } from 'lucide-svelte';
	import type { ApplicantChoiceStatsType, GenderType } from '$types/applicant';

	import downloadSvg from '$lib/utils/downloadSvg';
	import Doughnut from '$components/Charts/Doughnut.svelte';

	export let title: string;
	export let data: ApplicantChoiceStatsType;

	let ref: SVGElement;

	function mapData(data: ApplicantChoiceStatsType) {
		return Object.entries(data).map(([label, value]) => ({ label: mapGender(label as GenderType), value }));
	}
</script>

<div>
	<Doughnut data={mapData(data)} {title} bind:ref />

	<button type="button" on:click={() => downloadSvg(ref, title)}>
		<ArrowDownToLine size={16} />
	</button>
</div>

<style>
	div {
		@apply relative shadow-md rounded-xl p-5 border border-gray-300;
	}

	button {
		@apply absolute right-5 bottom-5 p-1 text-gray-700 rounded-full bg-black/10 duration-300;

		&:hover {
			@apply bg-black/30;
		}
	}
</style>
