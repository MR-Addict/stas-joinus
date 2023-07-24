import type { ComponentType } from 'svelte';
import FaChartLine from 'svelte-icons/fa/FaChartLine.svelte';
import FaFlagCheckered from 'svelte-icons/fa/FaFlagCheckered.svelte';

interface LinkType {
	name: string;
	path: string;
	Icon: ComponentType;
}

export const links: LinkType[] = [
	{ name: '报名', path: '/', Icon: FaFlagCheckered },
	{ name: '统计', path: '/', Icon: FaChartLine }
];
