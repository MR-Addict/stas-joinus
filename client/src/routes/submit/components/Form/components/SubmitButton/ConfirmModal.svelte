<script lang="ts">
	import Modal from '$components/Modal/Modal.svelte';
	import Spinner from '$components/Spinner/Spinner.svelte';
	import form from '$stores/form';

	export let pending: boolean;
	export let showModal: boolean;
</script>

<Modal bind:showModal disabled={pending}>
	<div class="wrapper">
		<h1>{`确认${$form.localApplicant === null ? '提交' : '修改'}`}</h1>

		{#if $form.localApplicant}
			<p>每人仅有一次修改机会，提交后<strong>志愿选择</strong>和<strong>自我介绍</strong>等所有内容将不可修改</p>
		{:else}
			<p>
				请确认信息填写正确，提交后<strong>个人信息</strong>、<strong>联系方式</strong>和<strong>学院及专业</strong
				>这些内容将不可修改
			</p>
		{/if}

		<div class="btns">
			<button type="button" on:click={() => (showModal = false)} disabled={pending}>取消</button>
			<button type="submit" disabled={pending}>
				{#if pending}
					<Spinner />
				{/if}
				<span>{$form.localApplicant === null ? '提交' : '修改'}</span>
			</button>
		</div>
	</div>
</Modal>

<style>
	.wrapper {
		@apply max-w-sm bg-white p-7 sm:p-10 flex flex-col items-center rounded-xl gap-2;
	}

	h1 {
		@apply text-xl font-semibold;
	}

	p {
		text-wrap: balance;
		@apply text-center text-gray-800;

		& strong {
			@apply text-black;
		}
	}

	.btns {
		@apply mt-2 w-full flex flex-row items-center justify-center gap-4;

		& button {
			@apply py-2 w-full border-2 border-black rounded-md font-semibold duration-300;

			&:first-of-type:enabled:hover {
				@apply bg-gray-200;
			}

			&:last-of-type {
				@apply flex flex-row items-center justify-center gap-1 bg-yellow-400;

				&:enabled:hover {
					@apply bg-yellow-500;
				}
			}

			&:disabled {
				@apply bg-gray-300 text-gray-400 cursor-not-allowed;
			}
		}
	}
</style>
