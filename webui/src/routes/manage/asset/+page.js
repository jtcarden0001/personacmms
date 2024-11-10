export async function load({ fetch }) {
    const response = await fetch("http://localhost:8080/api/v1/asset");
    const wos = await response.json();

    return { "content": wos };
}