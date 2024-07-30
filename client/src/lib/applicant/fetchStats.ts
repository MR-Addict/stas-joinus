import z from 'zod';

import url from '$lib/utils/url';
import type { ApiResultType } from '$types/app';
import { ApplicantStats, type ApplicantStatsType } from '$types/applicant';

export default async function fetchStats(): Promise<ApiResultType<ApplicantStatsType[]>> {
	const defaultMessage = '统计信息获取失败';

	try {
		const path = url('/api/applicants/stats');
		const res = await fetch(path, { credentials: 'include' }).then((res) => res.json());
		const parsed = z.object({ message: z.string(), data: z.array(ApplicantStats) }).safeParse(res);
		if (!parsed.success) return { success: false, message: res.message || defaultMessage };

		const data = parsed.data.data.sort((a, b) => b.name.localeCompare(a.name));
		return { success: true, message: parsed.data.message, data };
	} catch (err) {
		console.error(err);
		return { success: false, message: defaultMessage };
	}
}
