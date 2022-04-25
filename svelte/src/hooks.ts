import cookie from 'cookie';
import type { Handle } from '@sveltejs/kit';

/** @type {import('@sveltejs/kit').Handle} */
export const handle: Handle = async ({ event, resolve }) => {
	const response = await resolve(event)

	if (event.locals.setcookie) {
		response.headers.set('set-cookie', event.locals.setcookie)
	}

	return response;
};
