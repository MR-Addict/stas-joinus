<script lang="ts">
	import { onMount } from 'svelte';
	import { Mail } from 'lucide-svelte';

	import '../form.css';
	import inputs from '$stores/inputs';

	const id = 'email';

	let value = '';
	let error = '';
	let ref: HTMLDivElement;

	$: error = $inputs.find((input) => input.id === id)?.error || '';

	onMount(() => inputs.register(id, ref, validate));

	function handleChange() {
		inputs.update(id, value);
	}

	function validate(value: string) {
		if (value.length === 0) return '邮箱不能为空';
		else if (value.length > 320) return '你的邮箱太长啦';
		else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(value)) return '请输入正确格式的邮箱';
		else return '';
	}
</script>

<div bind:this={ref} class="form" class:error>
	<h1 class="title">
		<Mail size={18} />
		<span>邮箱</span>
	</h1>
	<label for={id} class="label">我们会将面试信息及面试结果通过邮件的形式发送给你</label>
	<input {id} name={id} type="email" class="input" placeholder="你的邮箱" bind:value on:change={handleChange} />
	<p class="err-msg">{error}</p>
</div>
