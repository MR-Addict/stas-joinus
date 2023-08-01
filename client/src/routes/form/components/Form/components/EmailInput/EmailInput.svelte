<script lang="ts">
	import '../form.css';
	import errorElements from '$stores/errorElements';

	let error = '';

	$: error = $errorElements.find((element) => element.id === 'email')?.error || '';

	function handleChange(event: Event) {
		const value = (event.target as HTMLInputElement).value;
		if (value.length === 0) errorElements.upsert('email', '邮箱不能为空');
		else if (value.length > 320) errorElements.upsert('email', '你的邮箱太长啦');
		else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(value)) errorElements.upsert('email', '请输入正确格式的邮箱');
		else errorElements.remove('email');
	}
</script>

<div class="form">
	<h1 class="title">
		<span>邮箱</span>
		<span class="text-red-600">*</span>
	</h1>
	<label for="email" class="label">我们会将面试信息以及面试结果通过邮件的形式发送给你</label>
	<input
		required
		id="email"
		name="email"
		type="email"
		class="input"
		maxlength="320"
		placeholder="输入你的邮箱"
		on:change={handleChange}
	/>
	<p class="err-msg" class:active={error}>{error}</p>
</div>
