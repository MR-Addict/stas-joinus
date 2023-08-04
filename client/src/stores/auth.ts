import url from '$lib/utils/url';
import { writable } from 'svelte/store';
import { Response } from '$types/response';

function createStore() {
	const authorized = writable(false);

	async function ping() {
		try {
			const res = await fetch(url('/api/user'), { credentials: 'include' }).then((res) => res.json());
			const { success } = Response.parse(res);
			authorized.set(success);
			return success;
		} catch (err) {
			console.error(err);
			authorized.set(false);
			return false;
		}
	}

	async function logout() {
		try {
			await fetch(url('/api/user/logout'));
			const success = await ping();
			if (!success) authorized.set(false);
			return !success;
		} catch (err) {
			console.error(err);
			return false;
		}
	}

	return {
		ping,
		logout,
		subscribe: authorized.subscribe
	};
}

const auth = createStore();

export default auth;
