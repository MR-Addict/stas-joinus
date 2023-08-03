export default function clickOutside(element: HTMLElement, callback: Function) {
	function handleClick(event: MouseEvent) {
		if (!element.contains(event.target as HTMLElement)) {
			callback();
		}
	}

	document.body.addEventListener('click', handleClick);

	return {
		update(newCallback: Function) {
			callback = newCallback;
		},
		destroy() {
			document.body.removeEventListener('click', handleClick);
		}
	};
}
