<script lang="ts">
	import { onMount } from 'svelte';

	import '../app.css';
	import auth from '$stores/auth';

	import Footer from '$components/Footer/Footer.svelte';
	import Navbar from '$components/Navbar/Navbar.svelte';
	import Toasts from '$components/Toasts/Toasts.svelte';

	onMount(auth.ping);
</script>

<Toasts />
<Navbar />
<slot />
<Footer />
