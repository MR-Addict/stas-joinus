<script lang="ts">
	import { onMount } from 'svelte';
	import { GraduationCap } from 'lucide-svelte';

	import '../form.css';
	import inputs from '$stores/inputs';

	const id = 'major';

	let value = '';
	let error = '';
	let ref: HTMLDivElement;

	$: error = $inputs.find((input) => input.id === id)?.error || '';

	onMount(() => inputs.register(id, ref, validate));

	function handleChange() {
		inputs.update(id, value);
	}

	function validate(value: string) {
		if (value.length === 0) return '专业不能为空';
		else if (value.length > 50) return '你的专业名字太长啦';
		else return '';
	}
</script>

<div bind:this={ref} class="form" class:error>
	<h1 class="title">
		<GraduationCap size={18} />
		<span>专业</span>
	</h1>
	<label for={id} class="label">你所学的专业</label>
	<input {id} name={id} type="text" class="input" placeholder="你的专业" bind:value on:change={handleChange} />
	<p class="err-msg">{error}</p>
</div>
