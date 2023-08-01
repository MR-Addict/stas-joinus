import { get, writable } from 'svelte/store';

interface ErrorElement {
	id: string;
	error: string;
}

function createStore() {
	const store = writable<ErrorElement[]>([]);

	function getElement(id: string) {
		for (const element of get(store)) {
			if (element.id === id) return true;
		}
		return false;
	}

	function upsertElement(id: string, error: string) {
		if (!getElement(id)) store.update((elements) => elements.concat({ id, error }));
		else {
			store.update((elements) => {
				for (const element of elements) if (element.id === id) element.error = error;
				return elements;
			});
		}
	}

	function removeElement(id: string) {
		store.update((elements) => elements.filter((element) => element.id !== id));
	}

	return {
		upsert: upsertElement,
		remove: removeElement,
		subscribe: store.subscribe
	};
}

const errorElements = createStore();

export default errorElements;
