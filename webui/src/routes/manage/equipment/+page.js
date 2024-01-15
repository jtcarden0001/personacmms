export async function load({ fetch }) {
    const response = await fetch("http://localhost:8080/api/v1/equipment");
    const wos = await response.json();

    return { "content": wos };
}