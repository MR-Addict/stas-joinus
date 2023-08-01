<script lang="ts">
	import '../form.css';
	import errorElements from '$stores/errorElements';

	let error = '';

	$: error = $errorElements.find((element) => element.id === 'student_id')?.error || '';

	function handleChange(event: Event) {
		const value = (event.target as HTMLInputElement).value;
		if (value.length === 0) errorElements.upsert('student_id', '学号不能为空');
		else if (value.length != 12) errorElements.upsert('student_id', '学号应该是11位');
		else errorElements.remove('student_id');
	}
</script>

<div class="form">
	<h1 class="title">
		<span>学号</span>
		<span class="text-red-600">*</span>
	</h1>
	<label for="student_id" class="label">你的学号</label>
	<input
		required
		id="student_id"
		name="student_id"
		type="number"
		class="input"
		placeholder="输入你的学号"
		on:change={handleChange}
	/>
	<p class="err-msg" class:active={error}>{error}</p>
</div>
