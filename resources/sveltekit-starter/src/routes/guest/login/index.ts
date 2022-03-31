import { api } from '$src/routes/_api';
import type { RequestHandler } from '@sveltejs/kit';

// Creating a user
export const post: RequestHandler = async ({ request, locals }) => {
	const form = await request.formData();

	const response = await api({
        method: 'post',
        resource: 'auth/login',
        data: {
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
