export default function url(path: string) {
	return new URL(path, 'http://localhost:4000').toString();
}
