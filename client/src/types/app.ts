import { z } from 'zod';

const exportFormats = ['xlsx', 'json'] as const;
const ExportFormat = z.enum(exportFormats);

type ApiResultType<T = undefined> =
	| ({
			readonly success: true;
			readonly message: string;
	  } & (T extends undefined ? {} : { readonly data: NonNullable<T> }))
	| {
			readonly success: false;
			readonly message: string;
	  };

const AppConfig = z.object({
	start_time: z.string().transform((value) => new Date(value)),
	end_time: z.string().transform((value) => new Date(value))
});

type ExportFormatType = z.infer<typeof ExportFormat>;
type AppConfigType = z.infer<typeof AppConfig>;

export { AppConfig, ExportFormat, exportFormats };
export type { ApiResultType, ExportFormatType, AppConfigType };
