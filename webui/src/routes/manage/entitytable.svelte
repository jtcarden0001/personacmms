<script>
    import "../../app.css";
    export let columnTitles;
    export let dataRows;
</script>

<div class="overflow-x-auto flex-auto">
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
          </tr>
        {/each}
      </tbody>
    </table>
  </div>

<style>

</style>