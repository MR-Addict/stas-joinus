import { get } from 'svelte/store';

import inputs from '$stores/inputs';

export default function scrollToFirstError() {
	const offset = get(inputs)
		.filter((input) => input.error !== '')
		.map((input) => document.getElementById(input.id))
		.map((input) => input?.offsetTop || 999999);
	document.documentElement.scrollTop = Math.min(...offset);
}
