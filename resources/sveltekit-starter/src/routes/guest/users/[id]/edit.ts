import { api } from '$src/routes/_api';
import type { RequestHandler } from '@sveltejs/kit';

export const get: RequestHandler = async ({ params }) => {
	const response = await api('get', `users/${params.id}/edit`);

	if (response.status === 404) {
		return {
			body: []
		};
	}

	if (response.status === 200) {
		return {
			body: await response.json()
		};
	}

	return {
		status: response.status
	};
};
