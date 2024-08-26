import { get, writable } from 'svelte/store';
import type { ApplicantType } from '$types/applicant';
import type { PaginationType } from '$types/pagination';
import fetchApplicantsApi from '$lib/applicant/fetchApplicants';

function createStore() {
	const store = writable<{ applicants: ApplicantType[]; pagination: PaginationType } | null>(null);

	async function refersh(page: number = 1) {
		const res = await fetchApplicantsApi(page, get(store)?.pagination.page_size || 50);
		if (res.success) store.set(res.data);
		return res;
	}

	return {
		refersh,
		subscribe: store.subscribe
	};
}

const view = createStore();

export default view;
