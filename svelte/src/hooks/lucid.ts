import { api } from '$src/routes/_api';
import cookie, { parse } from 'cookie';
import type { Handle } from '@sveltejs/kit';

export const handle: Handle = async ({ event, resolve }) => {
	// skip the hook to prevent from running if ever
	// we're calling ./setup-docker
	if (process.env.LUCID_DOCKER_SETTING_UP) {
		const response = await resolve(event)
		return response
	}

	let setCookie = null

	if (!event.locals.user) {
		const loggedIn = await api({
			method: 'get',
			resource: import.meta.env.VITE_LOGIN_PATH,
			event,
		});
		setCookie = loggedIn.headers.get('set-cookie')
		event.locals.user = (await loggedIn.json()).user
	}

	const sessionName = import.meta.env.VITE_SESSION_NAME
	const cookies = cookie.parse(event.request.headers.get('cookie') || '')
	event.locals.session = cookies[sessionName]

	const response = await resolve(event)

	if (!event.locals.session && setCookie) {
		response.headers.set('set-cookie', setCookie)
	}

	return response
};

export async function getSession({ locals }) {
	return {
		user: locals.user,
	};
}
