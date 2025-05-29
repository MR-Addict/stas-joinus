<script lang="ts">
	import { Filter, RefreshCw } from 'lucide-svelte';

	import view from '$stores/view';
	import toast from 'svelte-french-toast';
	import clickOutside from '$hooks/clickOutside';
	import type { Writable } from 'svelte/store';
	import type { TableFilterType } from '$types/tableFilter';

	export let tableFilter: Writable<TableFilterType>;

	let pending = false;
	let refreshed = false;
	let showFilter = false;

	async function handleRefresh() {
		pending = true;
		refreshed = true;
		setTimeout(() => (refreshed = false), 500);

		const toastId = toast.loading('数据刷新中，请稍后...');
		const res = await view.refersh();
		if (res.success) toast.success(res.message, { id: toastId });
		else toast.error(res.message, { id: toastId });

		pending = false;
	}

	function updateFilter(e: Event) {
		const target = e.target as HTMLInputElement;
		tableFilter.update((t) => ({ ...t, [target.id]: target.checked }));
	}
</script>

<header>
	<h1>
		<span>所有数据</span>
		<button disabled={pending} type="button" class:refreshed on:click={handleRefresh}><RefreshCw size={16} /></button>
	</h1>

	<div class="ml-auto relative" use:clickOutside={() => (showFilter = false)}>
		<button type="button" class="action-btn" on:click={() => (showFilter = !showFilter)}>
			<div><Filter size={16} /></div>
		</button>

		<div class="filter" class:active={showFilter}>
			<div class="option">
				<input type="checkbox" id="name" checked={$tableFilter.name} on:change={updateFilter} />
				<label for="name">姓名</label>
			</div>
			<div class="option">
				<input type="checkbox" id="gender" checked={$tableFilter.gender} on:change={updateFilter} />
				<label for="gender">性别</label>
			</div>
			<div class="option">
				<input type="checkbox" id="phone" checked={$tableFilter.phone} on:change={updateFilter} />
				<label for="phone">手机</label>
			</div>
			<div class="option">
				<input type="checkbox" id="email" checked={$tableFilter.email} on:change={updateFilter} />
				<label for="email">邮箱</label>
			</div>
			<div class="option">
				<input type="checkbox" id="qq" checked={$tableFilter.qq} on:change={updateFilter} />
				<label for="qq">QQ</label>
			</div>
			<div class="option">
				<input type="checkbox" id="student_id" checked={$tableFilter.student_id} on:change={updateFilter} />
				<label for="student_id">学号</label>
			</div>
			<div class="option">
				<input type="checkbox" id="college" checked={$tableFilter.college} on:change={updateFilter} />
				<label for="college">学院</label>
			</div>
			<div class="option">
				<input type="checkbox" id="major" checked={$tableFilter.major} on:change={updateFilter} />
				<label for="major">专业</label>
			</div>
			<div class="option">
				<input type="checkbox" id="created_at" checked={$tableFilter.created_at} on:change={updateFilter} />
				<label for="created_at">提交时间</label>
			</div>
			<div class="option">
				<input type="checkbox" id="first_choice" checked={$tableFilter.first_choice} on:change={updateFilter} />
				<label for="first_choice">第一志愿</label>
			</div>
			<div class="option">
				<input type="checkbox" id="second_choice" checked={$tableFilter.second_choice} on:change={updateFilter} />
				<label for="second_choice">第二志愿</label>
			</div>
			<div class="option">
				<input type="checkbox" id="introduction" checked={$tableFilter.introduction} on:change={updateFilter} />
				<label for="introduction">自我介绍</label>
			</div>
		</div>
	</div>
</header>

<style>
	header {
		@apply flex flex-row items-center gap-4;
	}
	h1 {
		@apply flex flex-row items-center gap-1;

		& span {
			@apply text-gray-800 text-lg font-semibold;
		}

		& button {
			@apply text-gray-600 w-4;

			&.refreshed {
				@apply rotate-180 duration-500;
			}
		}
	}

	.filter {
		box-shadow: 0 0 5px #d1d5db;
		@apply z-10 flex flex-col border border-gray-300;
		@apply origin-top invisible opacity-0 duration-200 scale-y-95;
		@apply rounded-md py-3 px-4 bg-white mt-1 absolute right-0 top-full;

		&.active {
			@apply visible opacity-100 scale-y-100;
		}
	}

	.option {
		@apply whitespace-nowrap flex flex-row items-center gap-2;

		& input {
			appearance: none;
			@apply relative w-4 h-4 bg-gray-300 rounded-sm cursor-pointer;

			&::before {
				content: '';
				@apply block w-2.5 h-2.5 bg-white scale-0 duration-100;
				@apply absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2;
				clip-path: polygon(14% 44%, 0 65%, 50% 100%, 100% 16%, 80% 0%, 43% 62%);
			}

			&:checked {
				@apply bg-blue-600;
				&::before {
					@apply scale-100;
				}
			}
		}
	}

	.action-btn {
		@apply flex flex-row items-center;
	}
</style>
