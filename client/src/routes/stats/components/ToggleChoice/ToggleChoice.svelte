<script lang="ts">
	import stats from '$stores/stats';
	import type { ChoiceType } from '$types/applicant';
	import { firstChoice, secondChoice } from '$data/choice';

	export let choice: ChoiceType;

	function handleToggle() {
		stats.refersh();
		choice = choice.label === 'first_choice' ? secondChoice : firstChoice;
	}
</script>

<button type="button" on:click={handleToggle}>
	<span class="bubble" class:active={choice.label === 'first_choice'} />
	<span class="option" class:active={choice.label === 'first_choice'}>第一志愿</span>
	<span class="option" class:active={choice.label === 'second_choice'}>第二志愿</span>
</button>

<style>
	button {
		--py: 0.45rem;
		--px: 0.5rem;
		padding: var(--py) var(--px);
		gap: var(--gap);
		@apply bg-black/10 isolate backdrop-blur-lg;
		@apply sticky top-20 z-10 flex flex-row items-center rounded-md;

		& .option {
			@apply py-1 px-2 rounded-md duration-300;
		}

		& .bubble {
			--gap: 0.25rem;
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
