<script>
    let basePath = '/manage';
    let woPath = basePath + '/workorders';
    let eqPath = basePath + '/equipment';
    let taskPath = basePath + '/tasks';
    let toolPath = basePath + '/tools';
    let coPath = basePath + '/consumables';
</script>
<div class="h-4/5 flex items-center border-r-4 rounded-md">
  <ul class="menu bg-transparent rounded-box m-5">
    <li><a href={woPath}>Work Orders</a></li>
    <li><a href={eqPath}>Equipment</a></li>
    <li><a href={taskPath}>Tasks</a></li>
    <li><a href={toolPath}>Tools</a></li>
    <li><a href={coPath}>Consumables</a></li>
  </ul>
</div>

<style></style>