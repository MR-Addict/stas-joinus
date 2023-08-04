import z from 'zod';

import url from '$lib/utils/url';
import toasts from '$stores/toasts';

export default async function submitForm(formData: FormData) {
	const toastId = toasts.add('表单提交中，请耐心等待', 'pending');
	try {
		const res = await fetch(url('/api/applicant'), { method: 'POST', body: formData }).then((res) => res.json());
		const { message, success } = z.object({ success: z.boolean(), message: z.string() }).parse(res);
		toasts.update(toastId, { message, status: success ? 'success' : 'failed' });
		return success;
	} catch (err) {
		console.error(err);
		toasts.update(toastId, { message: '提交失败，请稍后重试或联系我们', status: 'failed' });
		return false;
	}
}
