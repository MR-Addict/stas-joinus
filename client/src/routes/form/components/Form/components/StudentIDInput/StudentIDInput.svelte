<script lang="ts">
	import '../form.css';
	import { onMount } from 'svelte';
	import inputs from '$stores/inputs';

	let error = '';
	$: error = $inputs.find((input) => input.id === 'student_id')?.error || '';

	onMount(() => inputs.register('student_id', validate));

	function handleChange(event: Event) {
		const value = (event.target as HTMLInputElement).value;
		inputs.update('student_id', value);
	}

	function validate(value: string) {
		if (value.length === 0) return '学号不能为空';
		else if (value.length != 12) return '学号应该是12位';
		else return '';
	}
</script>

<div class="form" class:error>
	<h1 class="title">
		<span>学号</span>
		<span class="text-red-600">*</span>
	</h1>
	<label for="student_id" class="label">你的学号</label>
	<input
		id="student_id"
		name="student_id"
		type="number"
		class="input"
		placeholder="输入你的学号"
		on:change={handleChange}
	/>
	<p class="err-msg">{error}</p>
</div>
