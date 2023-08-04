import z from 'zod';

const Options = z.enum(['技术开发部', '组织策划部', '科普活动部', '新闻宣传部', '对外联络部', '双创联合服务部']);

const SubmitApplicant = z.object({
	name: z.string().min(2).max(20),
	gender: z.enum(['男', '女']),
	phone: z.string().length(11),
	qq: z.string().min(5).max(11),
	email: z.string().email().max(320),
	student_id: z.string().length(12),
	college: z.string().max(50),
	major: z.string().max(50),
	first_choice: Options,
	second_choice: Options,
	introduction: z.string().min(4).max(500)
});

const Applicant = z.object({ id: z.string(), create_at: z.number() }).merge(SubmitApplicant);

type ApplicantType = z.TypeOf<typeof Applicant>;

export { Applicant, SubmitApplicant };
export type { ApplicantType };
