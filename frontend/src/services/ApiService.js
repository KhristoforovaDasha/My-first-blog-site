export async function ApiService(url, params = { headers: {} }) {
    const newParams = {
        ...params,
    };
    const response = await fetch(`http://localhost:3001/${url}`, newParams)
    let data = null;
    data = await response.json()
    console.log(data)
    return data;
}
