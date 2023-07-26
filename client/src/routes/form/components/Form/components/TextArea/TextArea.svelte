<script lang="ts">
	export let title: string;
	export let description: string;
	export let name: string;
	export let validate: (value: string) => string | undefined;

	let error: string | undefined = undefined;
	const handleChange = (event: Event) => (error = validate((event.target as HTMLTextAreaElement).value));
</script>

<div class="wrapper">
	<h1>
		<span>{title}</span>
		<span class="text-red-600">*</span>
	</h1>
	<label for={name}>{description}</label>
	<textarea required id={name} {name} placeholder={title} on:change={handleChange} />
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
	textarea {
		@apply h-32 bg-white text-sm outline-none p-2 border border-gray-400 rounded-sm;
		&:focus {
			@apply border-teal-600;
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
