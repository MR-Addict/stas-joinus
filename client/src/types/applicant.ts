import z from 'zod';

import { firstChoice, secondChoice } from '$data/choice';

const Options = z.enum(['技术开发部', '组织策划部', '科普活动部', '新闻宣传部', '对外联络部', '双创联合服务部']);

const ApplicantChoiceStats = z.object({ boy: z.number(), girl: z.number() });

const ApplicantStats = z.object({
	name: Options,
	first_choice: ApplicantChoiceStats,
	second_choice: ApplicantChoiceStats
});

type ChoiceType = typeof firstChoice | typeof secondChoice;

const ApplicantSubmit = z.object({
	name: z.string().min(2).max(20),
	gender: z.enum(['boy', 'girl']),
	phone: z.string().length(11),
	email: z.string().email().max(320),
	qq: z.string().min(5).max(11),
	student_id: z.string().length(12),
	college: z.string().max(50),
	major: z.string().max(50),
	first_choice: Options,
	second_choice: Options,
	introduction: z.string().min(4).max(500)
});

const Applicant = z.object({ id: z.number(), modified: z.boolean(), submitted_at: z.string() }).merge(ApplicantSubmit);

type ApplicantChoiceStatsType = z.TypeOf<typeof ApplicantChoiceStats>;
type ApplicantStatsType = z.TypeOf<typeof ApplicantStats>;
type ApplicantSubmitType = z.TypeOf<typeof ApplicantSubmit>;
type ApplicantType = z.TypeOf<typeof Applicant>;

export { Applicant, ApplicantSubmit, ApplicantStats, ApplicantChoiceStats };
export type { ApplicantType, ApplicantSubmitType, ApplicantStatsType, ApplicantChoiceStatsType, ChoiceType };
