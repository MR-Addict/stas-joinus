<script lang="ts">
	import '../form.css';
	import errorElements from '$stores/errorElements';

	let error = '';

	$: error = $errorElements.find((element) => element.id === 'college')?.error || '';

	function handleChange(event: Event) {
		const value = (event.target as HTMLInputElement).value;
		if (value.length === 0) errorElements.upsert('college', '学院不能为空');
		else if (value.length > 50) errorElements.upsert('college', '你的学院名字太长啦');
		else errorElements.remove('college');
	}
</script>

<div class="form">
	<h1 class="title">
		<span>学院</span>
		<span class="text-red-600">*</span>
	</h1>
	<label for="college" class="label">你所在的学院</label>
	<input
		required
		id="college"
		name="college"
		type="text"
		class="input"
		maxlength="50"
		placeholder="输入的学院"
		on:change={handleChange}
	/>
	<p class="err-msg" class:active={error}>{error}</p>
</div>
