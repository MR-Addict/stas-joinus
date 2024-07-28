<script lang="ts">
	import { onMount } from 'svelte';
	import { NotebookPen } from 'lucide-svelte';

	import '../form.css';
	import inputs from '$stores/inputs';

	const id = 'introduction';

	let value = '';
	let error = '';
	let ref: HTMLDivElement;

	$: error = $inputs.find((input) => input.id === id)?.error || '';

	onMount(() => inputs.register(id, ref, validate));

	function handleChange() {
		inputs.update(id, value);
	}

	function validate(value: string) {
		if (value.length === 0) return '自我介绍不能为空';
		else if (value.length < 4) return '你的自我介绍太短啦';
		else if (value.length > 500) return '你的自我介绍太长啦';
		else return '';
	}
</script>

<div bind:this={ref} class="form" class:error>
	<h1 class="title">
		<NotebookPen size={18} />
		<span>自我介绍</span>
	</h1>
	<label for={id} class="label">
		写写你的性格、爱好、特长、个人项目、取得的成就等等，尽量内容丰富，让我们对你有更好的了解
	</label>
	<textarea {id} name={id} class="input h-40" placeholder="你的自我介绍" bind:value on:change={handleChange} />
	<p class="err-msg">{error}</p>
</div>
