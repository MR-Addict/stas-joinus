<script lang="ts">
	import '../app.css';

	import Footer from '$components/Footer/Footer.svelte';
	import Navbar from '$components/Navbar/Navbar.svelte';
	import Toasts from '$components/Toasts/Toasts.svelte';
</script>

<Toasts />
<Navbar />
<slot />
<Footer />
