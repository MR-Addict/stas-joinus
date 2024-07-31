<script lang="ts">
	import formatDate from '$lib/utils/formatDate';
	import type { AppConfigType } from '$types/app';

	export let config: AppConfigType;
</script>

<svelte:head>
	<title>校大学生科协 • 参与报名</title>
</svelte:head>

{#if new Date() <= config.start_time}
	<div class="space-y-1 text-center">
		<p class="text-gray-700">开始时间：{formatDate(config.start_time)}</p>
		<p>报名未开始，你可以点击右上角加群了解</p>
	</div>
{:else if new Date() >= config.end_time}
	<div class="space-y-1 text-center">
		<p class="text-gray-700">结束时间：{formatDate(config.end_time)}</p>
		<p>报名结束啦，有问题可以点击右上角加群了解</p>
	</div>
{:else}
	<slot />
{/if}
