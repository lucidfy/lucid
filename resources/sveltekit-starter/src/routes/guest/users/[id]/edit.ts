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

export const post: RequestHandler = async ({ request, params, locals }) => {
	const form = await request.formData();

	const response = await api('put', `users/${params.id}`, {
		'name': form.has('name') ? form.get('name') : undefined,
		'email': form.has('email') ? form.get('email') : undefined,
		'password': form.has('password') ? form.get('password') : undefined,
	});

	if (response.status === 404) {
		return {
			body: []
		};
	}

	if (response.status === 200) {
		return {
			body: await response.json(),
		};
	}

	return {
		status: response.status
	};
};
