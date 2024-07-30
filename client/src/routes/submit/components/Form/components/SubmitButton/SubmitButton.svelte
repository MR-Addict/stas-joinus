<script lang="ts">
	import form from '$stores/form';
	import inputs from '$stores/inputs';

	import ConfirmModal from './ConfirmModal.svelte';

	export let pending: boolean;

	let showModal = false;

	function handleCancel() {
		window.scroll({ top: 0 });
		form.updateShowForm(false);
	}

	function handleConfirm() {
		if (inputs.checkHasErrors()) inputs.scrollToFirstError();
		else showModal = true;
	}
</script>

<ConfirmModal {pending} bind:showModal />

<div class="btns">
	{#if $form.localApplicant}
		<button type="button" on:click={handleCancel}>返回</button>
	{/if}

	{#if !$form.localApplicant?.modified}
		<button data-type="action" type="button" on:click={handleConfirm}>
			{$form.localApplicant === null ? '提交' : '修改'}
		</button>
	{/if}
</div>

<style>
	.btns {
		@apply w-full flex flex-row items-center justify-center gap-4;

		& button {
			@apply py-3 w-full border-2 border-black rounded-md font-semibold duration-300;

			&:first-of-type:hover {
				@apply bg-gray-200;
			}

			&[data-type='action'] {
				@apply flex flex-row items-center justify-center gap-1 bg-yellow-400;

				&:hover {
					@apply bg-yellow-500;
				}
			}
		}
	}
</style>
