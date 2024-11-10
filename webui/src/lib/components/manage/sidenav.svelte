<script>
    let basePath = '/manage';
    let woPath = 'work-orders';
    let eqPath = 'asset';
    let taskPath = 'tasks';
    let toolPath = 'tools';
    let coPath = 'consumables';

    const pathSuffixes = [woPath, eqPath, taskPath, toolPath, coPath]

    const pathTitleMap = new Map([
        [woPath, "Work Orders"],
        [eqPath, "Asset"],
        [taskPath, "Tasks"],
        [toolPath, "Tools"],
        [coPath, "Consumables"]
    ])

    const currentPath = window.location.pathname
    const selected = currentPath.split("/").pop();
</script>

<!-- currently the side nav hides its content behind the top nave bar once the window height gets too small.  
     TODO: fix this scenario where the sidnavholds its position and is scrollable once the window gets too vertically small -->
<div class="sidenav flex items-center border-r-4 rounded-md">
  <ul class="menu bg-transparent rounded-box m-5">
    {#each pathSuffixes as pathsfx}
      {#if pathsfx === selected}
        <li class="text-white bg-indigo-800 rounded-lg"><a href={basePath.concat("/", pathsfx)}>{pathTitleMap.get(pathsfx)}</a></li>
      {:else}
        <li><a href={basePath.concat("/", pathsfx)}>{pathTitleMap.get(pathsfx)}</a></li>
      {/if}
    {/each}
  </ul>
</div>

<style>
  .sidenav {
    height: 75vh;
    position: -webkit-sticky;
    position: sticky;
    top: 12.5%;
  }
</style>