<script lang="ts">
	export let title: string;
	export let description: string;
	export let name: string;
	export let minlength: number | undefined = undefined;
	export let maxlength: number | undefined = undefined;
</script>

<div class="wrapper">
	<h1>
		<span>{title}</span>
		<span class="text-red-600">*</span>
	</h1>
	<label for={name}>{description}</label>
	<textarea id={name} {minlength} {maxlength} required {name} placeholder={title} />
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
</style>
