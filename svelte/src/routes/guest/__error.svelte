<script context="module">
  /** @type {import('@sveltejs/kit').ErrorLoad} */
  export function load({ error, status }) {
    return {
      props: {
        code: status,
        error: error,
      }
    };
  }
</script>

<script>
  export let code;
  export let error;
  console.log(error)
</script>

<svelte:head>
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link rel="preconnect" href="https://fonts.gstatic.com">
  <link href="https://fonts.googleapis.com/css2?family=Raleway:ital,wght@0,100;0,200;0,300;0,400;1,100;1,200;1,300&display=swap" rel="stylesheet">
  <style>
    .raleway {
        font-family: 'Raleway', sans-serif;
    }
  </style>
</svelte:head>

<main class="bg-orange-600 flex h-screen raleway m-auto">
  <div class="text-white m-auto font-thin w-10/12">
    <div class="text-center mb-5">
      <h1 class="text-8xl">
        {code}
      </h1>
      <i class="text-sm -mt-3">{import.meta.env.VITE_PROJECT_NAME} &amp; SvelteKit</i>
    </div>

    {#if error}
      <h1 class="text-5xl mb-5">
        { error.name }
      </h1>
      <div class="rounded-md bg-white text-black text-sm p-5 font-normal whitespace-pre-wrap">
        {#if error.stack}
          { error.stack }
        {:else}
          {@html error.message}
        {/if}
      </div>
    {/if}
  </div>
</main>
