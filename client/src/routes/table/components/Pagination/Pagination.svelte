<script lang="ts">
	import toasts from '$stores/toasts';
	import pagination from '$stores/pagination';
	import fetchApplicants from '$lib/applicant/fetchApplicants';

	let page = 1;
	let total = 0;
	let pageSize = 20;
	let totalPages = 1;

	let pending = false;

	$: {
		page = $pagination.page;
		total = $pagination.total;
		pageSize = $pagination.page_size;
		totalPages = Math.ceil(total / pageSize);
	}

	async function handleClick(page: number) {
		pending = true;
		if (await fetchApplicants(page)) window.scrollTo({ top: 0 });
		else toasts.add('无法加载下一页', 'failed');
		pending = false;
	}
</script>

{#if totalPages > 1}
	<div class="wrapper">
		<p>{page}/{totalPages}页(共{total}记录)</p>
		<div class="ml-auto flex flex-row flex-wrap gap-2">
			{#each { length: totalPages } as item, index (index)}
				{@const display = index + 1}
				<button
					type="button"
					class:active={page === display}
					on:click={() => handleClick(display)}
					disabled={page === display || pending}
				>
					{display}
				</button>
			{/each}
		</div>
	</div>
{/if}

<style lang="postcss">
	.wrapper {
		@apply w-full flex flex-row items-center flex-wrap gap-2 text-sm text-gray-600;
	}
	button {
		@apply w-6 h-6 flex items-center justify-center border border-gray-400;
		&.active {
			@apply text-blue-600;
		}
	}
</style>
