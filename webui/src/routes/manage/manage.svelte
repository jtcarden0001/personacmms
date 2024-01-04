<script>
    import EntityTable from "./entitytable.svelte";
    import SideNav from "./sidenav.svelte";

    let columns = ["Name", "Job", "Favorite Color"];
    let data = [
        ["Cy Ganderton", "Quality Control Specialist", "Blue"],
        ["Hart Hagerty", "Desktop Support Technician", "Purple"],
        ["Brice Swyre", "Tax Accountant", "Red"],
        ["Dorolice Dugdale", "Senior Cost Accountant", "Red"],
        ["Dorolice Dugdale", "Senior Cost Accountant", "Red"],
        ["Dorolice Dugdale", "Senior Cost Accountant", "Red"],
        ["Dorolice Dugdale", "Senior Cost Accountant", "Red"],
        ["Dorolice Dugdale", "Senior Cost Accountant", "Red"],
        ["Dorolice Dugdale", "Senior Cost Accountant", "Red"],
        ["Dorolice Dugdale", "Senior Cost Accountant", "Red"],
        ["Dorolice Dugdale", "Senior Cost Accountant", "Red"],
        ["Dorolice Dugdale", "Senior Cost Accountant", "Red"],
        ["Hart Hagerty", "Desktop Support Technician", "Purple"],
        ["Hart Hagerty", "Desktop Support Technician", "Purple"],
        ["Hart Hagerty", "Desktop Support Technician", "Purple"],
        ["Hart Hagerty", "Desktop Support Technician", "Purple"],
        ["Hart Hagerty", "Desktop Support Technician", "Purple"],
        ["Hart Hagerty", "Desktop Support Technician", "Purple"],
        ["Hart Hagerty", "Desktop Support Technician", "Purple"],
    ]
</script>

<div class="flex grow z-20">
    <SideNav />
    <EntityTable columnTitles={columns} dataRows={data} />
</div>

<style></style>