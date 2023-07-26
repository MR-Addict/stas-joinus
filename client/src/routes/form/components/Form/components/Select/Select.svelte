<script lang="ts">
	export let title: string;
	export let description: string;
	export let name: string;
	export let options: string[];
	export let validate: ((value: string) => string | undefined) | undefined = undefined;

	let error: string | undefined = undefined;
	const handleChange = (event: Event) =>
		(error = validate ? validate((event.target as HTMLSelectElement).value) : undefined);
</script>

<div class="wrapper">
	<h1>
		<span>{title}</span>
		<span class="text-red-600">*</span>
	</h1>
	<label for={name}>{description}</label>
	<select id={name} {name} required on:change={handleChange}>
		{#each options as option (option)}
			<option value={option}>{option}</option>
		{/each}
	</select>
	<p class:active={error !== undefined}>{error}</p>
</div>

<style lang="postcss">
	.wrapper {
		@apply w-full bg-white p-5 rounded-md flex flex-col gap-2 border border-gray-200;
	}
	h1 {
		@apply font-semibold;
	}
	label {
		@apply text-sm text-gray-500;
	}
	select {
		@apply rounded-none bg-white text-sm outline-none pt-1 border-b border-b-gray-400;
		&:focus {
			@apply border-b-teal-600;
		}
	}
	p {
		@apply text-xs max-h-0 overflow-hidden text-white;
		transition: max-height 0.25s ease-out;
		&.active {
			@apply text-red-600 max-h-10;
			transition: max-height 0.25s ease-in;
		}
	}
</style>
