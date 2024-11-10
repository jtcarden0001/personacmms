<script>
    import ManageTab from "$lib/components/manage/managetab.svelte";

    export let data = {"content":[]};

    let columns = ["Title", "Asset", "Status"];
    const processedRowData = [];
    for (let i = 0; i < data.content.length; i++) {
        let id = data.content[i].id;
        // ensure this order matches the order of the columns
        let row = [
            data.content[i].taskTitle,
            data.content[i].assetTitle,
            data.content[i].statusTitle
        ];
        
        const scopedPath = "asset/" + data.content[i].assetId + "/tasks/" + data.content[i].taskId + "/work-orders/" + id;
        processedRowData.push({
            "id": id,
            "itemRoute": "/manage/" + scopedPath + "/view",
            "cells": row,
            "scopedUrl": "http://localhost:8080/api/v1/" + scopedPath
        });

        console.log(processedRowData)
    }

</script>

<ManageTab columns={columns} dataRows={processedRowData} />

<style></style>