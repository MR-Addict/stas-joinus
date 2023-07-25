<script lang="ts">
	import '../app.css';

	import Footer from '$components/Footer/Footer.svelte';
	import Navbar from '$components/Navbar/Navbar.svelte';
</script>

<Navbar />
<slot />
<Footer />
