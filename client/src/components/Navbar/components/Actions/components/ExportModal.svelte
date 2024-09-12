<script lang="ts">
	import { z } from 'zod';
	import toast from 'svelte-french-toast';
	import Modal from '$components/Modal/Modal.svelte';

	import Spinner from '$components/Spinner/Spinner.svelte';
	import fetchApplicantsApi from '$lib/applicant/fetchApplicants';
	import persistantStore from '$lib/utils/persistantStore';
	import formatDate from '$lib/utils/formatDate';
	import exportData from '$lib/utils/exportData';
	import mapGender from '$lib/utils/mapGender';
	import { Applicant } from '$types/applicant';
	import { exportFormats, type ExportFormatType } from '$types/app';

	export let showModal = false;

	let pending = false;
	let exportFormat = persistantStore<ExportFormatType>('exportFormat', 'xlsx');

	async function handleExport() {
		pending = true;

		const res = await fetchApplicantsApi({ all: true });
		if (!res.success) {
			toast.error(res.message);
			return;
		}

		const applicants = z
			.array(Applicant.omit({ modified: true }))
			.parse(res.data.applicants)
			.map((a) => ({ ...a, gender: mapGender(a.gender), submitted_at: formatDate(new Date(a.submitted_at)) }));

		const fileName = `科协报名表单-${formatDate(new Date()).replaceAll(':', '-')}.${$exportFormat}`;
		if ($exportFormat === 'json') exportData.json(fileName, applicants);
		else if ($exportFormat === 'csv') exportData.csv(fileName, applicants);
		else if ($exportFormat === 'xlsx') {
			await exportData.xlsx(fileName, applicants, (sheet, key, index) => {
				const column = sheet.getColumn(index + 1);
				if (['id', 'name', 'gender'].includes(key)) column.width = 10;
				else if (key === 'introduction') column.width = 100;
				else column.width = 20;
				column.alignment = { wrapText: true, vertical: 'top', horizontal: 'left' };
			});
		}

		pending = false;
	}
</script>

<Modal bind:showModal disabled={pending}>
	<form on:submit|preventDefault={handleExport}>
		<h1>导出数据</h1>

		<p>请选择需要导出的数据格式</p>

		<div class="options-wrapper" style="grid-template-columns: repeat({exportFormats.length}, minmax(0, 1fr));">
			<div class="bubble" data-active-format-index={exportFormats.indexOf($exportFormat)} />
			{#each exportFormats as format (format)}
				<button
					type="button"
					class="option"
					class:active={$exportFormat === format}
					on:click={() => exportFormat.set(format)}
				>
					{'.' + format}
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
		@apply flex flex-col items-center gap-2;
		@apply bg-white p-7 w-full max-w-xs rounded-lg text-center;
	}
	h1 {
		@apply font-semibold text-xl;
	}
	p {
		@apply text-gray-600;
	}
	.options-wrapper {
		--py: 0.45rem;
		--px: 0.5rem;
		--gap: 0.25rem;
		padding: var(--py) var(--px);
		gap: var(--gap);
		@apply bg-black/10 isolate backdrop-blur-lg grid rounded-md;

		& .option {
			@apply py-1 px-2 rounded-md;
		}

		& .bubble {
			top: var(--py);
			left: var(--px);
			height: calc(100% - var(--py) * 2);
			width: calc((100% - var(--px) * 2 - var(--gap) * 2) / 3);
			@apply absolute bg-black/20 -z-10 rounded-md duration-500 ease-in-out;

			&[data-active-format-index='1'] {
				transform: translateX(calc(100% * 1 + var(--gap)));
			}

			&[data-active-format-index='2'] {
				transform: translateX(calc(100% * 2 + var(--gap) * 2));
			}
		}
	}
	button[type='submit'] {
		@apply flex items-center justify-center gap-2 mt-2;
		@apply py-2 px-4 rounded-md duration-300 text-white bg-teal-600;

		&:enabled:hover {
			@apply bg-teal-700;
		}
		&:disabled {
			@apply bg-gray-300 text-gray-400 cursor-not-allowed;
		}
	}
</style>
