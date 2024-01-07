<script>
    import EntityTable from "./entitytable.svelte";
    import SideNav from "./sidenav.svelte";

    export let basePath
    export let columns = ["Title", "Equipment", "Status"];
    export let dataRows = [{
      data: []
    }];
    export const itemData = new Map();
</script>

<div class="flex grow z-20">
  <SideNav />
  <div class="flex-auto">
    <EntityTable basePath={basePath} columnTitles={columns} dataRows={dataRows} itemData={itemData} />
  </div>
</div>

<style></style>