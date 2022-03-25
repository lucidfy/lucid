<script>
  import Table from '$src/lib/guest/Table.svelte'
  import Gorvel from '$src/lib/guest/Gorvel.svelte';

  export let title

  export let data
  export let csrfField
  export let links
</script>

<Gorvel {title}>
  <div class="my-5 shadow-orange-800 shadow-lg">
    <Table {title} {data} {csrfField} />
  </div>

  {@html links}
</Gorvel>
