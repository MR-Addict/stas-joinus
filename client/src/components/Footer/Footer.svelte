<script lang="ts">
	import { Forward } from 'lucide-svelte';

	import { links } from './config';
</script>

<footer>
	<ul class="w-full grid grid-cols-2 md:grid-cols-3 gap-5">
		{#each links as category (category.group)}
			<li class="space-y-2">
				<h1 class="font-semibold">{category.group}</h1>
				<ul>
					{#each category.items as item (item.name)}
						<li>
							<a href={item.path} target="_blank">
								<div>{item.name}</div>
								<div class="icon"><Forward size={14} /></div>
							</a>
						</li>
					{/each}
				</ul>
			</li>
		{/each}
	</ul>

	<p>
		{`Copyright © ${new Date().getFullYear()} 校大学生科协`}
	</p>
</footer>

<style>
	footer {
		@apply px-4 sm:px-32 py-10 flex flex-col gap-4 border-t border-t-gray-300;
	}

	a {
		@apply flex flex-row items-center gap-1 w-fit duration-300 text-gray-700;

		&:hover {
			@apply text-green-600;
		}

		& .icon {
			@apply block w-4;
		}
	}
</style>
