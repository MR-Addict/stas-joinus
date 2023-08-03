<script lang="ts">
	import '../form.css';
	import { onMount } from 'svelte';
	import inputs from '$stores/inputs';

	let error = '';
	$: error = $inputs.find((input) => input.id === 'email')?.error || '';

	onMount(() => inputs.register('email', validate));

	function handleChange(event: Event) {
		const value = (event.target as HTMLInputElement).value;
		inputs.update('email', value);
	}

	function validate(value: string) {
		if (value.length === 0) return '邮箱不能为空';
		else if (value.length > 320) return '你的邮箱太长啦';
		else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(value)) return '请输入正确格式的邮箱';
		else return '';
	}
</script>

<div class="form" class:error>
	<h1 class="title">
		<span>邮箱</span>
		<span class="text-red-600">*</span>
	</h1>
	<label for="email" class="label">我们会将面试信息以及面试结果通过邮件的形式发送给你</label>
	<input id="email" name="email" type="email" class="input" placeholder="输入你的邮箱" on:change={handleChange} />
	<p class="err-msg">{error}</p>
</div>
