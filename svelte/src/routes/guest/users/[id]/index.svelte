<script lang="ts">
  import Lucid from "$src/lib/guest/Lucid.svelte";
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

<Lucid {title}>
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
</Lucid>
