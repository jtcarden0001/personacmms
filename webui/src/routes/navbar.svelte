<script>
    import "../app.css";

    function handleDetailsClick() {
        document.getElementById("AccountDetails").removeAttribute("open");
    }

</script>

<!-- based on https://daisyui.com/components/navbar/ -->
<nav class="navbar bg-base-100 flex-none z-50">
    <div class="flex-1">
      <a class="btn btn-ghost text-xl" href="/">PersonaCMMS</a>
    </div>
    <div class="flex-none">
      <ul class="menu menu-horizontal px-1">
        <li><a href="/manage">Manage</a></li>
        <li><a href="/about">About</a></li>
        <li>
          <details id="AccountDetails">
            <summary>
              Account
            </summary>
            <ul class="p-2 bg-base-100 rounded-t-none">
              <li><a href="/signup" on:click={handleDetailsClick}>Sign Up</a></li>
              <li><a href="/login" on:click={handleDetailsClick}>Log In</a></li>
            </ul>
          </details>
        </li>
      </ul>
    </div>
</nav>

<style>

</style>
