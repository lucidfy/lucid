import cookie from 'cookie';
import type { Handle } from '@sveltejs/kit';

/** @type {import('@sveltejs/kit').Handle} */
export const handle: Handle = async ({ event, resolve }) => {
	const cookies = cookie.parse(event.request.headers.get('cookie') || '');
	event.locals.gorvelsession = cookies.gorvel_session;

	// set all requests to have the gorvel_session
	if (event.locals.gorvelcookie) {
		event.request.headers.set('gorvel_session', event.locals.gorvelcookie)
	}

	const response = await resolve(event);

	// if the cookies.gorvel_session is empty, then
	// set cookie based on what gorvelcookie provided
	if (!cookies.gorvel_session && event.locals.gorvelcookie) {
		response.headers.set('set-cookie', event.locals.gorvelcookie)
	}

	return response;
};
