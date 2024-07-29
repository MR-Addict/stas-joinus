<script lang="ts">
	import form from '$stores/form';

	import Modal from '$components/Modal/Modal.svelte';

	export let showModal: boolean;

	function handleConfirm() {
		showModal = false;
		form.updateShowForm(true);
		form.refreshLocalApplicant();
	}
</script>

<Modal bind:showModal>
	<div class="wrapper">
		<h1>确认修改志愿</h1>
		<p>每人仅有<strong>一次</strong>修改志愿的机会，并且只能修改<strong>志愿</strong>和<strong>自我介绍</strong>部分</p>

		<div class="btns">
			<button type="button" on:click={() => (showModal = false)}>取消</button>
			<button type="button" on:click={handleConfirm}>
				<span>修改</span>
			</button>
		</div>
	</div>
</Modal>

<style lang="postcss">
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

			&:first-of-type:hover {
				@apply bg-gray-200;
			}

			&:last-of-type {
				@apply flex flex-row items-center justify-center gap-1 bg-yellow-400;

				&:enabled:hover {
					@apply bg-yellow-500;
				}
				&:disabled {
					@apply bg-gray-300 text-gray-400 cursor-not-allowed;
				}
			}
		}
	}
</style>
