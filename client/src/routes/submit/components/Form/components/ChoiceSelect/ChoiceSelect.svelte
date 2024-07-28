<script lang="ts">
	import { onMount } from 'svelte';
	import { FlaskConical } from 'lucide-svelte';

	import '../form.css';
	import inputs from '$stores/inputs';

	const first_label = 'first_choice';
	const second_label = 'second_choice';

	let error = '';
	let ref: HTMLDivElement;
	let first_choice = '';
	let second_choice = '';

	$: error = $inputs.find((input) => input.id === first_label)?.error || '';

	onMount(() => inputs.register(first_label, ref, validate));

	function handleChange() {
		inputs.update(first_label, first_choice + ',' + second_choice);
	}

	function validate(value: string) {
		const [first, second] = value.split(',');
		if (!first) return '请选择第一志愿';
		else if (!second) return '请选择第二志愿';
		else if (first === second) return '志愿重复，请修改第一或第二志愿';
		else return '';
	}
</script>

<div bind:this={ref} class="form" class:error>
	<h1 class="title">
		<FlaskConical size={18} />
		<span>第一志愿</span>
	</h1>
	<label for={first_label} class="label">校大学生科协你最想加入的部门，我们会优先考虑你的第一志愿</label>
	<select id={first_label} name={first_label} class="input" bind:value={first_choice} on:change={handleChange}>
		<option disabled value="">--选择部门--</option>
		<option value="技术开发部">技术开发部</option>
		<option value="科普活动部">科普活动部</option>
		<option value="组织策划部">组织策划部</option>
		<option value="新闻宣传部">新闻宣传部</option>
		<option value="对外联络部">对外联络部</option>
		<option value="双创联合服务部">双创联合服务部</option>
	</select>
	<p class="err-msg">{error}</p>
</div>

<div class="form" class:error>
	<h1 class="title">
		<FlaskConical size={18} />
		<span>第二志愿</span>
	</h1>
	<label for={second_label} class="label">
		校大学生科协你同样感兴趣的部门，或是可以接受调剂的志愿，请勿与第一志愿相同
	</label>
	<select id={second_label} name={second_label} class="input" bind:value={second_choice} on:change={handleChange}>
		<option disabled value="">--选择部门--</option>
		<option value="技术开发部">技术开发部</option>
		<option value="科普活动部">科普活动部</option>
		<option value="组织策划部">组织策划部</option>
		<option value="新闻宣传部">新闻宣传部</option>
		<option value="对外联络部">对外联络部</option>
		<option value="双创联合服务部">双创联合服务部</option>
	</select>
	<p class="err-msg">{error}</p>
</div>
