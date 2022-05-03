import { api } from '$src/routes/_api';
import type { RequestHandler } from '@sveltejs/kit';

export const get: RequestHandler = async (event) => {
	const url = event.url
	const response = await api({
		method: 'get',
		resource: `/docs/${url.pathname}`,
		event,
	});

	if (response.status === 200) {
		const body = await response.json()
		return {
			body,
		};
	}

	return {
		status: response.status,
	};
};
