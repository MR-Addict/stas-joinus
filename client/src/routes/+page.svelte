<script lang="ts">
	import Fireworks from '$components/Icons/Fireworks/Fireworks.svelte';
</script>

<svelte:head>
	<title>欢迎加入校科协</title>
</svelte:head>

<main>
	<div class="w-36"><Fireworks /></div>

	<h1 class="text-3xl font-semibold mt-2 mb-3">欢迎加入校科协</h1>

	<p class="text-gray-700 max-w-md">
		欢迎2023级的新生来到我们南京工业大学，如果你对校大学生科协技术协会感兴趣，想和我们一起度过一个充实且精彩的大学生活，可以填写以下表单加入我们，或者加群和我们互动交流
	</p>

	<p class="text-gray-500 mt-1 mb-3">报名时间：2023-09-04至2023-09-15</p>

	<div class="flex flex-row gap-5">
		<a href="/contact">加群了解</a>
		<a href="/form">参与报名</a>
	</div>
</main>

<style lang="postcss">
	main {
		min-height: calc(100vh - 60px);
		min-height: calc(100svh - 60px);
		@apply flex flex-col items-center justify-center text-center px-4;
	}
	a {
		@apply font-semibold py-2.5 px-7 rounded-full duration-300;
	}
	a:first-of-type {
		@apply border-2 border-teal-800 text-gray-800;
		&:hover {
			@apply text-teal-800;
		}
	}
	a:last-of-type {
		@apply text-white bg-teal-800;
		&:hover {
			@apply bg-teal-900;
		}
	}
</style>
