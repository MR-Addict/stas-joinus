import { writable, derived } from 'svelte/store';

import type { ToastStatusType, ToastType } from '$types/toast';

function uuid() {
	return Math.random().toString(36).slice(-9);
}

function createStore() {
	const store = writable<ToastType[]>([]);

	const derivedStore = derived(store, ($store) => ({
		all: $store,
		finished: $store.filter((toast) => toast.status !== 'pending')
	}));

	function add(message: string, status: ToastStatusType) {
		const id = uuid();
		store.update((toasts) => toasts.concat({ id, message, status }));
		return id;
	}

	function update(id: string, update: Partial<ToastType>) {
		store.update((toasts) => {
			const toastId = toasts.findIndex((toast) => toast.id === id);
			if (toastId !== -1) Object.assign(toasts[toastId], update);
			return toasts;
		});
	}

	function remove(id: string) {
		store.update((toasts) => toasts.filter((toast) => toast.id !== id));
	}

	return {
		add,
		update,
		remove,
		subscribe: derivedStore.subscribe
	};
}

const toasts = createStore();

export default toasts;
