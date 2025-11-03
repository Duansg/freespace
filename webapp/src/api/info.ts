import {apiFetch} from './fetcher';

export type Info = {
    email: string;
    github: string;
    img: string;
    introduction: string;
    name?: string;
};

export async function getInfo(astro?: any): Promise<Info[]> {
    const data = await apiFetch<{ data: Info[] }>('/api/v1/info', {astro});
    return data.data;
}
