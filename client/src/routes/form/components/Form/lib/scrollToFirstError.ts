import { get } from 'svelte/store';

import errorElements from '$stores/errorElements';

export default function scrollToFirstError() {
	const offset = get(errorElements)
		.map((ele) => document.getElementById(ele.id))
		.map((ele) => ele?.offsetTop || 10000);
	document.documentElement.scrollTop = Math.min(...offset);
}
