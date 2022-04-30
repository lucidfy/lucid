<script lang="ts">
  export let title: string
  export let IpAddress: string
  export let userAgent: string
  console.log("Here's the data we received at lucid endpoint", title, IpAddress, userAgent)
</script>

<svelte:head>
  <title>{title}</title>
</svelte:head>

<main class="bg-orange-600 flex h-screen raleway font-thin">
  <div class="text-white m-auto">
    <h1 class="text-8xl mb-5">
        Welcome to <a href="https://github.com/daison12006013/lucid" class="font-light">Lucid!</a>
    </h1>

    <div class="text-center mt-5">
      <p class="my-3">Running using SvelteKit</p>
      <a class="
          uppercase text-xs
          bg-transparent hover:bg-white
          text-white hover:text-gray2-700
          font-light
          py-2 px-4
          border border-white-500 hover:border
          rounded" href="/users"
          >Sample Forms</a>

      <a class="
          uppercase text-xs
          bg-transparent hover:bg-white
          text-white hover:text-gray2-700
          font-light
          py-2 px-4
          border border-white-500 hover:border
          rounded" href="{import.meta.env.VITE_LOGIN_PATH}"
          >Login</a>
    </div>
  </div>
</main>
