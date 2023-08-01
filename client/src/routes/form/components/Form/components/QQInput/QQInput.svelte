<script lang="ts">
	import '../form.css';
	import '../form.css';
	import errorElements from '$stores/errorElements';

	let error = '';

	$: error = $errorElements.find((element) => element.id === 'qq')?.error || '';

	function handleChange(event: Event) {
		const value = (event.target as HTMLInputElement).value;
		if (value.length === 0) errorElements.upsert('qq', 'QQ号不能为空');
		else if (value.length < 5) errorElements.upsert('qq', '你的QQ号太短啦');
		else if (value.length > 11) errorElements.upsert('qq', '你的QQ号太长啦');
		else errorElements.remove('qq');
	}
</script>

<div class="form">
	<h1 class="title">
		<span>QQ</span>
		<span class="text-red-600">*</span>
	</h1>
	<label for="qq" class="label">我们联系你的备用方案</label>
	<input required id="qq" name="qq" type="number" class="input" placeholder="输入你的QQ号" on:change={handleChange} />
	<p class="err-msg" class:active={error}>{error}</p>
</div>
