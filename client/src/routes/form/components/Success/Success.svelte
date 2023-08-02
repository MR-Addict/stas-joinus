<script>
	import Flower from '$components/Icons/Flower/Flower.svelte';
</script>

<div>
	<Flower />
	<h1>提交成功</h1>
	<p>恭喜你提交成功啦，我们会尽快通知你面试时间，请耐心等待，你也可以加群和我们交流</p>
	<a href="/">返回</a>
</div>

<style lang="postcss">
	div {
		@apply w-full flex flex-col items-center justify-center gap-2;
	}
	h1 {
		@apply text-2xl font-semibold text-gray-800;
	}
	p {
		@apply max-w-xs text-center text-gray-500;
	}
	a {
		@apply text-gray-500 text-sm py-1 px-2 rounded-3xl border border-gray-400 duration-300;
		&:hover {
			@apply text-gray-700 border-gray-500;
		}
	}
</style>
