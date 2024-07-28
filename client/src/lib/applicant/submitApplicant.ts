import z from 'zod';

import url from '$lib/utils/url';
import type { ApiResultType } from '$types/app';
import { Applicant, type ApplicantType } from '$types/applicant';

export default async function submitApplicantApi(formData: FormData): Promise<ApiResultType<ApplicantType>> {
	const defaultMessage = '提交失败，请稍后重试或联系我们';

	try {
		const path = url('/api/applicant');
		const res = await fetch(path, { method: 'POST', body: formData }).then((res) => res.json());
		const parsed = z.object({ data: Applicant }).safeParse(res);
		if (parsed.success) return { success: true, data: parsed.data.data };
		else return { success: false, message: res.message || defaultMessage };
	} catch (err) {
		console.error(err);
		return { success: false, message: defaultMessage };
	}
}
