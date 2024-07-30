<script lang="ts">
	import { onMount } from 'svelte';
	import { Toaster } from 'svelte-french-toast';

	import '../app.css';
	import auth from '$stores/auth';

	import Footer from '$components/Footer/Footer.svelte';
	import Navbar from '$components/Navbar/Navbar.svelte';
	import config from '$stores/config';

	onMount(() => {
		auth.ping();
		config.ping();
	});
</script>

<div>
	<Navbar />
	<slot />
</div>
<Footer />
<Toaster />

<style>
	div {
		min-height: 100vh;
		min-height: 100dvh;
		@apply flex flex-col;
	}
</style>
