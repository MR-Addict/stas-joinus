import z from 'zod';

import { toasts } from '$stores/toasts';

export default async function submitForm(formData: FormData) {
	try {
		const res = await fetch('https://joinus.mraddict.top/api/applicant', { method: 'POST', body: formData }).then(
			(res) => res.json()
		);
		const result = z.object({ success: z.boolean(), message: z.string() }).parse(res);
		if (!result.success) toasts.add(result.message, 'failed');
		return result.success;
	} catch (err) {
		console.error(err);
		toasts.add('提交失败，未知错误，请稍后重试或联系我们', 'failed');
		return false;
	}
}
