<script lang="ts">
	import flower from './flower.svg';
</script>

<div>
	<img src={flower} alt="flower" />
	<h1>提交成功</h1>
	<p>恭喜你提交成功啦，点击右上角加入群聊，以获取更多后续的通知内容</p>
	<button type="button">修改志愿</button>
</div>

<style lang="postcss">
	div {
		@apply w-full flex flex-col items-center justify-center gap-2;
	}
	h1 {
		@apply text-2xl font-semibold text-gray-800;
	}
	p {
		text-wrap: balance;
		@apply max-w-sm text-lg text-center text-gray-700;
	}
	button {
		@apply mt-2 text-lg bg-yellow-400 border-2 border-black;
		@apply font-semibold py-2.5 px-7 rounded-md duration-300;
	}
</style>
