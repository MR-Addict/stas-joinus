<script lang="ts">
	import Header from './components/Header/Header.svelte';
	import Actions from './components/Actions/Actions.svelte';
</script>

<nav class="gradienter">
	<Header />
	<Actions />
</nav>

<style>
	nav {
		@apply z-20 flex flex-row items-center justify-between gap-4;
		@apply py-4 px-4 sm:px-10 border-b border-b-gray-300 sticky top-0 bg-white;
	}
</style>
