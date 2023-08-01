<script lang="ts">
	import '../form.css';
	import errorElements from '$stores/errorElements';

	let error = '';

	$: error = $errorElements.find((element) => element.id === 'phone')?.error || '';

	function handleChange(event: Event) {
		const value = (event.target as HTMLInputElement).value;
		if (value.length === 0) errorElements.upsert('phone', '手机号不能为空');
		else if (value.length !== 11) errorElements.upsert('phone', '手机号应该是11位');
		else errorElements.remove('phone');
	}
</script>

<div class="form">
	<h1 class="title">
		<span>手机</span>
		<span class="text-red-600">*</span>
	</h1>
	<label for="phone" class="label">我们会将面试信息以及面试结果通过短信的形式发送给你</label>
	<input
		required
		id="phone"
		name="phone"
		type="number"
		class="input"
		placeholder="输入你的手机号"
		on:change={handleChange}
	/>
	<p class="err-msg" class:active={error}>{error}</p>
</div>
