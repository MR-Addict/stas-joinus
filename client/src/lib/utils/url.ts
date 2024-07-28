export default function url(path: string) {
	const backendUrl = import.meta.env.VITE_BACKEND_URL;
	if (backendUrl) return new URL(path, backendUrl).toString();
	return path;
}
