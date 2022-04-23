import type { RequestEvent } from "@sveltejs/kit/types/internal";

const base = 'http://localhost:8080';

interface ApiParams {
	method: string;
	event: RequestEvent<Record<string, string>>;
	resource?: string;
	data?: Record<string, unknown>;
}

export async function api(params: ApiParams) {
	let fullurl = base
	if (params.resource) {
		fullurl = `${base}/${params.resource}`
	}

	const response = await fetch(fullurl, {
		method: params.method,
		headers: {
			'content-type': 'application/json',
			'accept': 'application/json',
			'cookie': params.event.request.headers.get('cookie') || '',
		},
		body: params.data && JSON.stringify(params.data)
	})

	params.event.locals.setcookie = response.headers.get('set-cookie')

	return response;
}
