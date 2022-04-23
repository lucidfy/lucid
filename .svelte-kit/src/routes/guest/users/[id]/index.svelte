<script lang="ts">
  import Gorvel from "$src/lib/guest/Gorvel.svelte";
  import Form from "$src/routes/guest/users/_form.svelte";

  export let title: string
  export let success: string
  export let error: string

  export let fails: any
  export let inputs: { name: string; email: string; }
  export let record: { name: string; email: string; id: number; }
  export let isCreate: boolean
  export let isShow: boolean
  export let csrfField: string
</script>

<Gorvel {title}>
  <Form
    {title}
    {success}
    {error}
    {fails}
    {inputs}
    {record}
    {isCreate}
    {isShow}
    {csrfField}
  />
</Gorvel>
