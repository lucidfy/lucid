import { api } from '$src/routes/_api';
import type { RequestHandler } from '@sveltejs/kit';

// Getting the create details
export const get: RequestHandler = async (event) => {
	const response = await api({method: 'get', resource: 'users/create', event});

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

// Creating a user
export const post: RequestHandler = async (event) => {
	const form = await event.request.formData();

	const response = await api({
		method: 'post',
		resource: 'users',
		data: {
			'name': form.has('name') ? form.get('name') : undefined,
			'email': form.has('email') ? form.get('email') : undefined,
			'password': form.has('password') ? form.get('password') : undefined,
		},
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
