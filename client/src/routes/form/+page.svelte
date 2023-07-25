<script>
	import Header from './components/Header/Header.svelte';
	import Input from './components/Input/Input.svelte';
	import Select from './components/Select/Select.svelte';
	import TextArea from './components/TextArea/TextArea.svelte';

	const options = ['技术开发部', '组织策划部', '科普活动部', '新闻宣传部', '对外联络部', '双创联合服务部'];
</script>

<svelte:head>
	<title>欢迎加入校科协 • 参与报名</title>
</svelte:head>

<main>
	<form class="w-full max-w-2xl flex flex-col items-end gap-5">
		<Header />
		<Input title="姓名" description="您的姓名" name="name" type="text" minlength={2} maxlength={20} />
		<Select title="性别" description="您的性别" name="gender" options={['男', '女']} />
		<Input title="手机" description="我们会将面试信息以及面试结果通过短信的形式发送给您" name="phone" type="number" />
		<Input title="邮箱" description="我们会将面试信息以及面试结果通过邮件的形式发送给您" name="email" type="email" />
		<Input title="QQ" description="我们联系您的备用方案" name="qq" type="number" />
		<Input title="学号" description="您的学号" name="student_id" type="number" />
		<Input title="学院" description="您所在的学院" name="college" type="text" maxlength={50} />
		<Input title="专业" description="您现在的专业" name="major" type="text" maxlength={50} />
		<Select title="第一志愿" description="您最想加入的部门，我们会优先考虑您的第一志愿" name="first_choice" {options} />
		<Select title="第二志愿" description="您还想加入的部门，请不要与第一志愿相同" name="second_choice" {options} />
		<TextArea title="自我介绍" description="写写您擅长的领域、个人爱好、取得的成就等等" name="introduction" />
		<button type="submit">提交</button>
	</form>
</main>

<style lang="postcss">
	main {
		min-height: calc(100vh - 60px);
		min-height: calc(100svh - 60px);
		@apply bg-gray-100 flex flex-col items-center justify-center px-4 py-5 md:py-10;
	}
	button {
		@apply bg-teal-600 text-white py-2 px-6 rounded-md duration-300;
		&:hover {
			@apply bg-teal-700;
		}
	}
</style>
