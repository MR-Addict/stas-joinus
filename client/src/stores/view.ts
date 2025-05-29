import toast from 'svelte-french-toast';
import { get, writable } from 'svelte/store';

import type { ApplicantType } from '$types/applicant';
import type { PaginationType } from '$types/pagination';
import fetchApplicantsApi from '$lib/applicant/fetchApplicants';

function createStore() {
	const store = writable<{ applicants: ApplicantType[]; pagination: PaginationType } | null>(null);

	async function refersh(page?: number, page_size?: number, query?: string) {
		const storeData = get(store);
		if (page === undefined) page = storeData?.pagination.page || 1;
		if (page_size === undefined) page_size = storeData?.pagination.page_size || 50;
		if (query === undefined) query = storeData?.pagination.query || '';

		const res = await fetchApplicantsApi({ page, page_size, query, all: false });
		if (res.success) store.set(res.data);
		else toast.error(res.message);
		return res;
	}

	return {
		refersh,
		subscribe: store.subscribe
	};
}

const view = createStore();

export default view;
