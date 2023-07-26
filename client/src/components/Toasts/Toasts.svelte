<script lang="ts">
	import { fly } from 'svelte/transition';
	import { toasts } from '$stores/toasts';

	import GoInfo from 'svelte-icons/go/GoInfo.svelte';
	import MdClose from 'svelte-icons/md/MdClose.svelte';
	import MdCancel from 'svelte-icons/md/MdCancel.svelte';
	import FaCheckCircle from 'svelte-icons/fa/FaCheckCircle.svelte';

	let finishedToastIds: string[] = [];

	$: {
		$toasts.finished
			.filter((item) => !finishedToastIds.includes(item.id))
			.forEach((toast) => {
				setTimeout(() => {
					finishedToastIds.push(toast.id);
					toasts.remove(toast.id);
				}, 2000);
			});
	}
</script>

<ul aria-label="toasts">
	{#each $toasts.all as toast (toast.id)}
		<li class="toast" transition:fly={{ x: '10rem' }} data-status={toast.status}>
			<div class="icon">
				{#if toast.status === 'pending'}
					<GoInfo />
				{:else if toast.status === 'failed'}
					<MdCancel />
				{:else}
					<FaCheckCircle />
				{/if}
			</div>

			<p title={toast.message} class="message">{toast.message}</p>

			<button type="button" title="隐藏" on:click={() => toasts.remove(toast.id)}>
				<div class="w-5 h-5 text-gray-500 p-0.5 rounded-full hover:bg-black/10 hover:text-gray-700"><MdClose /></div>
			</button>
		</li>
	{/each}
</ul>

<style lang="postcss">
	ul {
		max-height: 60vh;
		max-height: 60svh;
		@apply top-10 z-10 fixed right-2.5 flex flex-col gap-4 overflow-y-auto overflow-x-hidden p-1.5;
	}
	ul::-webkit-scrollbar {
		display: none;
	}
	.toast {
		box-shadow: 0 0 5px #b9bcc1;
		@apply bg-white border-b-4 border-b-blue-600;
		@apply w-full py-4 px-4 rounded-md flex flex-row items-center justify-between gap-2;
	}
	.toast[data-status='success'] {
		@apply border-b-green-600;
	}
	.toast[data-status='failed'] {
		@apply border-b-red-600;
	}
	.icon {
		@apply w-5 h-5 text-blue-600;
	}
	.toast[data-status='success'] .icon {
		@apply text-green-600;
	}
	.toast[data-status='failed'] .icon {
		@apply text-red-600;
	}
	.message {
		max-width: 15rem;
		@apply w-screen whitespace-nowrap overflow-hidden text-ellipsis text-sm;
	}
</style>
