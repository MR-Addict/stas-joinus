<script lang="ts">
	import { onMount } from 'svelte';
	import { CreditCard } from 'lucide-svelte';

	import '../form.css';
	import form from '$stores/form';
	import inputs from '$stores/inputs';

	const id = 'student_id';

	let error = '';
	let ref: HTMLDivElement;
	let value = $form.localApplicant?.student_id || '';

	$: error = $inputs.find((input) => input.id === id)?.error || '';

	onMount(() => inputs.register(id, value, ref, validate));

	function handleChange() {
		inputs.update(id, value);
	}

	function validate(value: string) {
		if (value.length === 0) return '学号不能为空';
		else if (value.length != 12) return '学号应该是12位';
		else return '';
	}
</script>

<div bind:this={ref} class="form" class:error>
	<h1 class="title">
		<CreditCard size={18} />
		<span>学号</span>
	</h1>
	<label for={id} class="label">你的学号</label>
	<input
		{id}
		name={id}
		type="number"
		class="input"
		placeholder="你的学号"
		disabled={$form.localApplicant !== null}
		bind:value
		on:change={handleChange}
	/>
	<p class="err-msg">{error}</p>
</div>
