import z from 'zod';
import url from '$lib/utils/url';

import type { ApiResultType } from '$types/app';
import { Applicant, type ApplicantType } from '$types/applicant';
import { Pagination, type PaginationType } from '$types/pagination';

export default async function fetchApplicantsApi(
	page: number = 1,
	pageSize: number = 20
): Promise<ApiResultType<{ applicants: ApplicantType[]; pagination: PaginationType }>> {
	const defaultMessage = '报名信息获取失败';

	try {
		const path = `/api/applicants?page=${page}&page_size=${pageSize}`;
		const res = await fetch(url(path), { credentials: 'include' }).then((res) => res.json());

		const parsed = z.object({ data: z.array(Applicant), pagination: Pagination, message: z.string() }).safeParse(res);
		if (parsed.success)
			return {
				success: true,
				message: parsed.data.message,
				data: { applicants: parsed.data.data, pagination: parsed.data.pagination }
			};
		else return { success: false, message: res.message || defaultMessage };
	} catch (err) {
		console.error(err);
		return { success: false, message: defaultMessage };
	}
}
