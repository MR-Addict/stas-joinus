import z from 'zod';
import url from '$lib/utils/url';

import { AppConfig, type ApiResultType, type AppConfigType } from '$types/app';

export default async function fetchAppConfigApi(): Promise<ApiResultType<AppConfigType>> {
	const defaultMessage = '无法获取应用配置';

	try {
		const path = url('/api/app/config');
		const res = await fetch(path, { credentials: 'include' }).then((res) => res.json());
		const parsed = z.object({ data: AppConfig }).safeParse(res);
		if (parsed.success) return { success: true, data: parsed.data.data };
		else return { success: false, message: res.message || defaultMessage };
	} catch (err) {
		console.error(err);
		return { success: false, message: defaultMessage };
	}
}
