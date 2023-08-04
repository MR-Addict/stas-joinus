import type { ApplicantType } from '$types/applicant';
import { writable } from 'svelte/store';

function createStore() {
	const store = writable<ApplicantType[]>([]);

	return {
		set: store.set,
		subscribe: store.subscribe
	};
}

const applicants = createStore();

export default applicants;
