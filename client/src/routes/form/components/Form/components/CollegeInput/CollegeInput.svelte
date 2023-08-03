<script lang="ts">
	import '../form.css';
	import { onMount } from 'svelte';
	import inputs from '$stores/inputs';

	let error = '';
	$: error = $inputs.find((input) => input.id === 'college')?.error || '';

	onMount(() => inputs.register('college', validate));

	function handleChange(event: Event) {
		const value = (event.target as HTMLInputElement).value;
		inputs.update('college', value);
	}

	function validate(value: string) {
		if (value.length === 0) return '学院不能为空';
		else if (value.length > 50) return '你的学院名字太长啦';
		else return '';
	}
</script>

<div class="form" class:error>
	<h1 class="title">
		<span>学院</span>
		<span class="text-red-600">*</span>
	</h1>
	<label for="college" class="label">你所在的学院</label>
	<input id="college" name="college" type="text" class="input" placeholder="输入的学院" on:change={handleChange} />
	<p class="err-msg">{error}</p>
</div>
