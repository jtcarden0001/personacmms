<script>
    let basePath = '/manage';
    let woPath = 'workorders';
    let eqPath = 'equipment';
    let taskPath = 'tasks';
    let toolPath = 'tools';
    let coPath = 'consumables';

    const pathSuffixes = [woPath, eqPath, taskPath, toolPath, coPath]

    const pathTitleMap = new Map([
        [woPath, "Work Orders"],
        [eqPath, "Equipment"],
        [taskPath, "Tasks"],
        [toolPath, "Tools"],
        [coPath, "Consumables"]
    ])

    const currentPath = window.location.pathname
    const selected = currentPath.split("/").pop();
</script>

<div class="h-4/5 flex items-center border-r-4 rounded-md">
  <ul class="menu bg-transparent rounded-box m-5">
    {#each pathSuffixes as pathsfx}
      {#if pathsfx === selected}
        <li class="text-white bg-purple-900 rounded-lg"><a href={basePath.concat("/", pathsfx)}>{pathTitleMap.get(pathsfx)}</a></li>
      {:else}
        <li><a href={basePath.concat("/", pathsfx)}>{pathTitleMap.get(pathsfx)}</a></li>
      {/if}
    {/each}
  </ul>
</div>

<style></style>