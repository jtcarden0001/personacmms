<script>
    import Manage from "$lib/components/manage/manage.svelte";

    export let data = {"content":[]};

    let basePath = "/manage/work-orders";
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
            "data": row
        });
    }

    const processedViewData = new Map();
    for (let i = 0; i < data.content.length; i++) {
        processedViewData.set(data.content[i].id, [
            ["Title", data.content[i].taskTitle],
            ["Equipment", data.content[i].equipmentTitle],
            ["Status", data.content[i].statusTitle],
            ["Created Date", data.content[i].createdDate],
            ["Completed Date", data.content[i].completedDate]
        ]);
    }

        
</script>

<Manage basePath={basePath} columns={columns} dataRows={processedRowData} itemData={processedViewData} />

<style></style>