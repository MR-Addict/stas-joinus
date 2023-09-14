<script lang="ts">
	import '../form.css';
	import { onMount } from 'svelte';
	import inputs from '$stores/inputs';

	let error = '';
	let first_choice = '技术开发部';
	let second_choice = '科普活动部';

	$: error = $inputs.find((input) => input.id === 'first_choice')?.error || '';

	onMount(() => inputs.register('first_choice', validate));

	function handleChange() {
		inputs.update('first_choice', first_choice + ',' + second_choice);
	}

	function validate(value: string) {
		const [first, second] = value.split(',');
		if (first && second && first === second) return '志愿重复，请修改第一或第二志愿';
		else return '';
	}
</script>

<div class="form" class:error>
	<h1 class="title">
		<span>第一志愿</span>
		<span class="text-red-600">*</span>
	</h1>
	<label for="first_choice" class="label">校大学生科协你最想加入的部门，我们会优先考虑你的第一志愿</label>
	<select id="first_choice" name="first_choice" class="input" bind:value={first_choice} on:change={handleChange}>
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
		<span>第二志愿</span>
		<span class="text-red-600">*</span>
	</h1>
	<label for="second_choice" class="label">校大学生科协你还想加入的部门，请勿与第一志愿相同</label>
	<select id="second_choice" name="second_choice" class="input" bind:value={second_choice} on:change={handleChange}>
		<option value="科普活动部">科普活动部</option>
		<option value="技术开发部">技术开发部</option>
		<option value="组织策划部">组织策划部</option>
		<option value="新闻宣传部">新闻宣传部</option>
		<option value="对外联络部">对外联络部</option>
		<option value="双创联合服务部">双创联合服务部</option>
	</select>
	<p class="err-msg">{error}</p>
</div>
