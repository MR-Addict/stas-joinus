import ExcelJS from 'exceljs';

function downloadFile(name: string, url: string) {
	const link = document.createElement('a');
	link.setAttribute('href', url);
	link.setAttribute('download', name);
	link.click();
}

function exportJSON<T>(filename: string, data: T) {
	const header = 'data:text/json;charset=utf-8,\uFEFF';
	const content = JSON.stringify(data, null, 2);
	const encodedUri = header + encodeURIComponent(content);

	downloadFile(filename, encodedUri);
}

function exportCSV<T>(filename: string, data: T[]) {
	const header = 'data:text/csv;charset=utf-8,\uFEFF';
	const content =
		Object.keys(data[0] as any).join(',') + '\n' + data.map((row: any) => Object.values(row).join(',')).join('\n');
	const encodedUri = header + encodeURIComponent(content);

	downloadFile(filename, encodedUri);
}

async function exportXlsx<T>(
	filename: string,
	data: T[],
	setStyle: (column: ExcelJS.Worksheet, key: string, index: number) => void
) {
	// create workbook and sheet
	const workbook = new ExcelJS.Workbook();
	workbook.modified = new Date();
	const sheet = workbook.addWorksheet('报名表单');

	// extract keys and values
	const keys = Object.keys(data[0] as any);
	const values = data.map((row: any) => Object.values(row));

	// set header
	sheet.addRow(keys);
	values.forEach((row) => sheet.addRow(row));

	// set column style
	keys.forEach((key, index) => setStyle(sheet, key, index));

	// download file
	await workbook.xlsx.writeBuffer().then((buffer) => {
		const blob = new Blob([buffer], { type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet' });
		const url = URL.createObjectURL(blob);

		downloadFile(filename, url);
	});
}

const exportData = {
	json: exportJSON,
	csv: exportCSV,
	xlsx: exportXlsx
};

export default exportData;
