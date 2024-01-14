<script>
    import ManageTab from "$lib/components/manage/managetab.svelte";

    export let data = {"content":[]};

    let columns = ["Title", "Equipment", "Status"];
    const processedRowData = [];
    for (let i = 0; i < data.content.length; i++) {
        let id = data.content[i].id;
        // ensure this order matches the order of the columns
        let row = [
            data.content[i].taskTitle,
            data.content[i].equipmentTitle,
            data.content[i].statusTitle
        ];
        
        processedRowData.push({
            "id": id,
            "itemRoute": "/manage/equipment/" + data.content[i].equipmentId + "/tasks/" + data.content[i].taskId + "/work-orders/" + id + "/view",
            "entityTableRow": row
        });
    }

</script>

<ManageTab columns={columns} dataRows={processedRowData} />
<style></style>