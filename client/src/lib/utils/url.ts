import { PUBLIC_BACKEND_URL } from '$env/static/public';

export default function url(path: string) {
	if (PUBLIC_BACKEND_URL) return new URL(path, PUBLIC_BACKEND_URL).toString();
	return path;
}
