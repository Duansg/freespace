export type FetcherOptions = {
    astro?: any;
};

export async function apiFetch<T>(path: string, options?: FetcherOptions): Promise<T> {
    let url: URL;

    try {
        if (options?.astro) {
            const reqUrl = new URL(options.astro.request.url);
            const backend = import.meta.env.BACKEND_URL || `http://${reqUrl.hostname}:8240`;
            url = new URL(path, backend);
        } else {
            url = new URL(path, import.meta.env.BACKEND_URL || window.location.origin);
        }

        const res = await fetch(url.toString(), {
            method: 'GET',
            headers: {'Content-Type': 'application/json'},
        });
        if (!res.ok) throw new Error(`Failed to fetch ${url}: ${res.status}`);
        return res.json();
    } catch (err) {
        console.error("API request failed:", err);
        throw err;
    }

}
