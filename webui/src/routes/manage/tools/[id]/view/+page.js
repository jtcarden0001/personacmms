export async function load({ fetch, params }) {
    const request = "http://localhost:8080/api/v1/tools/" + params.id
    const response = await fetch(request);
    const wo = await response.json();

    return wo;
}