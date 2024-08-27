<script lang="ts">
	import config from '$stores/config';
	import formatDate from '$lib/utils/formatDate';
	import billboard from './images/billboard.png';
</script>

<svelte:head>
	<title>校大学生科协</title>
</svelte:head>

<main>
	<img src={billboard} alt="billboard" />

	<h1 class="text-xl font-semibold mt-5 mb-2">欢迎加入校大学生科协</h1>

	<p class="text-gray-700 max-w-sm text-center" style="text-wrap: balance;">
		填写下面的表单加入校大学生科协技术协会，和我们一起度过一个精彩的大学生活吧
	</p>

	<div class="mt-1 flex flex-col">
		<p class="config-time">
			<span>开始时间：</span>
			{#if $config === null}
				<span>无法连接至服务器</span>
			{:else if $config === undefined}
				<span class="skeleton"></span>
			{:else if $config}
				<span>{formatDate($config.start_time)}</span>
			{/if}
		</p>

		<p class="config-time">
			<span>结束时间：</span>
			{#if $config === null}
				<span>无法连接至服务器</span>
			{:else if $config === undefined}
				<span class="skeleton"></span>
			{:else if $config}
				<span>{formatDate($config.end_time)}</span>
			{/if}
		</p>
	</div>

	<a href="/submit">现在报名</a>
</main>

<style>
	main {
		@apply flex flex-col items-center justify-center px-4;
	}
	img {
		@apply max-w-xs w-full;
	}
	.config-time {
		@apply flex flex-row items-center text-gray-500 sm:text-sm;

		& .skeleton {
			width: 144.88px;
			@apply h-5 rounded-sm bg-gray-300 animate-pulse;
		}
	}
	a {
		@apply mt-4 text-lg bg-yellow-400 border-2 border-black;
		@apply font-semibold py-2.5 px-7 rounded-md duration-300;

		&:hover {
			@apply bg-yellow-500;
		}
	}
</style>
