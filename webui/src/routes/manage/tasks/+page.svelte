<script>
    import ManageTab from "$lib/components/manage/managetab.svelte";

    export let data = {"content":[]};

    let columns = ["Title", "Asset"];
    const processedRowData = [];
    for (let i = 0; i < data.content.length; i++) {
        let id = data.content[i].id;
        // ensure this order matches the order of the columns
        let row = [
            data.content[i].title,
            data.content[i].assetId
        ];
        
        processedRowData.push({
            "id": id,
            "itemRoute": "/manage/asset/" + data.content[i].assetId + "/tasks/" + data.content[i].id + "/view",
            "cells": row
        });
    }

</script>

<ManageTab columns={columns} dataRows={processedRowData} />

<style></style>