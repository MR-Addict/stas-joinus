<script lang="ts">
	import { onMount } from 'svelte';
	import { User } from 'lucide-svelte';

	import '../form.css';
	import inputs from '$stores/inputs';

	const id = 'name';

	let value = '';
	let error = '';
	let ref: HTMLDivElement;

	$: error = $inputs.find((input) => input.id === id)?.error || '';

	onMount(() => inputs.register(id, ref, validate));

	function handleChange() {
		inputs.update(id, value);
	}

	function validate(value: string) {
		if (value.length === 0) return '姓名不能为空';
		else if (value.length < 2) return '你的名字太短啦';
		else if (value.length > 20) return '你的名字太长啦';
		else return '';
	}
</script>

<div bind:this={ref} class="form" class:error>
	<h1 class="title">
		<User size={18} />
		<span>姓名</span>
	</h1>
	<label for={id} class="label">你的姓名</label>
	<input {id} name={id} type="text" class="input" placeholder="你的姓名" bind:value on:change={handleChange} />
	<p class="err-msg">{error}</p>
</div>
