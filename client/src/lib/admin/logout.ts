import z from 'zod';
import url from '$lib/utils/url';

import type { ApiResultType } from '$types/app';

export default async function userLogoutApi(): Promise<ApiResultType> {
	const defaultMessage = '退出登录失败';

	try {
		const path = url('/api/user/logout');
		const res = await fetch(path, { credentials: 'include' }).then((res) => res.json());
		const parsed = z.object({ success: z.boolean(), message: z.string() }).safeParse(res);
		if (parsed.success && parsed.data.success) return { success: true, message: parsed.data.message };
		else return { success: false, message: res.message || defaultMessage };
	} catch (err) {
		console.error(err);
		return { success: false, message: defaultMessage };
	}
}
