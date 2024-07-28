<script lang="ts">
	import billboard from './images/billboard.png';
</script>

<svelte:head>
	<title>校大学生科协</title>
</svelte:head>

<main>
	<img src={billboard} alt="billboard" />

	<h1 class="text-2xl font-semibold my-5">欢迎加入校大学生科协</h1>

	<p class="text-lg text-gray-700 max-w-sm text-center" style="text-wrap: balance;">
		欢迎2024级新生来到南京工业大学，填写下面的表单加入校大学生科协技术协会，和我们一起度过一个精彩的大学生活吧
	</p>

	<p class="text-gray-500 mt-1 mb-3">报名时间：2024-09-04至2024-09-15</p>

	<a href="/form">立即报名</a>
</main>

<style lang="postcss">
	main {
		min-height: calc(100vh - 4rem);
		min-height: calc(100dvh - 4rem);
		@apply flex flex-col items-center justify-center px-4;
	}
	img {
		@apply max-w-sm w-full;
	}
	a {
		@apply text-lg bg-yellow-400 border-2 border-black;
		@apply font-semibold py-2.5 px-7 rounded-md duration-300;

		&:hover {
			@apply bg-yellow-500;
		}
	}
</style>
