<script lang="ts">
	export let showModal: boolean;

	let dialog: HTMLDialogElement;

	function openModal() {
		dialog.showModal();
		showModal = true;
		document.body.style.overflowY = 'hidden';
	}

	function closeModal() {
		dialog.close();
		showModal = false;
		document.body.style.overflowY = 'auto';
	}

	$: if (dialog) {
		if (showModal) openModal();
		else closeModal();
	}
</script>

<!-- svelte-ignore a11y-click-events-have-key-events a11y-no-noninteractive-element-interactions -->
<dialog bind:this={dialog} on:click|self={closeModal}>
	<slot />
</dialog>

<style>
	dialog {
		background-color: transparent;
		animation: zoom 300ms cubic-bezier(0.34, 1.56, 0.64, 1);
	}
	dialog::backdrop {
		background: rgba(0, 0, 0, 0.3);
	}
	@keyframes zoom {
		from {
			transform: scale(0.9);
		}
		to {
			transform: scale(1);
		}
	}
</style>
