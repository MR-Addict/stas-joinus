<script lang="ts">
	import '../form.css';
	import errorElements from '$stores/errorElements';

	let error = '';

	$: error = $errorElements.find((element) => element.id === 'major')?.error || '';

	function handleChange(event: Event) {
		const value = (event.target as HTMLInputElement).value;
		if (value.length === 0) errorElements.upsert('major', '专业不能为空');
		else if (value.length > 50) errorElements.upsert('major', '你的专业名字太长啦');
		else errorElements.remove('major');
	}
</script>

<div class="form">
	<h1 class="title">
		<span>专业</span>
		<span class="text-red-600">*</span>
	</h1>
	<label for="major" class="label">你现在的专业</label>
	<input
		required
		id="major"
		name="major"
		type="text"
		class="input"
		maxlength="50"
		placeholder="输入你的专业"
		on:change={handleChange}
	/>
	<p class="err-msg" class:active={error}>{error}</p>
</div>
