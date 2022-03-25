<script lang="ts">
  export let title: string
  export let success: string
  export let error: string

  export let fails: Object
  export let inputs: { name: string; email: string; }
  export let record: { name: string; email: string; id: number; }
  export let isCreate: boolean
  export let isShow: boolean
  export let csrfField: string
</script>

<div class="text-white m-auto w-6/12">
  {#if success}
    <div
      class="p-4 mb-4 text-sm text-green-700 bg-green-100 rounded-lg dark:bg-green-200 dark:text-green-800"
      role="alert"
    >
      {success}
    </div>
  {/if}

  {#if error}
    <div
      class="p-4 mb-4 text-sm text-red-700 bg-red-100 rounded-lg dark:bg-red-200 dark:text-red-800"
      role="alert"
    >
      {error}
    </div>
  {/if}

  {#if fails && Object.keys(fails).length > 0}
    <div
      class="p-4 mb-4 text-sm text-red-700 bg-red-100 rounded-lg dark:bg-red-200 dark:text-red-800"
      role="alert"
    >
      Validation Error!
      <ul class="list-disc pl-5">
        {#each  Object.keys(fails) as key, i (i)}
          <li>{fails[key]}</li>
        {/each}
      </ul>
    </div>
  {/if}

  <div class="flex justify-between mb-3">
    <div class="text-xs inline-block my-auto">
      <span class="border-dotted border-gray-400 border-b"><a href="/">Home</a></span>
      <i class="mx-2">&#8725;</i>
      <span><a href="/users">Users List</a></span>
      <i class="mx-2">&#8725;</i>
      <span>{title}</span>
    </div>
  </div>

  <div class="-my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
    <div class="py-2 align-middle max-w-full sm:px-6 lg:px-8">
      <div class="p-5 bg-gray-100 shadow overflow-hidden border-b border-gray-200 sm:rounded-lg">
        <form
          class="w-full"
          method="POST"
          action={isCreate ? '/users/create' : `/users/${record.id}/edit`}
        >
          <div class="flex flex-wrap -mx-3 mb-6">
            <div class="w-1/2 md:w-full px-3 mb-6 md:mb-0">
              <label
                class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2"
                for="grid-name"
              >
                Name
              </label>
              <input
                class="appearance-none block w-full bg-gray-200 text-gray-700 border border-red-500 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white"
                id="grid-name"
                type="text"
                placeholder="John"
                name="name"
                required={isCreate}
                value={isCreate ? inputs?.name ?? '' : record.name}
                disabled={isShow}
              />
            </div>
            <div class="w-1/2 md:w-full px-3 mb-6 md:mb-0">
              <label
                class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2"
                for="grid-email"
              >
                Email
              </label>
              <input
                class="appearance-none block w-full bg-gray-200 text-gray-700 border border-red-500 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white"
                id="grid-email"
                type="text"
                placeholder="john@doe.com"
                name="email"
                required={isCreate}
                value={isCreate ? inputs?.email ?? '' : record.email}
                disabled={isShow}
              />
            </div>

            {#if !isShow}
              <div class="w-1/2 md:w-full px-3 mb-6 md:mb-0">
                <label
                  class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2"
                  for="grid-password"
                >
                  Password
                </label>
                <input
                  class="appearance-none block w-full bg-gray-200 text-gray-700 border border-red-500 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white"
                  id="grid-password"
                  type="password"
                  placeholder="*********"
                  name="password"
                  required={isCreate}
                  disabled={isShow}
                />
                {#if !isCreate && !isShow}
                  <p class="text-red-500 text-xs italic">
                    Leave it blank if you don't want to update the password!
                  </p>
                {/if}
              </div>
            {/if}

            {#if isShow}
              <div class="w-1/2 md:w-full px-3 mb-6 mt-5 md:mb-0 text-center">
                <a
                  href={`/users/${record.id}/edit`}
                  class="
                        uppercase text-xs
                        bg-gray-900 hover:bg-white
                        text-white hover:text-gray-700
                        font-light
                        py-3 px-6
                        border border-gray-700 hover:border
                        rounded
                        ">Edit</a
                >
              </div>
            {:else}
              {@html csrfField}
              <div class="w-1/2 md:w-full px-3 mb-6 md:mb-0 text-right">
                <button
                  class="
                        uppercase text-xs
                        bg-gray-900 hover:bg-white
                        text-white hover:text-gray-700
                        font-light
                        py-3 px-6
                        border border-gray-700 hover:border
                        rounded"
                  type="submit">Save</button
                >
              </div>
            {/if}
          </div>
        </form>
      </div>
    </div>
  </div>
</div>
