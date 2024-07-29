<script lang="ts">
	import { onMount } from 'svelte';
	import { Smartphone } from 'lucide-svelte';

	import '../form.css';
	import form from '$stores/form';
	import inputs from '$stores/inputs';

	const id = 'phone';

	let error = '';
	let ref: HTMLDivElement;
	let value = $form.localApplicant?.phone || '';

	$: error = $inputs.find((input) => input.id === id)?.error || '';

	onMount(() => inputs.register(id, value, ref, validate));

	function handleChange() {
		inputs.update(id, value);
	}

	function validate(value: string) {
		if (value.length === 0) return '手机号不能为空';
		else if (value.length !== 11) return '手机号应该是11位';
		else return '';
	}
</script>

<div bind:this={ref} class="form" class:error>
	<h1 class="title">
		<Smartphone size={18} />
		<span>手机</span>
	</h1>
	<label for={id} class="label">我们会将面试信息及面试结果通过短信的形式发送给你</label>
	<input
		{id}
		name={id}
		type="number"
		class="input"
		placeholder="你的手机号"
		disabled={$form.localApplicant !== null}
		bind:value
		on:change={handleChange}
	/>
	<p class="err-msg">{error}</p>
</div>
