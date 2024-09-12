<script lang="ts">
	import toast from 'svelte-french-toast';
	import { Check, Download, Files } from 'lucide-svelte';

	import qrcode from './images/qrcode.png';
	import Modal from '$components/Modal/Modal.svelte';

	export let showModal = false;

	let copied = false;
	const groupId = '1002668803';

	function handleCopy() {
		if (!navigator.clipboard) return toast.error('当前浏览器环境不支持复制操作');

		copied = true;
		toast.success('群号复制成功');
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
			<p class="text-gray-600">
				<span>群号: </span>
				<span>{groupId}</span>
			</p>

			<button type="button" on:click={handleCopy} class="btn" aria-label="copy">
				<div class:active={copied}>
					<Check size={16} />
				</div>

				<div class:active={!copied}>
					<Files size={14} />
				</div>
			</button>

			<button type="button" aria-label="download" class="btn" on:click={handleDownload}>
				<Download size={16} />
			</button>
		</div>
	</div>
</Modal>

<style>
	.wrapper {
		@apply bg-white rounded-xl flex flex-col items-center justify-center p-10 gap-1;
	}
	.btn {
		@apply w-7 aspect-square relative text-gray-600 rounded-full bg-black/10 grid place-content-center;

		& div {
			@apply absolute top-1/2 left-1/2 -translate-y-1/2 -translate-x-1/2 scale-0 invisible opacity-0;

			&.active {
				@apply transform scale-100 visible opacity-100 duration-300;
			}
		}
	}
</style>
