<script lang="ts">
	import '../app.css';

	import Footer from '$components/Footer/Footer.svelte';
	import Navbar from '$components/Navbar/Navbar.svelte';
</script>

<svelte:head>
	<title>欢迎加入校科协</title>
</svelte:head>

<Navbar />
<slot />
<Footer />
