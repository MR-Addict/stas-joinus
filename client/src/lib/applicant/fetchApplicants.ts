import z from 'zod';
import url from '$lib/utils/url';
import { get } from 'svelte/store';

import applicants from '$stores/applicants';
import pagination from '$stores/pagination';
import { Applicant } from '$types/applicant';
import { Pagination } from '$types/pagination';

export default async function fetchApplicants(page?: number) {
	try {
		const pageData = get(pagination);
		const path = `/api/applicants?page=${page || pageData.page}&page_size=${pageData.page_size}`;
		const res = await fetch(url(path), { credentials: 'include' }).then((res) => res.json());
		const result = z.object({ data: z.array(Applicant), pagination: Pagination }).parse(res);
		applicants.set(result.data);
		pagination.set(result.pagination);
		return true;
	} catch (err) {
		console.error(err);
		applicants.set([]);
		return false;
	}
}
