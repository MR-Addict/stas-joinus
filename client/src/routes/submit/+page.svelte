<script lang="ts">
	import { onMount } from 'svelte';

	import form from '$stores/form';

	import Form from './components/Form/Form.svelte';
	import Success from './components/Success/Success.svelte';
	import TimeValidation from './components/TimeValidation/TimeValidation.svelte';

	onMount(() => {
		const localApplicant = form.refreshLocalApplicant();
		if (localApplicant) form.updateShowForm(false);
		else form.updateShowForm(true);
	});
</script>

<svelte:head>
	<title>校大学生科协 • 参与报名</title>
</svelte:head>

<main>
	<TimeValidation>
		{#if $form.showForm}
			<Form />
		{:else}
			<Success />
		{/if}
	</TimeValidation>
</main>

<style lang="postcss">
	main {
		@apply flex flex-col items-center justify-center px-4 py-5 sm:py-10;
	}
</style>
