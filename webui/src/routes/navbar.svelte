<script>
  import "../app.css";

  function handleDetailsClick() {
    document.getElementById("AccountDetails").removeAttribute("open");
  }

  function getBaseManageLink(url) {
    const splitLink = url.split("/")
    if (splitLink.length < 3) {
      return "/manage/work-orders"
    }

    const targetlink = "/" + splitLink[1] + "/" + splitLink[2]
    return targetlink
  }

  //TODO: dynamically generate this link so that it redirects to parent manage page of current location
  // default to work-orders if navigating from a non-manage page
  // i.e. when at /manage/asset/1 it should redirect to /manage/asset
  let manageLink = "/manage/work-orders"

</script>

<!-- based on https://daisyui.com/components/navbar/ -->
<nav class="navbar bg-base-100 flex-none z-50">
    <div class="flex-1">
      <a class="btn btn-ghost text-xl" href="/">PersonaCMMS</a>
    </div>
    <div class="flex-none">
      <ul class="menu menu-horizontal px-1">
        <li><a href={manageLink}>Manage</a></li>
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
