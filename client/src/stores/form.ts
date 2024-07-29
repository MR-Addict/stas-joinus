import { writable } from 'svelte/store';
import { Applicant } from '$types/applicant';
import type { ApplicantType } from '$types/applicant';

function createStore() {
	const form = writable<{ showForm: boolean; localApplicant: ApplicantType | null }>({
		showForm: true,
		localApplicant: null
	});

	function refreshLocalApplicant() {
		const local = localStorage.getItem('applicant');
		if (!local) {
			form.update((prev) => ({ ...prev, localApplicant: null }));
			return null;
		}

		const parsed = Applicant.safeParse(JSON.parse(local));
		if (!parsed.success) {
			form.update((prev) => ({ ...prev, localApplicant: null }));
			return null;
		}

		form.update((prev) => ({ ...prev, localApplicant: parsed.data }));
		return parsed.data;
	}

	function updateShowForm(showForm: boolean) {
		form.update((state) => ({ ...state, showForm }));
	}

	return {
		updateShowForm,
		refreshLocalApplicant,
		subscribe: form.subscribe
	};
}

const form = createStore();

export default form;
