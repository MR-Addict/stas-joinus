<script lang="ts">
	import { Forward } from 'lucide-svelte';

	import { links } from './config';
</script>

<footer>
	<ul class="w-full grid grid-cols-2 md:grid-cols-3 gap-5">
		{#each links as category (category.group)}
			<li class="space-y-2">
				<h1 class="font-semibold text-gray-200">{category.group}</h1>
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

	<p class="copyright">
		<span>Copyright &copy; {new Date().getFullYear()} 校大学生科协.</span>
		<a href="https://beian.miit.gov.cn" target="_blank">苏ICP备2022032826号</a>
	</p>
</footer>

<style lang="postcss">
	footer {
		@apply bg-slate-800 text-white px-4 sm:px-32 py-10 flex flex-col gap-4;
	}

	a {
		@apply flex flex-row items-center gap-1 w-fit duration-300;
		&:hover {
			@apply text-green-600;
		}
		& .icon {
			@apply block w-4;
		}
	}
	.copyright {
		@apply flex flex-row flex-wrap gap-2;
	}
</style>
