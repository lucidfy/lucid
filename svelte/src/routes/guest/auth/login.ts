import { api } from '$src/routes/_api';
import type { RequestHandler } from '@sveltejs/kit';

// Creating a user
export const post: RequestHandler = async (event) => {
	const form = await event.request.formData();
	const email = form.has('email') ? form.get('email') : undefined

	const response = await api({
        method: 'post',
        resource: import.meta.env.VITE_LOGIN_PATH,
        data: {
            email,
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
		const data = await response.json()
		return {
			status: response.status,
			body: {
				email: email.toString(),
				data,
			}
		};
	}

	return {
		status: response.status
	};
};
