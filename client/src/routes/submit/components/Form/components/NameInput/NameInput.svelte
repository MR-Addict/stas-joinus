<script lang="ts">
	import { onMount } from 'svelte';
	import { User } from 'lucide-svelte';

	import '../form.css';
	import form from '$stores/form';
	import inputs from '$stores/inputs';

	const id = 'name';

	let error = '';
	let ref: HTMLDivElement;
	let value = $form.localApplicant?.name || '';

	$: error = $inputs.find((input) => input.id === id)?.error || '';

	onMount(() => inputs.register(id, value, ref, validate));

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
	<label for={id}>
		<h1>
			<User size={18} />
			<span>姓名</span>
		</h1>
		<p>你的姓名</p>
	</label>
	<input
		{id}
		name={id}
		type="text"
		class="input"
		placeholder="你的姓名"
		disabled={$form.localApplicant !== null}
		bind:value
		on:change={handleChange}
	/>
	<p class="err-msg">{error}</p>
</div>
