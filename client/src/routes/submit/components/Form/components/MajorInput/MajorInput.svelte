<script lang="ts">
	import { onMount } from 'svelte';
	import { GraduationCap } from 'lucide-svelte';

	import '../form.css';
	import form from '$stores/form';
	import inputs from '$stores/inputs';

	const id = 'major';

	let error = '';
	let ref: HTMLDivElement;
	let value = $form.localApplicant?.major || '';

	$: error = $inputs.find((input) => input.id === id)?.error || '';

	onMount(() => inputs.register(id, value, ref, validate));

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
	<label for={id}>
		<h1>
			<GraduationCap size={18} />
			<span>专业</span>
		</h1>
		<p>你所学的专业</p>
	</label>
	<input
		{id}
		name={id}
		type="text"
		class="input"
		placeholder="你的专业"
		disabled={$form.localApplicant !== null}
		bind:value
		on:change={handleChange}
	/>
	<p class="err-msg">{error}</p>
</div>
