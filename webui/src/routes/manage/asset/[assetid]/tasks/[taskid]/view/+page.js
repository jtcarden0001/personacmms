export async function load({ fetch, params }) {
    const request = "http://localhost:8080/api/v1/asset/" + params.eqid + "/tasks/" + params.taskid
    const response = await fetch(request);
    const task = await response.json();

    return task;
}