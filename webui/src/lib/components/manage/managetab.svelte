<script>
    import EntityTable from "./entitytable.svelte";
    import SideNav from "./sidenav.svelte";

    export let columns = ["Title", "Asset", "Status"];
    export let dataRows = [];
</script>

<div class="flex grow z-20">
  <SideNav />
  <div class="flex-auto">
    <EntityTable columnTitles={columns} dataRows={dataRows} />
  </div>
</div>

<style></style>