import { writable } from 'svelte/store';

import userPingApi from '$lib/admin/ping';
import userLoginApi from '$lib/admin/login';
import userLogoutApi from '$lib/admin/logout';

function createStore() {
	const authorized = writable(false);

	async function login(formData: FormData) {
		const res = await userLoginApi(formData);
		authorized.set(res.success);
		return res;
	}

	async function ping() {
		const res = await userPingApi();
		authorized.set(res.success);
		return res;
	}

	async function logout() {
		const res = await userLogoutApi();
		authorized.set(!res.success);
		return res;
	}

	return {
		ping,
		login,
		logout,
		subscribe: authorized.subscribe
	};
}

const auth = createStore();

export default auth;
