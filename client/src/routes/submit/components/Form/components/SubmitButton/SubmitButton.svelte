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
		@apply text-lg bg-yellow-400 border-2 border-black;
		@apply font-semibold py-2.5 px-7 rounded-md duration-300;
		&:enabled:hover {
			@apply bg-yellow-500;
		}
		&:disabled {
			@apply bg-gray-300 text-gray-400 cursor-not-allowed;
		}
	}
</style>
