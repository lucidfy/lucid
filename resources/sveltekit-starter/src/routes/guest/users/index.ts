import { api } from '$src/routes/_api';
import type { RequestHandler } from '@sveltejs/kit';

export const get: RequestHandler = async ({ locals }) => {
	// locals.userid comes from src/hooks.js
	const response = await api('get', `users`);

	if (response.status === 404) {
		// start with an empty array
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

export const post: RequestHandler = async ({ request, locals }) => {
	const form = await request.formData();

	await api('post', `users/${locals.userid}`, {
		text: form.get('text')
	});

	return {};
};

// If the user has JavaScript disabled, the URL will change to
// include the method override unless we redirect back to /users
const redirect = {
	status: 303,
	headers: {
		location: '/users'
	}
};

export const patch: RequestHandler = async ({ request, locals }) => {
	const form = await request.formData();

	await api('patch', `users/${locals.userid}/${form.get('uid')}`, {
		text: form.has('text') ? form.get('text') : undefined,
		done: form.has('done') ? !!form.get('done') : undefined
	});

	return redirect;
};

export const del: RequestHandler = async ({ request, locals }) => {
	const form = await request.formData();

	await api('delete', `users/${locals.userid}/${form.get('uid')}`);

	return redirect;
};
