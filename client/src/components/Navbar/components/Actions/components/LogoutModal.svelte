<script lang="ts">
	import toast from 'svelte-french-toast';
	import { goto } from '$app/navigation';
	import { page } from '$app/state';

	import auth from '$stores/auth';
	import Modal from '$components/Modal/Modal.svelte';
	import Spinner from '$components/Spinner/Spinner.svelte';

	export let showModal = false;

	let pending = false;

	async function handleLogout() {
		pending = true;

		const res = await auth.logout();
		if (res.success) {
			showModal = false;
			toast.success(res.message);
			if (['/stats/', '/view/'].includes(page.url.pathname)) setTimeout(() => goto('/'), 0);
		} else toast.error(res.message);

		pending = false;
	}
</script>

<Modal bind:showModal disabled={pending}>
	<form on:submit|preventDefault={handleLogout}>
		<h1>确认退出登录？</h1>

		<p>退出后将无法访问管理页面</p>

		<button disabled={pending} type="submit">
			<span>退出登录</span>
			{#if pending}
				<Spinner />
			{/if}
		</button>
	</form>
</Modal>

<style>
	form {
		@apply flex flex-col gap-2;
		@apply bg-white p-7 w-full max-w-xs rounded-lg text-center;
	}
	h1 {
		@apply font-semibold text-xl;
	}
	p {
		@apply text-gray-600 text-sm;
	}
	button {
		@apply mt-2 flex items-center justify-center gap-2;
		@apply py-2.5 px-4 rounded-md duration-300 text-white bg-teal-600;

		&:enabled:hover {
			@apply bg-teal-700;
		}
		&:disabled {
			@apply bg-gray-300 text-gray-400 cursor-not-allowed;
		}
	}
</style>
