import z from 'zod';
import url from '$lib/utils/url';

import type { ApiResultType } from '$types/app';

export default async function userPingApi(): Promise<ApiResultType> {
	const defaultMessage = '无法获取管理员信息';

	try {
		const path = url('/api/user');
		const res = await fetch(path, { credentials: 'include' }).then((res) => res.json());
		const parsed = z.object({ success: z.boolean(), message: z.string() }).safeParse(res);
		if (parsed.success && parsed.data.success) return { success: true, message: parsed.data.message };
		else return { success: false, message: res.message || defaultMessage };
	} catch (err) {
		console.error(err);
		return { success: false, message: defaultMessage };
	}
}
