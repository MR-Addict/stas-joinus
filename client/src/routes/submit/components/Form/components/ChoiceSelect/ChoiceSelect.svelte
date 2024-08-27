<script lang="ts">
	import { onMount } from 'svelte';
	import { FlaskConical } from 'lucide-svelte';

	import '../form.css';
	import form from '$stores/form';
	import inputs from '$stores/inputs';

	const first_label = 'first_choice';
	const second_label = 'second_choice';

	let first_error = '';
	let second_error = '';
	let firstRef: HTMLDivElement;
	let secondRef: HTMLDivElement;
	let first_choice = $form.localApplicant?.first_choice || '';
	let second_choice = $form.localApplicant?.second_choice || '';

	$: first_error = $inputs.find((input) => input.id === first_label)?.error || '';
	$: second_error = $inputs.find((input) => input.id === second_label)?.error || '';

	onMount(() => {
		inputs.register(first_label, first_choice, firstRef, validateFirst);
		inputs.register(second_label, second_choice, secondRef, validateSecond);
	});

	function handleFirstChange() {
		inputs.update(first_label, first_choice);
	}

	function handleSecondChange() {
		inputs.update(second_label, second_choice);
	}

	function validateFirst(value: string) {
		if (!value) return '请选择第一志愿';
		else if (value === second_choice) return '志愿重复，请修改第一或第二志愿';
		else return '';
	}

	function validateSecond(value: string) {
		if (!value) return '请选择第二志愿';
		else if (value === first_choice) return '志愿重复，请修改第一或第二志愿';
		else return '';
	}
</script>

<div bind:this={firstRef} class="form" class:error={first_error}>
	<label for={first_label}>
		<h1>
			<FlaskConical size={18} />
			<span>第一志愿</span>
		</h1>
		<p>校大学生科协你最想加入的部门，我们会优先考虑你的第一志愿</p>
	</label>
	<select
		id={first_label}
		name={first_label}
		class="input"
		disabled={$form.localApplicant?.modified}
		bind:value={first_choice}
		on:change={handleFirstChange}
	>
		<option disabled value="">--选择部门--</option>
		<option value="技术开发部">技术开发部</option>
		<option value="科普活动部">科普活动部</option>
		<option value="组织策划部">组织策划部</option>
		<option value="新闻宣传部">新闻宣传部</option>
		<option value="对外联络部">对外联络部</option>
		<option value="双创联合服务部">双创联合服务部</option>
	</select>
	<p class="err-msg">{first_error}</p>
</div>

<div bind:this={secondRef} class="form" class:error={second_error}>
	<label for={second_label}>
		<h1>
			<FlaskConical size={18} />
			<span>第二志愿</span>
		</h1>
		<p>校大学生科协你同样感兴趣的部门，或是可以接受调剂的志愿，请勿与第一志愿相同</p>
	</label>
	<select
		id={second_label}
		name={second_label}
		class="input"
		disabled={$form.localApplicant?.modified}
		bind:value={second_choice}
		on:change={handleSecondChange}
	>
		<option disabled value="">--选择部门--</option>
		<option value="技术开发部">技术开发部</option>
		<option value="科普活动部">科普活动部</option>
		<option value="组织策划部">组织策划部</option>
		<option value="新闻宣传部">新闻宣传部</option>
		<option value="对外联络部">对外联络部</option>
		<option value="双创联合服务部">双创联合服务部</option>
	</select>
	<p class="err-msg">{second_error}</p>
</div>
