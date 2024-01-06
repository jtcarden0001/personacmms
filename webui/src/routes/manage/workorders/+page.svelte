<script>
    import Manage from "../manage.svelte";

    export let data = {"content":[]};

    let columns = ["Title", "Equipment", "Status"];
    let processedData = [];
    for (let i = 0; i < data.content.length; i++) {
        let row = [];
        row.push(data.content[i].taskTitle);
        row.push(data.content[i].equipmentTitle);
        row.push(data.content[i].statusTitle);
        processedData.push(row);
    }

    console.log(processedData);
</script>

<Manage columns={columns} dataRows={processedData} />

<style></style>