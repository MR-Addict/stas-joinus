import z from 'zod';

const Pagination = z.object({
	page: z.number(),
	total: z.number(),
	page_size: z.number(),
	query: z.string().optional()
});

type PaginationType = z.TypeOf<typeof Pagination>;

export { Pagination };
export type { PaginationType };
