<script lang="ts">
	import { z } from 'zod';
	import { Database, Table } from 'lucide-svelte';

	import toast from 'svelte-french-toast';

	import Modal from '$components/Modal/Modal.svelte';
	import Spinner from '$components/Spinner/Spinner.svelte';

	import { Applicant } from '$types/applicant';
	import { exportFormats, type ExportFormatType } from '$types/app';

	import mapGender from '$lib/utils/mapGender';
	import formatDate from '$lib/utils/formatDate';
	import persistantStore from '$lib/utils/persistantStore';
	import fetchApplicantsApi from '$lib/applicant/fetchApplicants';

	export let showModal = false;

	let pending = false;
	let exportFormat = persistantStore<ExportFormatType>('exportFormat', 'xlsx');

	async function handleExport() {
		pending = true;

		const res = await fetchApplicantsApi({ all: true });
		if (!res.success) {
			toast.error(res.message);
			pending = false;
			return;
		} else if (res.data.pagination.total === 0) {
			toast.error('没有可导出的数据');
			pending = false;
			return;
		}

		const applicants = z
			.array(Applicant.omit({ modified: true }))
			.parse(res.data.applicants)
			.map((a) => ({ ...a, gender: mapGender(a.gender), submitted_at: formatDate(new Date(a.submitted_at)) }));

		const fileName = `科协报名表单 ${formatDate(new Date()).replaceAll(':', '-')}.${$exportFormat}`;

		const { default: exportData } = await import('$lib/utils/exportData');

		if ($exportFormat === 'json') exportData.json(fileName, applicants);
		else if ($exportFormat === 'xlsx') {
			await exportData.xlsx(fileName, applicants, (sheet: any, key: string, index: number) => {
				const column = sheet.getColumn(index + 1);
				if (['id', 'name', 'gender'].includes(key)) column.width = 10;
				else if (key === 'introduction') column.width = 100;
				else column.width = 20;
				column.alignment = { wrapText: true, vertical: 'top', horizontal: 'left' };
			});
		}

		showModal = false;
		pending = false;
	}
</script>

<Modal bind:showModal disabled={pending}>
	<form on:submit|preventDefault={handleExport}>
		<h1>导出数据</h1>

		<p>请选择导出的数据格式</p>

		<div class="options-wrapper" style="grid-template-columns: repeat({exportFormats.length}, minmax(0, 1fr));">
			<div class="bubble" data-active-format-index={exportFormats.indexOf($exportFormat)}></div>
			{#each exportFormats as format (format)}
				<button
					type="button"
					class="option"
					class:active={$exportFormat === format}
					on:click={() => exportFormat.set(format)}
				>
					{#if format === 'xlsx'}
						<Table size={18} />
						<span>Excel</span>
					{:else if format === 'json'}
						<Database size={18} />
						<span>JSON</span>
					{/if}
				</button>
			{/each}
		</div>

		<button type="submit" disabled={pending}>
			<span>导出数据</span>
			{#if pending}
				<Spinner />
			{/if}
		</button>
	</form>
</Modal>

<style>
	form {
		@apply flex flex-col gap-2;
		@apply bg-white p-7 w-full max-w-xs rounded-lg text-center;
	}
	h1 {
		@apply font-semibold text-xl;
	}
	p {
		@apply text-gray-600 text-sm;
	}
	.options-wrapper {
		--py: 0.45rem;
		--px: 0.5rem;
		--gap: 0.5rem;
		padding: var(--py) var(--px);
		gap: var(--gap);
		@apply bg-black/10 isolate backdrop-blur-lg grid rounded-md;

		& .option {
			@apply flex items-center justify-center gap-1;
			@apply py-2.5 px-4 rounded-md;
		}

		& .bubble {
			top: var(--py);
			left: var(--px);
			height: calc(100% - var(--py) * 2);
			width: calc((100% - var(--px) * 2 - var(--gap)) / 2);
			@apply absolute bg-white -z-10 rounded-md duration-500 ease-in-out;

			&[data-active-format-index='1'] {
				transform: translateX(calc(100% * 1 + var(--gap)));
			}
		}
	}
	button[type='submit'] {
		@apply flex items-center justify-center gap-2 mt-2;
		@apply py-2.5 px-4 rounded-md duration-300 text-white bg-teal-600;

		&:enabled:hover {
			@apply bg-teal-700;
		}
		&:disabled {
			@apply bg-gray-300 text-gray-400 cursor-not-allowed;
		}
	}
</style>
