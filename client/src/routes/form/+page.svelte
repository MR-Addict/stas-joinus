<script lang="ts">
	import Form from './components/Form/Form.svelte';
</script>

<svelte:head>
	<title>欢迎加入校科协 • 参与报名</title>
</svelte:head>

<main>
	<Form />
</main>

<style lang="postcss">
	main {
		min-height: calc(100vh - 60px);
		min-height: calc(100svh - 60px);
		@apply bg-gray-100 flex flex-col items-center justify-center px-4 py-5 md:py-10;
	}
</style>
