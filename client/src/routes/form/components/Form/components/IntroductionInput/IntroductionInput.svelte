<script lang="ts">
	import '../form.css';
	import errorElements from '$stores/errorElements';

	let error = '';

	$: error = $errorElements.find((element) => element.id === 'introduction')?.error || '';

	function handleChange(event: Event) {
		const value = (event.target as HTMLInputElement).value;
		if (value.length === 0) errorElements.upsert('introduction', '自我介绍不能为空');
		else if (value.length < 4) errorElements.upsert('introduction', '你的自我介绍太短啦');
		else if (value.length > 500) errorElements.upsert('introduction', '你的自我介绍太长啦');
		else errorElements.remove('introduction');
	}
</script>

<div class="form">
	<h1 class="title">
		<span>自我介绍</span>
		<span class="text-red-600">*</span>
	</h1>
	<label for="introduction" class="label">写写你擅长的领域、个人爱好、取得的成就等等，让我们对你有更好的了解</label>
	<textarea
		required
		id="introduction"
		name="introduction"
		class="textarea"
		minlength="4"
		maxlength="500"
		placeholder="输入你的姓名"
		on:change={handleChange}
	/>
	<p class="err-msg" class:active={error}>{error}</p>
</div>
