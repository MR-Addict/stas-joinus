<script lang="ts">
	import auth from '$stores/auth';
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import { toasts } from '$stores/toasts';
	import clickOutside from '$hooks/clickOutside';

	import Login from '../Login/Login.svelte';
	import User from '$components/Icons/User/User.svelte';
	import Spinner from '$components/Spinner/Spinner.svelte';

	let pending = false;
	let showMenu = false;
	let showModal = false;
	async function handleLogout() {
		pending = true;
		if (await auth.logout()) {
			showMenu = false;
			const pathname = $page.url.pathname.split('/').slice(0, 2).join('/');
			if (pathname === '/table' || pathname === '/analytics') goto('/');
		} else toasts.add('退出失败，请稍后重试或联系我们', 'failed');
		pending = false;
	}
</script>

<Login bind:showModal />

<div use:clickOutside={() => (showMenu = false)} class="wrapper">
	<button type="button" aria-label="profile" class="avatar" on:click={() => (showMenu = !showMenu)}>
		<User />
	</button>

	<div class="menu-wrapper" class:active={showMenu}>
		<div class="menu">
			{#if $auth}
				<button disabled={pending} type="button" class="btn" on:click={handleLogout}>
					{#if pending}
						<Spinner />
					{/if}
					<span>退出登录</span>
				</button>
			{:else}
				<button type="button" class="btn" on:click={() => (showModal = true) && (showMenu = false)}>登录后台</button>
			{/if}
		</div>
	</div>
</div>

<style lang="postcss">
	.wrapper {
		@apply relative flex items-center justify-center;
	}
	.avatar {
		@apply w-6 text-gray-700 duration-300;
		&:hover {
			@apply text-black;
		}
	}
	.menu-wrapper {
		@apply hidden absolute top-full right-0 md:right-1/2 md:translate-x-1/2;
		&.active {
			@apply block;
		}
	}
	.menu {
		box-shadow: 0 0 5px #d1d5db;
		@apply mt-1.5 bg-white py-3 px-5 rounded-md border border-gray-300;
	}
	.btn {
		@apply flex flex-row items-center gap-1 whitespace-nowrap text-sm text-gray-700;
		&:hover {
			@apply text-black;
		}
		&:disabled {
			@apply text-gray-500 cursor-not-allowed;
		}
	}
</style>
