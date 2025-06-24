<script lang="ts">
	import Modal from '$components/Modal/Modal.svelte';
	import Login from '$components/Login/Login.svelte';
	import type { ApiResultType } from '$types/app';

	export let showModal = false;

	let pending = false;

	function onSubmit(res: ApiResultType) {
		if (res.success) showModal = false;
	}
</script>

<Modal bind:showModal disabled={pending}>
	<Login bind:pending {onSubmit} />
</Modal>
