<script lang="ts">
	export let toggleChoice: boolean;
</script>

<button type="button" on:click={() => (toggleChoice = !toggleChoice)}>
	<span class="bubble" class:active={toggleChoice} />
	<span class="option" class:active={toggleChoice}>第一志愿</span>
	<span class="option" class:active={!toggleChoice}>第二志愿</span>
</button>

<style>
	button {
		@apply py-2 px-2 bg-black/10 isolate backdrop-blur-lg;
		@apply sticky top-20 z-10 mx-auto flex flex-row items-center gap-2 rounded-md;

		& .option {
			@apply py-1 px-2 rounded-md duration-300;
		}

		& .bubble {
			--gap: 0.25rem;
			--py: 0.45rem;
			--px: 0.5rem;
			top: var(--py);
			right: var(--px);
			height: calc(100% - var(--py) * 2);
			width: calc(50% - var(--px) - var(--gap));
			@apply absolute bg-black/20 -z-10 rounded-md duration-500 ease-in-out;

			&.active {
				transform: translateX(calc(-100% - var(--px)));
			}
		}
	}
</style>
