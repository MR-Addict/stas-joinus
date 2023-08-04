<script lang="ts">
	import { toasts } from '$stores/toasts';
	import clickOutside from '$hooks/clickOutside';
	import fetchApplicants from '$lib/applicant/fetchApplicants';
	import type { TableFilter } from '$types/tableFilter';

	import TiFilter from 'svelte-icons/ti/TiFilter.svelte';
	import MdRefresh from 'svelte-icons/md/MdRefresh.svelte';
	import MdFileDownload from 'svelte-icons/md/MdFileDownload.svelte';

	export let tableFilter: TableFilter;

	let pending = false;
	let showFilter = false;

	async function handleRefresh() {
		pending = true;
		const id = toasts.add('数据刷新中', 'pending');
		if (await fetchApplicants()) toasts.update(id, { message: '刷新成功', status: 'success' });
		pending = false;
	}
</script>

<header>
	<h1>
		<span>所有提交</span>
		<button disabled={pending} type="button" on:click={handleRefresh}><MdRefresh /></button>
	</h1>

	<button type="button" class="ml-auto btn">
		<div><MdFileDownload /></div>
		<span>下载表格</span>
	</button>

	<div class="relative" use:clickOutside={() => (showFilter = false)}>
		<button type="button" class="btn" on:click={() => (showFilter = !showFilter)}>
			<div><TiFilter /></div>
			<span>筛选表格</span>
		</button>

		<div class="filter" class:active={showFilter}>
			<div class="option">
				<input type="checkbox" id="name" bind:checked={tableFilter.name} />
				<label for="name">姓名</label>
			</div>
			<div class="option">
				<input type="checkbox" id="gender" bind:checked={tableFilter.gender} />
				<label for="gender">性别</label>
			</div>
			<div class="option">
				<input type="checkbox" id="phone" bind:checked={tableFilter.phone} />
				<label for="phone">手机</label>
			</div>
			<div class="option">
				<input type="checkbox" id="qq" bind:checked={tableFilter.qq} />
				<label for="qq">QQ</label>
			</div>
			<div class="option">
				<input type="checkbox" id="email" bind:checked={tableFilter.email} />
				<label for="email">邮箱</label>
			</div>
			<div class="option">
				<input type="checkbox" id="student_id" bind:checked={tableFilter.student_id} />
				<label for="student_id">学号</label>
			</div>
			<div class="option">
				<input type="checkbox" id="college" bind:checked={tableFilter.college} />
				<label for="college">学院</label>
			</div>
			<div class="option">
				<input type="checkbox" id="major" bind:checked={tableFilter.major} />
				<label for="major">学院</label>
			</div>
			<div class="option">
				<input type="checkbox" id="created_at" bind:checked={tableFilter.created_at} />
				<label for="created_at">提交时间</label>
			</div>
			<div class="option">
				<input type="checkbox" id="first_choice" bind:checked={tableFilter.first_choice} />
				<label for="first_choice">第一志愿</label>
			</div>
			<div class="option">
				<input type="checkbox" id="second_choice" bind:checked={tableFilter.second_choice} />
				<label for="second_choice">第二志愿</label>
			</div>
			<div class="option">
				<input type="checkbox" id="introduction" bind:checked={tableFilter.introduction} />
				<label for="introduction">自我介绍</label>
			</div>
		</div>
	</div>
</header>

<style lang="postcss">
	header {
		@apply flex flex-row items-center gap-4;
	}
	h1 {
		@apply flex flex-row items-center gap-0.5;
		& span {
			@apply text-gray-800;
		}
		& button {
			@apply text-gray-600 w-4;
			&:hover {
				@apply text-gray-800;
			}
		}
	}

	.filter {
		box-shadow: 0 0 5px #d1d5db;
		@apply hidden flex-col border border-gray-300;
		@apply rounded-md py-2.5 px-4 bg-white mt-1 absolute right-0 md:right-1/2 md:translate-x-1/2 top-full;
		&.active {
			@apply flex;
		}
	}

	.option {
		@apply whitespace-nowrap flex flex-row items-center gap-1.5;
		& input {
			appearance: none;
			@apply relative;
			@apply w-4 h-4 border border-gray-500 rounded-sm cursor-pointer;
			&::before {
				content: '';
				@apply block w-2.5 h-2.5 bg-blue-600 scale-0 duration-100;
				@apply absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2;
				clip-path: polygon(14% 44%, 0 65%, 50% 100%, 100% 16%, 80% 0%, 43% 62%);
			}
			&:checked::before {
				@apply scale-100;
			}
		}
	}

	.btn {
		@apply text-sm flex flex-row items-center text-gray-600;
		&:hover {
			@apply text-black border-gray-500;
		}
		& div {
			width: 17px;
		}
	}
</style>
