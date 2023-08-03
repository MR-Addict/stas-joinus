import { get, writable } from 'svelte/store';

interface InputElement {
	id: string;
	value: string;
	error: string;
	validate: (vale: string) => string;
}

function createStore() {
	const store = writable<InputElement[]>([]);

	function register(id: string, validate: (vale: string) => string) {
		if (get(store).find((input) => input.id === id)) return;
		store.update((inputs) => inputs.concat({ id, value: '', error: '', validate }));
	}

	function update(id: string, value: string) {
		store.update((inputs) => {
			for (const input of inputs) {
				if (input.id === id) {
					input.value = value;
					input.error = input.validate(value);
				}
			}
			return inputs;
		});
	}

	function validateAll() {
		let hasError = false;
		store.update((inputs) => {
			for (const input of inputs) {
				input.error = input.validate(input.value);
				if (input.error) hasError = true;
			}
			return inputs;
		});
		return hasError;
	}

	return {
		update,
		register,
		validateAll,
		subscribe: store.subscribe
	};
}

const inputs = createStore();

export default inputs;
