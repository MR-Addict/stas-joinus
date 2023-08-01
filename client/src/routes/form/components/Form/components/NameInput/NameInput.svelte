<script lang="ts">
	import '../form.css';
	import errorElements from '$stores/errorElements';

	let error = '';

	$: error = $errorElements.find((element) => element.id === 'name')?.error || '';

	function handleChange(event: Event) {
		const value = (event.target as HTMLInputElement).value;
		if (value.length === 0) errorElements.upsert('name', '姓名不能为空');
		else if (value.length < 2) errorElements.upsert('name', '你的名字太短啦');
		else if (value.length > 20) errorElements.upsert('name', '你的名字太长啦');
		else errorElements.remove('name');
	}
</script>

<div class="form">
	<h1 class="title">
		<span>姓名</span>
		<span class="text-red-600">*</span>
	</h1>
	<label for="name" class="label">你的姓名</label>
	<input
		required
		id="name"
		name="name"
		type="text"
		class="input"
		minlength="2"
		maxlength="20"
		placeholder="输入你的姓名"
		on:change={handleChange}
	/>
	<p class="err-msg" class:active={error}>{error}</p>
</div>
