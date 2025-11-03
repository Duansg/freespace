import {apiFetch} from './fetcher';
import { marked } from 'marked';
import sanitizeHtml from 'sanitize-html';

export type Blog = {
    id: string;
    title: string;
    content: string;
    img: string;
    pubDate: Date;
    contentHtml?: string;
};

const SANITIZE_OPTIONS = {
    allowedTags: sanitizeHtml.defaults.allowedTags.concat([
        'img','h1','h2','h3','h4','h5','h6','pre','code','figure','figcaption','table','thead','tbody','tr','th','td'
    ]),
    allowedAttributes: {
        '*': ['class', 'id', 'style'],
        a: ['href', 'name', 'target', 'rel'],
        img: ['src', 'alt', 'width', 'height', 'loading'],
        code: ['class']
    },
    transformTags: {
        'a': (tagName, attribs) => {
            if (attribs && attribs.href && attribs.href.startsWith('http')) {
                attribs.target = '_blank';
                attribs.rel = 'noopener noreferrer';
            }
            return { tagName, attribs };
        }
    }
};

export async function getBlogs(astro?: any, limit?: Number): Promise<Blog[]> {
    const data = await apiFetch<{ data: Blog[] }>('/api/v1/blogs?limit=' + limit, {astro});
    return data.data;
}


export async function getBlogInfo(astro?: any, id?: string): Promise<Blog> {
    const data = await apiFetch<{ data: Blog }>('/api/v1/getBlogInfo?id=' + id, {astro});


    const unsafeHtml = marked.parse(data.data.content || '');

    const cleanHtml = sanitizeHtml(unsafeHtml);
    data.data.contentHtml = cleanHtml
    return data.data;
}