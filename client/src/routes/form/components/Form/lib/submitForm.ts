import z from 'zod';

import url from '$lib/utils/url';
import toast from 'svelte-french-toast';

export default async function submitForm(formData: FormData) {
	const id = toast('表单提交中，请稍后...');
	try {
		const res = await fetch(url('/api/applicant'), { method: 'POST', body: formData }).then((res) => res.json());
		const { message, success } = z.object({ success: z.boolean(), message: z.string() }).parse(res);
		if (success) toast.success(message, { id });
		else toast.error(message, { id });
		return success;
	} catch (err) {
		console.error(err);
		toast.error('提交失败，请稍后重试或联系我们', { id });
		return false;
	}
}
