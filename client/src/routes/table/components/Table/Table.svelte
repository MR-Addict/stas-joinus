<script lang="ts">
	import applicants from '$stores/applicants';
	import pagination from '$stores/pagination';
	import formatDate from '$lib/utils/formatDate';
	import type { TableFilter } from '$types/tableFilter';

	export let tableFilter: TableFilter;
</script>

<div class="w-full overflow-x-auto">
	<table>
		<thead>
			<tr>
				<th>ID</th>
				{#if tableFilter.name}
					<th>姓名</th>
				{/if}
				{#if tableFilter.gender}
					<th>性别</th>
				{/if}
				{#if tableFilter.phone}
					<th>手机</th>
				{/if}
				{#if tableFilter.qq}
					<th>QQ</th>
				{/if}
				{#if tableFilter.email}
					<th>邮箱</th>
				{/if}
				{#if tableFilter.student_id}
					<th>学号</th>
				{/if}
				{#if tableFilter.college}
					<th>学院</th>
				{/if}
				{#if tableFilter.major}
					<th>专业</th>
				{/if}
				{#if tableFilter.created_at}
					<th>提交时间</th>
				{/if}
				{#if tableFilter.first_choice}
					<th>第一志愿</th>
				{/if}
				{#if tableFilter.second_choice}
					<th>第二志愿</th>
				{/if}
				{#if tableFilter.introduction}
					<th>自我介绍</th>
				{/if}
			</tr>
		</thead>
		<tbody>
			{#each $applicants as applicant, index (applicant.id)}
				<tr>
					<td>{index + 1 + ($pagination.page - 1) * $pagination.page_size}</td>
					{#if tableFilter.name}
						<td>{applicant.name}</td>
					{/if}
					{#if tableFilter.gender}
						<td>{applicant.gender}</td>
					{/if}
					{#if tableFilter.phone}
						<td>{applicant.phone}</td>
					{/if}
					{#if tableFilter.qq}
						<td>{applicant.qq}</td>
					{/if}
					{#if tableFilter.email}
						<td>{applicant.email}</td>
					{/if}
					{#if tableFilter.student_id}
						<td>{applicant.student_id}</td>
					{/if}
					{#if tableFilter.college}
						<td>{applicant.college}</td>
					{/if}
					{#if tableFilter.major}
						<td>{applicant.major}</td>
					{/if}
					{#if tableFilter.created_at}
						<td>{formatDate(applicant.create_at, true)}</td>
					{/if}
					{#if tableFilter.first_choice}
						<td>{applicant.first_choice}</td>
					{/if}
					{#if tableFilter.second_choice}
						<td>{applicant.second_choice}</td>
					{/if}
					{#if tableFilter.introduction}
						<td>{applicant.introduction}</td>
					{/if}
				</tr>
			{/each}
		</tbody>
	</table>
</div>

<style lang="postcss">
	table {
		@apply w-full border border-gray-300;
	}
	table thead th {
		@apply text-gray-800 bg-gray-200;
	}
	table tr {
		@apply text-left;
	}
	table :is(th, td) {
		@apply p-2;
	}
	table tr :is(th, td:not(:last-of-type)) {
		@apply whitespace-nowrap;
	}
	table tr td:last-of-type {
		width: 600px;
		min-width: 300px;
	}
	table tbody tr:nth-of-type(even) {
		@apply bg-gray-200;
	}
</style>
