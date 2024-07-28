import { z } from 'zod';

type ApiResultType<T = undefined> =
	| ({
			readonly success: true;
	  } & (T extends undefined ? {} : { readonly data: NonNullable<T> }))
	| {
			readonly success: false;
			readonly message: string;
	  };

const AppConfig = z.object({
	start_time: z.string().transform((value) => new Date(value)),
	end_time: z.string().transform((value) => new Date(value))
});

type AppConfigType = z.infer<typeof AppConfig>;

export { AppConfig };
export type { ApiResultType, AppConfigType };
