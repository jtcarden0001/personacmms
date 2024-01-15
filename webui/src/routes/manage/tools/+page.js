export async function load({ fetch }) {
    const response = await fetch("http://localhost:8080/api/v1/tools");
    const wos = await response.json();

    return { "content": wos };
}