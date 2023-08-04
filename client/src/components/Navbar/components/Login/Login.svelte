<script lang="ts">
	import auth from '$stores/auth';
	import url from '$lib/utils/url';
	import toasts from '$stores/toasts';
	import { Response } from '$types/response';

	import Modal from '$components/Modal/Modal.svelte';
	import Spinner from '$components/Spinner/Spinner.svelte';
	import MdLockOutline from 'svelte-icons/md/MdLockOutline.svelte';

	export let showModal = false;

	let pending = false;

	async function handleSubmit(event: SubmitEvent) {
		const formData = new FormData(event.target as HTMLFormElement);

		pending = true;
		try {
			const res = await fetch(url('/api/user/login'), { method: 'POST', body: formData }).then((res) => res.json());
			const { success, message } = Response.parse(res);
			if (!success) toasts.add(message, 'failed');
			else if (await auth.ping()) showModal = false;
			else toasts.add('登录失败，请稍后尝试或联系我们', 'failed');
		} catch (err) {
			console.error(err);
			toasts.add('登录失败，请稍后尝试或联系我们', 'failed');
		} finally {
			pending = false;
		}
	}
</script>

<Modal bind:showModal>
	<form on:submit|preventDefault={handleSubmit}>
		<h1>用户登录</h1>
		<p>非管理员请勿尝试登录</p>
		<div class="input">
			<div class="w-5 text-gray-500"><MdLockOutline /></div>
			<input required type="password" name="password" placeholder="输入密码" />
		</div>
		<button disabled={pending} type="submit">
			{#if pending}
				<Spinner />
			{/if}
			<span>登录</span>
		</button>
	</form>
</Modal>

<style lang="postcss">
	form {
		@apply flex flex-col items-center gap-2;
		@apply bg-white p-7 w-full max-w-xs rounded-lg text-center;
	}
	h1 {
		@apply font-semibold text-xl;
	}
	p {
		@apply text-sm text-gray-500;
	}
	.input {
		@apply bg-white w-full flex flex-row items-center gap-2;
		@apply py-2 px-2 rounded-md border border-gray-400 outline-none;
		& input {
			@apply w-full outline-none;
		}
		&:focus-within {
			@apply outline-teal-600;
		}
	}
	button {
		@apply flex flex-row items-center justify-center gap-1;
		@apply mt-2 w-full py-2 rounded-md bg-teal-700 text-white duration-300;
		&:hover {
			@apply bg-teal-800;
		}
		&:disabled {
			@apply bg-gray-200 cursor-not-allowed;
		}
	}
</style>
