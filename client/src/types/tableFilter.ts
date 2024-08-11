import { z } from 'zod';

const TableFilter = z.object({
	name: z.boolean(),
	gender: z.boolean(),
	phone: z.boolean(),
	qq: z.boolean(),
	email: z.boolean(),
	student_id: z.boolean(),
	college: z.boolean(),
	major: z.boolean(),
	created_at: z.boolean(),
	first_choice: z.boolean(),
	second_choice: z.boolean(),
	introduction: z.boolean()
});

type TableFilterType = z.infer<typeof TableFilter>;

export type { TableFilterType };
export { TableFilter };
