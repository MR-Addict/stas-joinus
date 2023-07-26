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
	{#if error}
		<p>{error}</p>
	{/if}
</div>

<style lang="postcss">
	.wrapper {
		@apply w-full bg-white p-5 rounded-md flex flex-col gap-2 border border-gray-300;
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
		@apply text-xs text-red-600;
	}
</style>
