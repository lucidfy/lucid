<script lang="ts">
  import { fly } from 'svelte/transition'

  export let title: String
</script>

<svelte:head>
  <title>{title}</title>
</svelte:head>

<main class="bg-orange-600 flex h-screen mt-auto">
  <div class="text-white m-auto w-full">
    <div in:fly={{ y: 200, duration: 500 }}>
      <slot />
    </div>
  </div>
</main>
