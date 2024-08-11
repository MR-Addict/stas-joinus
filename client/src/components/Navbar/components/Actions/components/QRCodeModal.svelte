<script lang="ts">
	import qrcode from './images/qrcode.png';
	import Modal from '$components/Modal/Modal.svelte';
	import { Check, Download, Files } from 'lucide-svelte';

	export let showModal = false;

	let copied = false;
	const groupId = '1002668803';

	function handleCopy() {
		if (!navigator.clipboard) return;

		copied = true;
		navigator.clipboard.writeText(groupId);
		setTimeout(() => (copied = false), 1000);
	}

	function handleDownload() {
		const a = document.createElement('a');
		a.href = qrcode;
		a.download = 'qrcode.png';
		a.click();
	}
</script>

<Modal bind:showModal>
	<div class="wrapper">
		<img src={qrcode} alt="qrcode" />

		<div class="flex flex-row items-center gap-2 mt-4">
			<p>{`群号: ${groupId}`}</p>

			<button type="button" on:click={handleCopy} class="copy-btn" aria-label="copy">
				<div class:active={copied}>
					<Check size={16} />
				</div>

				<div class:active={!copied}>
					<Files size={16} />
				</div>
			</button>
		</div>

		<button type="button" aria-label="download" class="download-btn" on:click={handleDownload}>
			<Download size={32} />
		</button>
	</div>
</Modal>

<style>
	.wrapper {
		@apply bg-white rounded-xl;
		@apply w-full sm:w-fit sm:h-full flex flex-col items-center justify-center p-10 gap-1;
	}
	img {
		@apply max-h-full object-contain rounded-2xl;
	}
	.download-btn {
		@apply text-gray-800;
	}
	.copy-btn {
		@apply w-7 aspect-square relative text-gray-600 rounded-full bg-black/10;

		& div {
			@apply absolute top-1/2 left-1/2 -translate-y-1/2 -translate-x-1/2 scale-0 invisible opacity-0;

			&.active {
				@apply scale-100 visible opacity-100 duration-300;
			}
		}
	}
</style>
