import { writable } from 'svelte/store';

import type { AppConfigType } from '$types/app';
import fetchAppConfigApi from '$lib/app/config';

function createStore() {
	const store = writable<AppConfigType | null>(null);

	async function ping() {
		const res = await fetchAppConfigApi();
		if (res.success) store.set(res.data);
		else store.set(null);
		return res;
	}

	return {
		ping: ping,
		subscribe: store.subscribe
	};
}

const config = createStore();

export default config;
