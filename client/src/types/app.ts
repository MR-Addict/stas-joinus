type ApiResultType<T = undefined> =
	| ({
			readonly success: true;
	  } & (T extends undefined ? {} : { readonly data: NonNullable<T> }))
	| {
			readonly success: false;
			readonly message: string;
	  };

export type { ApiResultType };
