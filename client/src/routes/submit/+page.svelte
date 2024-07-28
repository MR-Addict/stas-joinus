<script lang="ts">
	import Form from './components/Form/Form.svelte';
	import Success from './components/Success/Success.svelte';
	import TimeValidation from './components/TimeValidation/TimeValidation.svelte';

	let submitted = false;
</script>

<svelte:head>
	<title>校大学生科协 • 参与报名</title>
</svelte:head>

<main>
	<TimeValidation>
		{#if submitted}
			<Success />
		{:else}
			<Form bind:submitted />
		{/if}
	</TimeValidation>
</main>

<style lang="postcss">
	main {
		@apply flex-1 flex flex-col items-center justify-center px-4 py-5 sm:py-10;
	}
</style>
