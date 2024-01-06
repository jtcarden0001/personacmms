<script>
    import EntityTable from "./entitytable.svelte";
    import SideNav from "./sidenav.svelte";
    import AddNew from "./addnew.svelte";

    export let columns = ["Title", "Equipment", "Status"];
    export let dataRows = [];
</script>

<div class="flex grow z-20">
  <SideNav />
  <div class="flex-auto">
    <EntityTable columnTitles={columns} dataRows={dataRows} />
  </div>
</div>

<style></style>