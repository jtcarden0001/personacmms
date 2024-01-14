<script>
    import "../../../app.css";
    import EntityTableRowSuffix from "./entitytablerowsuffix.svelte"
    import AddNew from "./addnew.svelte"

    export let columnTitles = [];
    export let dataRows = [];
</script>

<div class="overflow-x-auto flex-auto m-10">
  <div class="table">
    <div class="table-header-group bg-slate-200">
      <div class="table-row">
        {#each columnTitles as column}
          <div class="table-cell">{column}</div>
        {/each}
        <div class="table-cell">Actions</div>
      </div>
    </div>
    <div class="table-row-group">
      {#each dataRows as row}
          <div class="table-row hover:bg-slate-100">
            {#each row.cells as cell}
              <a class="table-cell" href={row.itemRoute}>
                {cell}
              </a>
            {/each}
            <EntityTableRowSuffix />
          </div>
      {/each}
      <AddNew />
    </div>
  </div>
</div>

<style></style>