<script lang="ts">
	import inputs from '$stores/inputs';
	import toast from 'svelte-french-toast';
	import submitApplicantApi from '$lib/applicant/submitApplicant';

	import Header from './components/Header/Header.svelte';
	import SubmitButton from './components/SubmitButton/SubmitButton.svelte';

	import NameInput from './components/NameInput/NameInput.svelte';
	import GenderSelect from './components/GenderSelect/GenderSelect.svelte';
	import PhoneInput from './components/PhoneInput/PhoneInput.svelte';
	import EmailInput from './components/EmailInput/EmailInput.svelte';
	import QQInput from './components/QQInput/QQInput.svelte';
	import StudentIdInput from './components/StudentIDInput/StudentIDInput.svelte';
	import CollegeInput from './components/CollegeInput/CollegeInput.svelte';
	import MajorInput from './components/MajorInput/MajorInput.svelte';
	import ChoiceSelect from './components/ChoiceSelect/ChoiceSelect.svelte';
	import IntroductionInput from './components/IntroductionInput/IntroductionInput.svelte';

	export let submitted: boolean;

	let pending = false;

	async function handleSubmit(event: SubmitEvent) {
		if (inputs.validateAll()) return;

		pending = true;

		const res = await submitApplicantApi(new FormData(event.target as HTMLFormElement));
		if (res.success) {
			submitted = true;
			window.scrollTo({ top: 0 });
			localStorage.setItem('applicant', JSON.stringify(res.data));
		} else toast.error(res.message);

		pending = false;
	}
</script>

<form on:submit|preventDefault={handleSubmit}>
	<Header />

	<NameInput />
	<GenderSelect />
	<PhoneInput />
	<EmailInput />
	<QQInput />
	<StudentIdInput />
	<CollegeInput />
	<MajorInput />
	<ChoiceSelect />
	<IntroductionInput />

	<SubmitButton {pending} />
</form>

<style lang="postcss">
	form {
		@apply w-full max-w-xl flex flex-col gap-7;
	}
</style>
