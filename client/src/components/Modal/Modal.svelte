<script lang="ts">
	import { scale } from 'svelte/transition';
	import { browser } from '$app/environment';

	export let showModal: boolean;

	$: if (browser) document.body.style.overflowY = showModal ? 'hidden' : 'auto';
</script>

{#if showModal}
	<div class="backdrop">
		<!-- svelte-ignore a11y-no-static-element-interactions -->
		<!-- svelte-ignore a11y-click-events-have-key-events -->
		<div
			class="modal"
			on:click|self={() => (showModal = false)}
			transition:scale={{ duration: 400, start: 0.4, opacity: 0 }}
		>
			<slot />
		</div>
	</div>
{/if}

<style lang="postcss">
	.backdrop {
		@apply backdrop-blur-sm bg-black/30 fixed inset-0 z-10;
	}
	.modal {
		@apply w-full h-full flex flex-col items-center justify-center;
	}
</style>
