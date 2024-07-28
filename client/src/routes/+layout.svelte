<script lang="ts">
	import { onMount } from 'svelte';
	import { Toaster } from 'svelte-french-toast';

	import '../app.css';
	import auth from '$stores/auth';

	import Footer from '$components/Footer/Footer.svelte';
	import Navbar from '$components/Navbar/Navbar.svelte';

	onMount(auth.ping);
</script>

<Navbar />
<slot />
<Footer />
<Toaster />
