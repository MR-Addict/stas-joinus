<script lang="ts">
	import User from '$components/Icons/User/User.svelte';
</script>

<nav class="px-4 md:px-20 py-4 border-b border-b-gray-300">
	<a href="/" class="flex flex-row items-center gap-1">
		<img src="/favicon.png" alt="logo" class="w-7" />
		<h1 class="font-semibold text-gray-800">校大学生科协</h1>
	</a>

	<div class="flex items-center justify-center">
		<button type="button" class="w-6 text-gray-700 duration-300 hover:text-black"><User /></button>
	</div>
</nav>

<style lang="postcss">
	nav {
		@apply flex flex-row items-center justify-between;
	}
</style>
