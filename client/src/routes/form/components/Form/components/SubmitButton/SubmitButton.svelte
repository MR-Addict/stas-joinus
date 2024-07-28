<script lang="ts">
	import Spinner from '$components/Spinner/Spinner.svelte';

	export let pending: boolean;
</script>

<button disabled={pending} type="submit">
	{#if pending}
		<Spinner />
	{/if}
	<span>提交</span>
</button>

<style lang="postcss">
	button {
		@apply flex flex-row items-center gap-1 bg-blue-600 text-white py-2 px-6 rounded-md duration-300;
		&:hover {
			@apply bg-blue-700;
		}
		&:disabled {
			@apply bg-gray-300 text-gray-400 cursor-not-allowed;
		}
	}
</style>
