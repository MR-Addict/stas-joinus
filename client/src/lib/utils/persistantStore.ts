import { writable } from 'svelte/store';

export default function persistantStore<T>(key: string, defaultValue: T, prefix = 'persistant-store') {
	const store = writable<T>(defaultValue);

	if (typeof window === 'undefined') return store;

	const json = localStorage.getItem(`${prefix}-${key}`);
	if (json) store.set(JSON.parse(json));

	store.subscribe((value) => localStorage.setItem(`${prefix}-${key}`, JSON.stringify(value)));

	return store;
}
