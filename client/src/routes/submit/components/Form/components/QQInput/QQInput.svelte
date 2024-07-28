<script lang="ts">
	import { onMount } from 'svelte';
	import { PawPrint } from 'lucide-svelte';

	import '../form.css';
	import inputs from '$stores/inputs';

	const id = 'qq';

	let value = '';
	let error = '';
	let ref: HTMLDivElement;

	$: error = $inputs.find((input) => input.id === id)?.error || '';

	onMount(() => inputs.register(id, ref, validate));

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
	<h1 class="title">
		<PawPrint size={18} />
		<span>QQ</span>
	</h1>
	<label for={id} class="label">我们联系你的备用方案</label>
	<input {id} name={id} type="number" class="input" placeholder="你的QQ号" bind:value on:change={handleChange} />
	<p class="err-msg">{error}</p>
</div>
