<script lang="ts">
	import '../form.css';
	import { onMount } from 'svelte';
	import inputs from '$stores/inputs';

	let error = '';
	$: error = $inputs.find((input) => input.id === 'phone')?.error || '';

	onMount(() => inputs.register('phone', validate));

	function handleChange(event: Event) {
		const value = (event.target as HTMLInputElement).value;
		inputs.update('phone', value);
	}

	function validate(value: string) {
		if (value.length === 0) return '手机号不能为空';
		else if (value.length !== 11) return '手机号应该是11位';
		else return '';
	}
</script>

<div class="form">
	<h1 class="title">
		<span>手机</span>
		<span class="text-red-600">*</span>
	</h1>
	<label for="phone" class="label">我们会将面试信息以及面试结果通过短信的形式发送给你</label>
	<input id="phone" name="phone" type="number" class="input" placeholder="输入你的手机号" on:change={handleChange} />
	<p class="err-msg" class:active={error}>{error}</p>
</div>
