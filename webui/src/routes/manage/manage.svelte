<script>
    import EntityTable from "./entitytable.svelte";
    import SideNav from "./sidenav.svelte";
    import AddNew from "./addnew.svelte";

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
  <div class="flex-auto">
    <EntityTable columnTitles={columns} dataRows={data} />
  </div>
</div>

<style></style>