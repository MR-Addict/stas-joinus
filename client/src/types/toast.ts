import z from 'zod';

const ToastStatus = z.enum(['pending', 'success', 'failed']);

const Toast = z.object({
	id: z.string(),
	message: z.string(),
	status: ToastStatus
});

type ToastType = z.TypeOf<typeof Toast>;
type ToastStatusType = z.TypeOf<typeof ToastStatus>;

export { Toast, ToastStatus };
export type { ToastType, ToastStatusType };
