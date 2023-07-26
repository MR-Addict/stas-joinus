<script>
	import FormIcon from '$components/Icons/FormIcon/FormIcon.svelte';
</script>

<header>
	<FormIcon />
	<h1>校科协报名表单</h1>
</header>

<style lang="postcss">
	header {
		@apply w-full bg-white p-5 rounded-md flex flex-row items-center justify-center gap-2 border border-gray-300;
	}
	h1 {
		@apply text-2xl font-semibold text-gray-800;
	}
</style>
