<script lang="ts">
	import toast from 'svelte-french-toast';
	import { LockKeyhole } from 'lucide-svelte';

	import auth from '$stores/auth';

	import Modal from '$components/Modal/Modal.svelte';
	import Spinner from '$components/Spinner/Spinner.svelte';

	export let showModal = false;

	let pending = false;

	async function handleSubmit(event: SubmitEvent) {
		const formData = new FormData(event.target as HTMLFormElement);

		pending = true;

		const loginRes = await auth.login(formData);
		if (!loginRes.success) toast.error(loginRes.message);
		else showModal = false;

		pending = false;
	}
</script>

<Modal bind:showModal disabled={pending}>
	<form on:submit|preventDefault={handleSubmit}>
		<h1>用户登录</h1>
		<p>非管理员请勿尝试登录</p>

		<div class="input">
			<div class="text-gray-500"><LockKeyhole size={16} /></div>
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

<style>
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
		&:enabled:hover {
			@apply bg-teal-800;
		}
		&:disabled {
			@apply bg-gray-300 text-gray-400 cursor-not-allowed;
		}
	}
</style>
