type RequestMethod = "GET" | "POST" | "PUT" | "DELETE";

export default class ClientHttp {
    static API_URL: string = 'http://localhost:3000/api' as string;

    static async request(method: RequestMethod, path: string, rawBody?: any) {
        const body = rawBody ? JSON.stringify(rawBody) : undefined;
        const response = await fetch(`${ClientHttp.API_URL}/${path}`, {
            method,
            headers: {
                "Content-Type": "application/json"
            },
            body
        });
        const data = await response.json();
        if (!response.ok) throw new Error(data.message);
        return data;
    }

    static async get(path: string) {
        return ClientHttp.request("GET", path);
    }

    static async post(path: string, body: any) {
        return ClientHttp.request("POST", path, body);
    }

    static async put(path: string, body: any) {
        return ClientHttp.request("PUT", path, body);
    }

    static async delete(path: string) {
        return ClientHttp.request("DELETE", path);
    }
}

export const fetcher = (path: string) => ClientHttp.get(path);