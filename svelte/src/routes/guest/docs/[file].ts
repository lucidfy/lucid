import { api } from '$src/routes/_api';
import type { RequestHandler } from '@sveltejs/kit';

// Lists of users
export const get: RequestHandler = async (event) => {
	const url = event.url
	const response = await api({
		method: 'get',
		resource: `users?pagination_url=${url.origin}${url.pathname}&${url.searchParams.toString()}`,
		event,
	});

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
		status: response.status,
	};
};

// Deleting a user
export const del: RequestHandler = async (event) => {
	const request = event.request
	const form = await request.formData();
	const userId = form.has('id') ? form.get('id') : undefined

	if (userId == undefined) {
		return {
			body: {
				error: "Unknown ID to be deleted!"
			}
		};
	}

	const response = await api({
		method: 'delete',
		resource: `users/${userId}`,
		event,
	});

	if (response.status === 404) {
		return {
			body: []
		};
	}

	if (response.status >= 200 && response.status < 300 || response.status === 401) {
		return {
			status: response.status,
			body: await response.json()
		};
	}

	return {
		status: response.status
	};
};
