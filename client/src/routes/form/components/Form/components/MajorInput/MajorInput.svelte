<script lang="ts">
	import '../form.css';
	import { onMount } from 'svelte';
	import inputs from '$stores/inputs';

	let error = '';
	$: error = $inputs.find((input) => input.id === 'major')?.error || '';

	onMount(() => inputs.register('major', validate));

	function handleChange(event: Event) {
		const value = (event.target as HTMLInputElement).value;
		inputs.update('major', value);
	}

	function validate(value: string) {
		if (value.length === 0) return '专业不能为空';
		else if (value.length > 50) return '你的专业名字太长啦';
		else return '';
	}
</script>

<div class="form" class:error>
	<h1 class="title">
		<span>专业</span>
		<span class="text-red-600">*</span>
	</h1>
	<label for="major" class="label">你现在的专业</label>
	<input id="major" name="major" type="text" class="input" placeholder="输入你的专业" on:change={handleChange} />
	<p class="err-msg">{error}</p>
</div>
