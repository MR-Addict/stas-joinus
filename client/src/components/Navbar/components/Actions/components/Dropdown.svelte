<script lang="ts">
	import toast from 'svelte-french-toast';
	import { goto } from '$app/navigation';
	import { ChartArea, GripVertical, LogOut, Table, User } from 'lucide-svelte';

	import auth from '$stores/auth';

	import Spinner from '$components/Spinner/Spinner.svelte';

	export let showModal = false;

	let pending = false;

	async function handleLogout() {
		pending = true;

		const res = await auth.logout();
		if (res.success) goto('/');
		else toast.error(res.message);

		pending = false;
	}
</script>

<div class="menu-wrapper">
	<button
		type="button"
		class="text-gray-600 hover:bg-black/10 rounded-full p-1 duration-300"
		on:click={() => !$auth && (showModal = true)}
	>
		{#if $auth}
			<User />
		{:else}
			<GripVertical />
		{/if}
	</button>

	{#if $auth}
		<div class="menu">
			<a href="/table" class="btn">
				<span>报名数据</span>
				<span class="icon">
					<Table size={16} />
				</span>
			</a>

			<a href="/chart" class="btn">
				<span>报名统计</span>
				<span class="icon">
					<ChartArea size={16} />
				</span>
			</a>

			<button type="button" class="btn" on:click={handleLogout} disabled={pending}>
				{#if pending}
					<Spinner />
				{/if}
				<span>退出登录</span>
				<span class="icon">
					<LogOut size={16} />
				</span>
			</button>
		</div>
	{/if}
</div>

<style lang="postcss">
	.menu-wrapper {
		@apply relative flex items-center justify-center;

		&:hover .menu {
			@apply scale-100 opacity-100 visible;
		}
	}

	.menu {
		@apply delay-100 origin-top-right duration-300 scale-50 opacity-0 invisible;
		@apply w-36 absolute -bottom-3 right-0 translate-y-full;
		@apply shadow-md bg-white border border-gray-300 rounded-md py-1;
	}

	.btn {
		@apply flex flex-row items-center w-full py-1.5 px-4 text-sm;

		&:not(:last-child) {
			@apply border-b-[3px] border-gray-300;
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
</style>
