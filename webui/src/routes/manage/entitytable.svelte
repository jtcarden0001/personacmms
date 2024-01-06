<script>
    import "../../app.css";
    import EntityTableRowSuffix from "./entitytablerowsuffix.svelte"
    import AddNew from "./addnew.svelte"

    export let columnTitles;
    export let dataRows = [];
</script>

<div class="overflow-x-auto flex-auto m-10">
    <table class="table">
      <thead>
        <tr>
          {#each columnTitles as column}
            <th>{column}</th>
          {/each}
        </tr>
      </thead>
      <tbody>
        {#each dataRows as row}
          <tr class="hover">
            {#each row as cell}
              <td>{cell}</td>
            {/each}
            <EntityTableRowSuffix />
          </tr>
        {/each}
        <AddNew />
      </tbody>
    </table>
  </div>

<style>
  tbody{
    cursor: pointer;
  }
</style>