import z from 'zod';

const Response = z.object({
	success: z.boolean(),
	message: z.string()
});

type ResponseType = z.TypeOf<typeof Response>;

export { Response };
export type { ResponseType };
