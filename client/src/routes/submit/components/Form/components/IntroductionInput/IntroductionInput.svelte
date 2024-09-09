<script lang="ts">
	import { onMount } from 'svelte';
	import { NotebookPen } from 'lucide-svelte';

	import '../form.css';
	import form from '$stores/form';
	import inputs from '$stores/inputs';

	const id = 'introduction';

	let error = '';
	let ref: HTMLDivElement;
	let value = $form.localApplicant?.introduction || '';

	$: error = $inputs.find((input) => input.id === id)?.error || '';

	onMount(() => inputs.register(id, value, ref, validate));

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
	<label for={id}>
		<h1>
			<NotebookPen size={18} />
			<span>自我介绍</span>
		</h1>
		<p>写写你的性格、爱好、特长、获奖经历、班级职务等等，内容尽可能丰富具体，让我们对你有更好的了解</p>
	</label>
	<textarea
		{id}
		name={id}
		class="input h-64 sm:h-48"
		placeholder="你的自我介绍"
		disabled={$form.localApplicant?.modified}
		bind:value
		on:change={handleChange}
	/>
	<p class="err-msg">{error}</p>
</div>
