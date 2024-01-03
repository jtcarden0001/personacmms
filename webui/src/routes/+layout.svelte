<script>
  import "../app.css";
  import Navbar from "./navbar.svelte";
</script>

<div class="h-screen flex flex-col">
<Navbar />
<slot />
</div>

<style></style>