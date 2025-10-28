import {apiFetch} from './fetcher';

export type MenuItem = {
    link: string;
    content: string;
    icon?: string;
};

export async function getMenu(astro?: any): Promise<MenuItem[]> {
    const data = await apiFetch<{ data: MenuItem[] }>('/api/v1/menu', {astro});
    return data.data;
}
