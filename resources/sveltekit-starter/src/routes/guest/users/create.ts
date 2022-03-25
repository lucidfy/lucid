import { api } from '$src/routes/_api';
import type { RequestHandler } from '@sveltejs/kit';

export const get: RequestHandler = async ({ locals }) => {
	const response = await api('get', 'users/create');

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

export const post: RequestHandler = async ({ request, params }) => {
	const form = await request.formData();

	const response = await api('post', 'users', {
		'name': form.has('name') ? form.get('name') : undefined,
		'email': form.has('email') ? form.get('email') : undefined,
		'password': form.has('password') ? form.get('password') : undefined,
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
