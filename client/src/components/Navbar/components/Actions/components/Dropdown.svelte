<script lang="ts">
	import { ChartArea, GripVertical, ArrowDownToLine, LogOut, Table, User } from 'lucide-svelte';

	import auth from '$stores/auth';
	import clickOutside from '$hooks/clickOutside';

	export let showModal = { login: false, logout: false, export: false };

	let showDropdown = false;

	function closeDropwdown() {
		if ($auth) showDropdown = false;
	}

	function showLogout() {
		showDropdown = false;
		showModal.logout = true;
	}

	function showExport() {
		showDropdown = false;
		showModal.export = true;
	}

	function handleClick() {
		if (!$auth) showModal.login = true;
		else showDropdown = !showDropdown;
	}
</script>

<div class="menu-wrapper" use:clickOutside={closeDropwdown}>
	<button
		type="button"
		aria-label="action button"
		class="text-gray-600 hover:bg-black/10 rounded-full p-1 duration-300"
		on:click={handleClick}
	>
		{#if $auth}
			<GripVertical />
		{:else}
			<User />
		{/if}
	</button>

	{#if $auth}
		<div class="menu" class:active={showDropdown}>
			<a href="/view" class="btn" on:click={closeDropwdown}>
				<span>报名数据</span>
				<span class="icon"><Table size={16} /></span>
			</a>

			<a href="/stats" class="btn" on:click={closeDropwdown}>
				<span>报名统计</span>
				<span class="icon"><ChartArea size={16} /></span>
			</a>

			<hr />

			<button type="button" class="btn" on:click={showExport}>
				<span>导出数据</span>
				<span class="icon"><ArrowDownToLine size={16} /></span>
			</button>

			<button type="button" class="btn" on:click={showLogout}>
				<span>退出登录</span>
				<span class="icon"><LogOut size={16} /></span>
			</button>
		</div>
	{/if}
</div>

<style>
	.menu-wrapper {
		@apply relative flex items-center justify-center;
	}

	.menu {
		box-shadow: 0 0 5px #dbdee2;
		@apply w-48 absolute -bottom-2 right-1.5 translate-y-full;
		@apply origin-top-right duration-300 scale-50 opacity-0 invisible;
		@apply bg-white border border-gray-300 rounded-md py-1;

		&.active {
			@apply scale-100 opacity-100 visible;
		}
	}

	.btn {
		@apply flex flex-row items-center w-full py-2 px-4;

		&:not(:last-child) {
			@apply border-b border-gray-300;
		}

		&:hover {
			@apply bg-teal-700 text-white duration-200;

			& .icon {
				@apply text-white;
			}
		}

		& .icon {
			@apply ml-auto text-gray-600;
		}
	}

	hr {
		@apply h-1 bg-gray-300;
	}
</style>
