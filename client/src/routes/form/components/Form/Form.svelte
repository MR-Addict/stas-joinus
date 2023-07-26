<script lang="ts">
	import z from 'zod';
	import { toasts } from '$stores/toasts';
	import { SubmitApplicant } from '$types/applicant';

	import Input from './components/Input/Input.svelte';
	import Select from './components/Select/Select.svelte';
	import Header from './components/Header/Header.svelte';
	import TextArea from './components/TextArea/TextArea.svelte';
	import SubmitButton from './components/SubmitButton/SubmitButton.svelte';

	let disabled = false;

	async function handleSubmit(event: SubmitEvent) {
		const data: any = {};
		new FormData(event.target as HTMLFormElement).forEach((value, key) => {
			data[key] = value;
		});
		const parsedData = SubmitApplicant.safeParse(data);
		if (!parsedData.success) {
			toasts.add('表单内容不合法', 'failed');
			return;
		}

		disabled = true;
		try {
			const res = await fetch('/api/applicant', {
				method: 'POST',
				body: JSON.stringify(parsedData.data),
				headers: { 'Content-Type': 'application/json' }
			}).then((res) => res.json());
			const result = z.object({ success: z.boolean(), message: z.string() }).parse(res);
			if (!result.success) toasts.add(result.message, 'failed');
			else toasts.add(result.message, 'success');
		} catch (err) {
			console.error(err);
			toasts.add('提交失败', 'failed');
		} finally {
			disabled = false;
		}
	}
</script>

<form id="form" on:submit|preventDefault={handleSubmit} class="w-full max-w-2xl flex flex-col items-end gap-5">
	<Header />
	<Input
		title="姓名"
		description="你的姓名"
		name="name"
		type="text"
		validate={(value) => {
			if (value.length === 0) return '姓名不能为空哦';
			else if (value.length < 2) return '你的名字太短啦';
			else if (value.length > 20) return '你的名字太长啦';
		}}
	/>
	<Select title="性别" description="你的性别" name="gender" options={['男', '女']} />
	<Input
		title="手机"
		description="我们会将面试信息以及面试结果通过短信的形式发送给你"
		name="phone"
		type="number"
		validate={(value) => {
			if (value.length === 0) return '手机号不能为空哦';
			else if (value.length !== 11) return '手机号应该是11位哦';
		}}
	/>
	<Input
		title="邮箱"
		description="我们会将面试信息以及面试结果通过邮件的形式发送给你"
		name="email"
		type="email"
		validate={(value) => {
			if (value.length === 0) return '邮箱不能为空哦';
			else if (value.length > 320) return '你的邮箱太长啦';
			else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(value)) return '请输入正确格式的邮箱哦';
		}}
	/>
	<Input
		title="QQ"
		description="我们联系你的备用方案"
		name="qq"
		type="number"
		validate={(value) => {
			if (value.length === 0) return 'QQ不能为空哦';
			else if (value.length < 5) return '你的QQ太短啦';
			else if (value.length > 11) return '你的QQ太长啦';
		}}
	/>
	<Input
		title="学号"
		description="你的学号"
		name="student_id"
		type="number"
		validate={(value) => {
			if (value.length === 0) return '学号不能为空哦';
			else if (value.length != 12) return '学号应该是11位哦';
		}}
	/>
	<Input
		title="学院"
		description="你所在的学院"
		name="college"
		type="text"
		validate={(value) => {
			if (value.length === 0) return '学院不能为空哦';
			else if (value.length > 50) return '你的学院名字太长啦';
		}}
	/>
	<Input
		title="专业"
		description="你现在的专业"
		name="major"
		type="text"
		validate={(value) => {
			if (value.length === 0) return '专业不能为空哦';
			else if (value.length > 50) return '你的专业名字太长啦';
		}}
	/>
	<Select
		title="第一志愿"
		description="校科协你最想加入的部门，我们会优先考虑你的第一志愿"
		name="first_choice"
		options={['技术开发部', '科普活动部', '组织策划部', '新闻宣传部', '对外联络部', '双创联合服务部']}
		validate={(value) => {
			// @ts-expect-error
			if (value === document.getElementById('second_choice')?.value) return '志愿重复，请修改你的志愿';
		}}
	/>
	<Select
		title="第二志愿"
		description="校科协你还想加入的部门，请勿与第一志愿相同"
		name="second_choice"
		options={['科普活动部', '技术开发部', '组织策划部', '新闻宣传部', '对外联络部', '双创联合服务部']}
		validate={(value) => {
			// @ts-expect-error
			if (value === document.getElementById('first_choice')?.value) return '志愿重复，请修改你的志愿';
		}}
	/>
	<TextArea
		title="自我介绍"
		description="写写你擅长的领域、个人爱好、取得的成就等等，让我们对你有更好的了解"
		name="introduction"
		validate={(value) => {
			if (value.length === 0) return '自我介绍不能为空哦';
			else if (value.length < 4) return '你的自我介绍太短啦';
			else if (value.length > 500) return '你的自我介绍太长啦';
		}}
	/>
	<SubmitButton {disabled} />
</form>
