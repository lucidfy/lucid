import { api } from '$src/routes/_api';
import type { RequestHandler } from '@sveltejs/kit';

export const get: RequestHandler = async ({ params, locals }) => {
	const response = await api({
		method: 'get',
		resource: `users/${params.id}/edit`,
		event: {locals},
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
		status: response.status
	};
};

export const post: RequestHandler = async ({ request, params, locals }) => {
	const form = await request.formData();

	const response = await api({
		method: 'put',
		resource: `users/${params.id}`,
		data: {
			'name': form.has('name') ? form.get('name') : undefined,
			'email': form.has('email') ? form.get('email') : undefined,
			'password': form.has('password') ? form.get('password') : undefined,
		},
		event: {locals},
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
