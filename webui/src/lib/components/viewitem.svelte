<script>
  import BackButton from "$lib/components/backbutton.svelte";

  // json represenations of the item to view
  export let content;

  let entries = Object.entries(content);
</script>

<BackButton />
<div>
  <h1 class="text-2xl font-bold">View Item</h1>
  {#each entries as [key, value]}
    <div class="flex flex-row">
      <p class="font-bold">{key}</p>
      <p>{value}</p>
    </div>
  {/each}
</div>

<style></style>
```