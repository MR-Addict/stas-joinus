import z from 'zod';
import url from '$lib/utils/url';

import type { ApiResultType } from '$types/app';

export default async function userLoginApi(formData: FormData): Promise<ApiResultType> {
	const defaultMessage = '登录失败';

	try {
		const path = url('/api/user/login');
		const res = await fetch(path, { method: 'POST', body: formData, credentials: 'include' }).then((res) => res.json());
		const parsed = z.object({ success: z.boolean() }).safeParse(res);
		if (parsed.success) return { success: true };
		else return { success: false, message: res.message || defaultMessage };
	} catch (err) {
		console.error(err);
		return { success: false, message: defaultMessage };
	}
}
