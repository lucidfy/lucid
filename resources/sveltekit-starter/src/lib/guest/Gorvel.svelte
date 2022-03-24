<script>
  import { fly } from 'svelte/transition'
  import Table from '$src/lib/guest/Table.svelte'

  export let title
  export let success
  export let error
</script>

<svelte:head>
  <title>{title}</title>
</svelte:head>

<main class="bg-orange-600 flex h-screen mt-auto">
  <div class="text-white m-auto w-4/5">
    {#if success}
      <div
        class="p-4 mb-4 text-sm text-green-700 bg-green-100 rounded-lg dark:bg-green-200 dark:text-green-800"
        role="alert"
      >
        success
      </div>
    {/if}

    {#if error}
      <div
        class="p-4 mb-4 text-sm text-red-700 bg-red-100 rounded-lg dark:bg-red-200 dark:text-red-800"
        role="alert"
      >
        error
      </div>
    {/if}

    <div in:fly={{ y: 200, duration: 500 }}>
      <slot />
    </div>
  </div>
</main>
