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
		},
		body: params.data && JSON.stringify(params.data)
	});

	// const cookies = cookie.parse(response.headers.get('set-cookie') || '');
	// params.event.locals.gorvel_session = cookies.gorvel_session
	// params.event.locals = {
	// 	gorvel_session: cookies.gorvel_session,
	// }
	params.event.locals.gorvelcookie = response.headers.get('set-cookie')

	return response;
}
