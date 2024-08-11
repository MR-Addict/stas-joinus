import { TableFilter, type TableFilterType } from '$types/tableFilter';

const defautlFilter: TableFilterType = {
	name: true,
	gender: true,
	phone: false,
	qq: false,
	email: false,
	student_id: true,
	college: false,
	major: false,
	created_at: false,
	first_choice: true,
	second_choice: true,
	introduction: true
};

function loadFilter(): TableFilterType {
	const filter = localStorage.getItem('filter');
	if (!filter) return defautlFilter;

	const parsed = TableFilter.safeParse(JSON.parse(filter));
	if (parsed.success) return parsed.data;
	else return defautlFilter;
}

function saveFilter(filter: TableFilterType): void {
	localStorage.setItem('filter', JSON.stringify(filter));
}

const filter = {
	default: defautlFilter,
	load: loadFilter,
	save: saveFilter
};

export default filter;
