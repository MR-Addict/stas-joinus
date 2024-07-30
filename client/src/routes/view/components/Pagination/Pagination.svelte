<script lang="ts">
	import toast from 'svelte-french-toast';

	import view from '$stores/view';
	import type { PaginationType } from '$types/pagination';

	export let pagination: PaginationType;

	let totalPages = 1;

	let pending = false;

	$: {
		totalPages = Math.ceil(pagination.total / pagination.page_size);
	}

	async function handleClick(page: number) {
		pending = true;

		const res = await view.refersh(page);
		if (res.success) window.scrollTo({ top: 0, behavior: 'smooth' });
		else toast.error(res.message);

		pending = false;
	}
</script>

{#if totalPages > 1}
	<div class="wrapper">
		<p>{pagination.page}/{totalPages}页(共{pagination.total}记录)</p>

		<div class="ml-auto flex flex-row flex-wrap gap-2">
			{#each { length: totalPages } as _, index (index)}
				{@const display = index + 1}
				<button
					type="button"
					on:click={() => handleClick(display)}
					class:active={pagination.page === display}
					disabled={pagination.page === display || pending}
				>
					{display}
				</button>
			{/each}
		</div>
	</div>
{/if}

<style>
	.wrapper {
		@apply w-full flex flex-row items-center flex-wrap gap-2;
	}
	p {
		@apply text-gray-600;
	}
	button {
		@apply w-6 h-6 flex items-center justify-center border border-gray-400 rounded-sm;

		&.active {
			@apply bg-teal-700 text-white border-transparent;
		}
	}
</style>
