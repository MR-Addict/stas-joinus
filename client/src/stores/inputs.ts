import { get, writable } from 'svelte/store';

interface InputElement {
	id: string;
	ref: HTMLDivElement;
	value: string;
	error: string;
	validate: (vale: string) => string;
}

function createStore() {
	const store = writable<InputElement[]>([]);

	function register(id: string, ref: HTMLDivElement, validate: (vale: string) => string) {
		store.update((inputs) => {
			const inputIndex = inputs.findIndex((input) => input.id === id);
			if (inputIndex !== -1) inputs[inputIndex] = { id, ref, value: '', error: '', validate };
			else inputs.push({ id, ref, value: '', error: '', validate });
			return inputs;
		});
	}

	function update(id: string, value: string) {
		store.update((inputs) => {
			for (const input of inputs) {
				if (input.id === id) {
					input.value = value.toString();
					input.error = input.validate(input.value);
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
		if (hasError) scrollToFirstError();
		return hasError;
	}

	function scrollToFirstError() {
		get(inputs)
			.filter((input) => input.error !== '')
			.sort((a, b) => a.ref.offsetTop - b.ref.offsetTop)
			.at(0)
			?.ref.scrollIntoView({ behavior: 'smooth', block: 'center', inline: 'center' });
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
