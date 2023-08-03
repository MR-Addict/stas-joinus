<script lang="ts">
	import auth from '$stores/auth';
	import { toasts } from '$stores/toasts';

	import Login from '../Login/Login.svelte';
	import User from '$components/Icons/User/User.svelte';
	import Spinner from '$components/Spinner/Spinner.svelte';

	let pending = false;
	let showModal = false;

	async function handleLogout() {
		pending = true;
		const success = await auth.logout();
		if (!success) toasts.add('退出失败，请稍后重试或联系我们', 'failed');
		pending = false;
	}
</script>

<Login bind:showModal />

<div class="wrapper">
	<button type="button" aria-label="profile" class="w-6 text-gray-700 duration-300 hover:text-black">
		<User />
	</button>

	<div class="menu-wrapper">
		<div class="menu">
			{#if $auth}
				<button disabled={pending} type="button" class="btn" on:click={handleLogout}>
					{#if pending}
						<Spinner />
					{/if}
					<span>退出登录</span>
				</button>
			{:else}
				<button type="button" class="btn" on:click={() => (showModal = true)}>登录后台</button>
			{/if}
		</div>
	</div>
</div>

<style lang="postcss">
	.wrapper {
		@apply relative flex items-center justify-center;
		&:hover .menu-wrapper {
			@apply block;
		}
	}
	.menu-wrapper {
		@apply hidden absolute top-full right-0 md:right-1/2 md:translate-x-1/2;
	}
	.menu {
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
