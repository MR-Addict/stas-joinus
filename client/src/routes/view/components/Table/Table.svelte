<script lang="ts">
	import formatDate from '$lib/utils/formatDate';
	import type { TableFilter } from '$types/tableFilter';
	import type { ApplicantType } from '$types/applicant';
	import type { PaginationType } from '$types/pagination';

	export let tableFilter: TableFilter;
	export let applicants: ApplicantType[];
	export let pagination: PaginationType;
</script>

<div class="w-full overflow-x-auto">
	<table>
		<thead>
			<tr>
				<th>序号</th>
				{#if tableFilter.name}
					<th>姓名</th>
				{/if}
				{#if tableFilter.gender}
					<th>性别</th>
				{/if}
				{#if tableFilter.phone}
					<th>手机</th>
				{/if}
				{#if tableFilter.email}
					<th>邮箱</th>
				{/if}
				{#if tableFilter.qq}
					<th>QQ</th>
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
			{#each applicants as applicant, index (applicant.id)}
				<tr>
					<td>{index + 1 + (pagination.page - 1) * pagination.page_size}</td>
					{#if tableFilter.name}
						<td>{applicant.name}</td>
					{/if}
					{#if tableFilter.gender}
						<td>{applicant.gender === 'boy' ? '男' : '女'}</td>
					{/if}
					{#if tableFilter.phone}
						<td>{applicant.phone}</td>
					{/if}
					{#if tableFilter.email}
						<td>{applicant.email}</td>
					{/if}
					{#if tableFilter.qq}
						<td>{applicant.qq}</td>
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
						<td>{formatDate(applicant.submitted_at)}</td>
					{/if}
					{#if tableFilter.first_choice}
						<td>{applicant.first_choice}</td>
					{/if}
					{#if tableFilter.second_choice}
						<td>{applicant.second_choice}</td>
					{/if}
					{#if tableFilter.introduction}
						<td class="whitespace-pre-wrap">{applicant.introduction}</td>
					{/if}
				</tr>
			{/each}
		</tbody>
	</table>
</div>

<style lang="postcss">
	table {
		@apply w-full;
	}
	table th {
		@apply bg-gray-300 p-2 sm:p-3;
	}
	table tr {
		@apply text-left align-text-top border-b border-b-gray-300;
		&:nth-last-of-type(2n) {
			@apply bg-gray-100;
		}
	}
	table td {
		@apply text-gray-700 p-2 sm:p-3;
	}
	table tr :is(th, td:not(:last-of-type)) {
		@apply whitespace-nowrap;
	}
	table tr td:last-of-type {
		@apply min-w-72 max-w-2xl;
	}
</style>
