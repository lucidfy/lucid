const base = 'http://localhost:8080';

export function api(method: string, resource?: string, data?: Record<string, unknown>) {
	let fullurl = base
	if (resource) {
		fullurl = `${base}/${resource}`
	}

	return fetch(fullurl, {
		method,
		headers: {
			'content-type': 'application/json',
			'accept': 'application/json',
		},
		body: data && JSON.stringify(data)
	});
}
