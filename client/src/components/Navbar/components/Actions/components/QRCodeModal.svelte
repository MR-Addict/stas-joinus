<script lang="ts">
	import qrcode from './images/qrcode.jpg';
	import Modal from '$components/Modal/Modal.svelte';

	export let showModal = false;
</script>

<Modal bind:showModal>
	<div>
		<img src={qrcode} alt="qrcode" />
	</div>
</Modal>

<style lang="postcss">
	div {
		@apply w-full h-full pointer-events-none flex items-center justify-center p-10;
	}
	img {
		@apply max-h-full object-contain rounded-md;
	}
</style>
