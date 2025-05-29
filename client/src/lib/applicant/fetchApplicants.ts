import z from 'zod';
import url from '$lib/utils/url';

import type { ApiResultType } from '$types/app';
import { Applicant, type ApplicantType } from '$types/applicant';
import { Pagination, type PaginationType } from '$types/pagination';

type Props = (Omit<PaginationType, 'total'> & { all: false }) | { all: true };
type Return = Promise<ApiResultType<{ applicants: ApplicantType[]; pagination: PaginationType }>>;

export default async function fetchApplicantsApi(props: Props): Return {
	const defaultMessage = '报名信息获取失败';

	try {
		const searchParams = new URLSearchParams();
		if (!props.all) {
			searchParams.set('page_size', props.page_size.toString());
			if (props.query) searchParams.set('query', props.query);
			if (props.page > 1) searchParams.set('page', props.page.toString());
		} else searchParams.set('all', 'true');

		const path = `/api/applicants?${searchParams.toString()}`;
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
