import type { PaginationType } from '$types/pagination';
import { writable } from 'svelte/store';

function createStore() {
	const store = writable<PaginationType>({ page: 1, total: 0, page_size: 20 });

	return {
		set: store.set,
		subscribe: store.subscribe
	};
}

const pagination = createStore();

export default pagination;
