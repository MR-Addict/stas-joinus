export default function downloadSvg(svg: SVGElement, filename: string, scale = 3) {
	if (!svg.hasAttribute('xmlns')) svg.setAttribute('xmlns', 'http://www.w3.org/2000/svg');

	const svgData = new XMLSerializer().serializeToString(svg);
	const svgBlob = new Blob([svgData], { type: 'image/svg+xml;charset=utf-8' });
	const url = URL.createObjectURL(svgBlob);

	const img = new Image();
	img.src = url;
	img.onload = () => {
		const canvas = document.createElement('canvas');
		canvas.width = img.width * scale;
		canvas.height = img.height * scale;
		const ctx = canvas.getContext('2d');
		if (ctx) {
			ctx.setTransform(scale, 0, 0, scale, 0, 0);
			ctx.drawImage(img, 0, 0);
			const pngDataUrl = canvas.toDataURL('image/png');

			const a = document.createElement('a');
			a.download = filename.replace(/\.svg$/, '.png');
			a.href = pngDataUrl;
			a.click();
			URL.revokeObjectURL(url);
		}
	};
}
