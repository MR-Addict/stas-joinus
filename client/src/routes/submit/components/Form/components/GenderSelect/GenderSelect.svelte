<script lang="ts">
	import { onMount } from 'svelte';
	import { Baby } from 'lucide-svelte';

	import '../form.css';
	import form from '$stores/form';
	import inputs from '$stores/inputs';

	const id = 'gender';

	let error = '';
	let ref: HTMLDivElement;
	let value = $form.localApplicant?.gender || '';

	$: error = $inputs.find((input) => input.id === id)?.error || '';

	onMount(() => inputs.register(id, value, ref, validate));

	function handleChange() {
		inputs.update(id, value);
	}

	function validate(value: string) {
		if (value.length === 0) return '请选择你的性别';
		else return '';
	}
</script>

<div bind:this={ref} class="form" class:error>
	<h1 class="title">
		<Baby size={18} />
		<span>性别</span>
	</h1>
	<label for={id} class="label">你的性别</label>
	<select {id} name={id} class="input" disabled={$form.localApplicant !== null} bind:value on:change={handleChange}>
		<option disabled value="">--选择性别--</option>
		<option value="boy">男</option>
		<option value="girl">女</option>
	</select>
	<p class="err-msg">{error}</p>
</div>
