<script lang="ts">
	import inputs from '$stores/inputs';
	import submitForm from './lib/submitForm';
	import scrollToFirstError from './lib/scrollToFirstError';

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

	export let openForm: (value: boolean) => void;

	let pending = false;

	async function handleSubmit(event: SubmitEvent) {
		if (inputs.validateAll()) scrollToFirstError();
		else {
			pending = true;
			const success = await submitForm(new FormData(event.target as HTMLFormElement));
			if (success) {
				openForm(false);
				window.scrollTo({ top: 0 });
			}
			pending = false;
		}
	}
</script>

<form on:submit|preventDefault={handleSubmit} class="w-full max-w-2xl flex flex-col items-end gap-5">
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
