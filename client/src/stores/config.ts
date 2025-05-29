import { writable } from 'svelte/store';

import fetchAppConfigApi from '$lib/app/config';
import type { AppConfigType } from '$types/app';

function createStore() {
	const store = writable<AppConfigType | null | undefined>(undefined);

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
