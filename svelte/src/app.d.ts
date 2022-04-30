/// <reference types="@sveltejs/kit" />

// See https://kit.svelte.dev/docs#typescript
// for information about these interfaces
declare namespace App {
	interface Locals {
		// sveltekit default
		userid: string;

		// gorvel
		user: any;
		session: string;
	}

	interface Platform { }

	interface Session {
		// gorvel
		user: any;
	}

	interface Stuff { }
}
