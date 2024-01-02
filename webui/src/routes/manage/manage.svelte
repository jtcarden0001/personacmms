<script>
    import EntityTable from "./entitytable.svelte";
    import SideNav from "./sidenav.svelte";

    let columns = ["Name", "Job", "Favorite Color"];
    let data = [
        ["Cy Ganderton", "Quality Control Specialist", "Blue"],
        ["Hart Hagerty", "Desktop Support Technician", "Purple"],
        ["Brice Swyre", "Tax Accountant", "Red"]
    ]
</script>

<div class="flex">
    <SideNav />
    <EntityTable columnTitles={columns} dataRows={data} />
</div>

<style></style>