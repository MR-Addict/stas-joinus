<script lang="ts">
	import FormIcon from '$components/Icons/FormIcon/FormIcon.svelte';
</script>

<header>
	<FormIcon />
	<h1>校大学生科协报名表单</h1>
</header>

<style lang="postcss">
	header {
		@apply w-full my-3 flex flex-row items-center justify-center gap-2;
	}
	h1 {
		@apply text-2xl font-semibold text-gray-800;
	}
</style>
