<script lang="ts">
	import toast from 'svelte-french-toast';

	import form from '$stores/form';
	import inputs from '$stores/inputs';
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

	let pending = false;

	async function handleSubmit(event: SubmitEvent) {
		pending = true;

		const formData = new FormData();
		$inputs.forEach((input) => formData.append(input.id, input.value));

		const res = await submitApplicantApi(formData);
		if (res.success) {
			toast.success(res.message);
			window.scrollTo({ top: 0 });
			localStorage.setItem('applicant', JSON.stringify(res.data));

			form.updateShowForm(false);
			form.refreshLocalApplicant();
		} else toast.error(res.message);

		pending = false;
	}
</script>

<form on:submit|preventDefault={handleSubmit}>
	<Header />

	<div class="group">
		<p>个人信息</p>
		<NameInput />
		<GenderSelect />
	</div>

	<div class="group">
		<p>联系方式</p>
		<PhoneInput />
		<EmailInput />
		<QQInput />
	</div>

	<div class="group">
		<p>学院及专业</p>
		<StudentIdInput />
		<CollegeInput />
		<MajorInput />
	</div>

	<div class="group">
		<p>选择志愿</p>
		<ChoiceSelect />
	</div>

	<div class="group">
		<p>自我介绍</p>
		<IntroductionInput />
	</div>

	<SubmitButton {pending} />
</form>

<style>
	form {
		@apply w-full max-w-xl flex flex-col gap-6;
	}

	div {
		@apply flex flex-col gap-6;

		& > p {
			@apply mx-auto text-xl font-bold;
		}
	}
</style>
