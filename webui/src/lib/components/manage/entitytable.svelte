<script>
    import "../../../app.css";
    import EntityTableRowSuffix from "./entitytablerowsuffix.svelte"
    import AddNew from "./addnew.svelte"

    export let basePath = "";
    export let columnTitles;
    export let dataRows = [];
    export let itemData;

    console.log("basePath: " + basePath);
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
        <a href={basePath.concat("/", row.id, "/view")}>
          <div class="table-row hover:bg-slate-100">
            {#each row.data as cell}
              <div class="table-cell">{cell}</div>
            {/each}
            <EntityTableRowSuffix />
          </div>
        </a>
      {/each}
      <AddNew />
    </div>
  </div>
</div>

<style></style>