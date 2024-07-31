<script lang="ts">
	import type { ApplicantChoiceStatsType, GenderType } from '$types/applicant';

	import mapGender from '$lib/utils/mapGender';
	import Doughnut from '$components/Charts/Doughnut.svelte';

	export let title: string;
	export let data: ApplicantChoiceStatsType;

	function mapData(data: ApplicantChoiceStatsType) {
		return Object.entries(data).map(([label, value]) => ({ label: mapGender(label as GenderType), value }));
	}
</script>

<Doughnut data={mapData(data)} {title} />
