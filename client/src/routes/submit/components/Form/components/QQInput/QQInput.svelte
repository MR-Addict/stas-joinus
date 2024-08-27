<script lang="ts">
	import { onMount } from 'svelte';
	import { PawPrint } from 'lucide-svelte';

	import '../form.css';
	import form from '$stores/form';
	import inputs from '$stores/inputs';

	const id = 'qq';

	let error = '';
	let ref: HTMLDivElement;
	let value = $form.localApplicant?.qq || '';

	$: error = $inputs.find((input) => input.id === id)?.error || '';

	onMount(() => inputs.register(id, value, ref, validate));

	function handleChange() {
		inputs.update(id, value);
	}

	function validate(value: string) {
		if (value.length === 0) return 'QQ号不能为空';
		else if (value.length < 5) return '你的QQ号太短啦';
		else if (value.length > 11) return '你的QQ号太长啦';
		else return '';
	}
</script>

<div bind:this={ref} class="form" class:error>
	<label for={id}>
		<h1>
			<PawPrint size={18} />
			<span>QQ</span>
		</h1>
		<p>我们联系你的备用方案</p>
	</label>
	<input
		{id}
		name={id}
		type="number"
		class="input"
		placeholder="你的QQ号"
		disabled={$form.localApplicant !== null}
		bind:value
		on:change={handleChange}
	/>
	<p class="err-msg">{error}</p>
</div>
