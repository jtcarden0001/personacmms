<script>
    import "../../app.css";
    import EntityTableRowSuffix from "./entitytablerowsuffix.svelte"

    export let columnTitles;
    export let dataRows;
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
          <tr>
            {#each row as cell}
              <td>{cell}</td>
            {/each}
            <EntityTableRowSuffix />
          </tr>
        {/each}
      </tbody>
    </table>
  </div>

<style>

</style>