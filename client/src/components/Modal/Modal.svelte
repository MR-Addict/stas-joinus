<script lang="ts">
	import { browser } from '$app/environment';

	export let showModal: boolean;

	$: if (browser) document.body.style.overflowY = showModal ? 'hidden' : 'auto';
</script>

{#if showModal}
	<div class="backdrop">
		<!-- svelte-ignore a11y-no-static-element-interactions -->
		<!-- svelte-ignore a11y-click-events-have-key-events -->
		<div class="modal" on:click|self={() => (showModal = false)}>
			<slot />
		</div>
	</div>
{/if}

<style lang="postcss">
	.backdrop {
		@apply bg-black/30 fixed inset-0;
	}
	.modal {
		@apply w-full h-full flex flex-col items-center justify-center;
		animation: zoom 500ms cubic-bezier(0.34, 1.56, 0.64, 1);
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
