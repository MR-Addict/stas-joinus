<script lang="ts">
	import { links } from './config';
	import FaRegUserCircle from 'svelte-icons/fa/FaRegUserCircle.svelte';
</script>

<nav class="px-4 md:px-20 py-4 border-b border-b-gray-300">
	<a href="/" class="flex flex-row items-center gap-1">
		<img src="/favicon.png" alt="logo" class="w-7" />
		<h1 class="font-semibold text-gray-800">校大学生科协</h1>
	</a>

	<ul class="flex flex-row items-center gap-2 md:gap-4 ml-auto">
		{#each links as link (link.name)}
			<li>
				<a href={link.path}>
					<div class="w-4">
						<svelte:component this={link.Icon} />
					</div>
					<div>{link.name}</div>
				</a>
			</li>
		{/each}
	</ul>

	<div>
		<div class="w-5"><FaRegUserCircle /></div>
	</div>
</nav>

<style lang="postcss">
	nav {
		@apply flex flex-row items-center gap-2 md:gap-4;
	}
	li a {
		@apply flex flex-row items-center gap-0.5 md:gap-1 duration-300;
		&:hover {
			@apply text-green-600;
		}
	}
</style>
