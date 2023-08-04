import { writable } from 'svelte/store';
import type { ApplicantType } from '$types/applicant';

function createStore() {
	const store = writable<ApplicantType[]>([]);

	return {
		set: store.set,
		subscribe: store.subscribe
	};
}

const applicants = createStore();

export default applicants;
