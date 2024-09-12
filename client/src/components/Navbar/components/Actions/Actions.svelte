<script lang="ts">
	import Dropdown from './components/Dropdown.svelte';
	import ExportModal from './components/ExportModal.svelte';
	import LoginModal from './components/LoginModal.svelte';
	import LogoutModal from './components/LogoutModal.svelte';
	import QrCodeModal from './components/QRCodeModal.svelte';

	let showModal = { login: false, logout: false, qrcode: false, export: false };
</script>

<LoginModal bind:showModal={showModal.login} />
<LogoutModal bind:showModal={showModal.logout} />
<QrCodeModal bind:showModal={showModal.qrcode} />
<ExportModal bind:showModal={showModal.export} />

<div class="wrapper">
	<button type="button" class="btn" on:click={() => (showModal.qrcode = true)}>加群了解</button>
	<Dropdown bind:showModal />
</div>

<style>
	.wrapper {
		@apply relative flex items-center justify-center gap-3 sm:gap-5;
	}

	.btn {
		@apply py-1.5 px-3 sm:py-2 sm:px-4 rounded-md duration-300 text-white bg-teal-700 text-sm;

		&:hover {
			@apply bg-teal-800;
		}
	}
</style>
