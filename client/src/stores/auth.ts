import { writable } from 'svelte/store';

import userPingApi from '$lib/admin/ping';
import userLoginApi from '$lib/admin/login';
import userLogoutApi from '$lib/admin/logout';

function createStore() {
	const authorized = writable(false);

	function open(){
		const loginButton = document.querySelector('button[aria-label="login button"]')
		if(loginButton) (loginButton as HTMLButtonElement).click();
	}

	async function ping() {
		const res = await userPingApi();
		authorized.set(res.success);
		return res;
	}

	async function login(formData: FormData) {
		const loginRes = await userLoginApi(formData);
		if (!loginRes.success) {
			authorized.set(false);
			return loginRes;
		}

		const pingRes = await userPingApi();
		if (!pingRes.success) {
			authorized.set(false);
			return pingRes;
		}

		authorized.set(true);
		return loginRes;
	}

	async function logout() {
		const logoutRes = await userLogoutApi();
		if (!logoutRes.success) {
			authorized.set(true);
			return logoutRes;
		}

		const pingRes = await userPingApi();
		if (pingRes.success) {
			authorized.set(true);
			return { success: false, message: '无法退出登录' };
		}

		authorized.set(false);
		return logoutRes;
	}

	return {
		open,
		ping,
		login,
		logout,
		subscribe: authorized.subscribe
	};
}

const auth = createStore();

export default auth;
