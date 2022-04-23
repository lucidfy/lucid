<script lang="ts">
  import '$lib/tailwind.css'
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

<div id="body">
  <slot />
</div>
