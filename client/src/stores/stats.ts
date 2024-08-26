import { writable } from 'svelte/store';
import fetchStats from '$lib/applicant/fetchStats';
import type { ApplicantStatsType } from '$types/applicant';

function createStore() {
	const store = writable<ApplicantStatsType[] | null>(null);

	async function refersh() {
		const res = await fetchStats();
		if (res.success) store.set(res.data);
		return res;
	}

	return {
		refersh,
		subscribe: store.subscribe
	};
}

const stats = createStore();

export default stats;
