<script lang="ts">
	import { onMount } from 'svelte';

	import form from '$stores/form';
	import config from '$stores/config';

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
	<title>参与报名 • 校大学生科协</title>
</svelte:head>

<main>
	{#if $config === undefined}
		<p>数据加载中，请稍后...</p>
	{:else if $config === null}
		<p>无法连接至服务器，请联系我们</p>
	{:else if $config}
		<TimeValidation config={$config}>
			{#if $form.showForm}
				<Form />
			{:else}
				<Success />
			{/if}
		</TimeValidation>
	{/if}
</main>

<style>
	main {
		@apply flex flex-col items-center justify-center px-4 pb-20 pt-5 sm:pt-10;
	}
</style>
