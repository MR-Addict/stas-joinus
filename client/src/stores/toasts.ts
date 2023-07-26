import { writable, derived, get } from 'svelte/store';

import type { ToastStatusType, ToastType } from '$types/toast';

function uuid() {
	return Math.random().toString(36).slice(-9);
}

function createStore() {
	const store = writable<ToastType[]>([]);

	const derivedStore = derived(store, ($store) => ({
		all: $store,
		finished: $store.filter((item) => item.status !== 'pending')
	}));

	function add(message: string, status: ToastStatusType) {
		store.update((items) => [...items, { id: uuid(), message, status }]);
	}

	function update(id: string, callback: (toast: ToastType) => ToastType) {
		const toast = get(store).find((item) => item.id === id);
		if (toast) {
			const updatedToast = callback(toast);
			store.update((items) => {
				items[items.findIndex((item) => item.id === updatedToast.id)] = updatedToast;
				return items;
			});
		}
	}

	function remove(id: string) {
		store.update((items) => items.filter((item) => item.id !== id));
	}

	return {
		add: add,
		update: update,
		remove: remove,
		subscribe: derivedStore.subscribe
	};
}

export const toasts = createStore();
