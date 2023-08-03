<script lang="ts">
	import '../form.css';
	import { onMount } from 'svelte';
	import inputs from '$stores/inputs';

	let error = '';
	$: error = $inputs.find((input) => input.id === 'name')?.error || '';

	onMount(() => inputs.register('name', validate));

	function handleChange(event: Event) {
		const value = (event.target as HTMLInputElement).value;
		inputs.update('name', value);
	}

	function validate(value: string) {
		if (value.length === 0) return '姓名不能为空';
		else if (value.length < 2) return '你的名字太短啦';
		else if (value.length > 20) return '你的名字太长啦';
		else return '';
	}
</script>

<div class="form">
	<h1 class="title">
		<span>姓名</span>
		<span class="text-red-600">*</span>
	</h1>
	<label for="name" class="label">你的姓名</label>
	<input id="name" name="name" type="text" class="input" placeholder="输入你的姓名" on:change={handleChange} />
	<p class="err-msg" class:active={error}>{error}</p>
</div>
