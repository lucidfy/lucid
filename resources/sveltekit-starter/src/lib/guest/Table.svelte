<script>
  export let title
  export let data
  export let csrfField
</script>

<div class="flex flex-col">
    <div class="flex flex-row mb-2">
      <div class="text-xs inline-block my-auto basis-1/4">
        <span class="border-dotted border-gray-400 border-b"><a href="/">Home</a></span>
        <i class="mx-2">&#8725;</i>
        <span>{ title }</span>
      </div>

      <div class="text-xs inline-block my-auto basis-1/4">
        <a href="/users/create" class="
          uppercase text-xs
          bg-transparent hover:bg-white
          text-white hover:text-gray-700
          font-light
          py-2 px-4
          border border-white-500 hover:border
          rounded inline-block">Create</a>
      </div>

      <div class="basis-1/4">
        <!-- Add sorting and page limit -->
      </div>

      <div class="w-full basis-1/4">
        <form method="GET" class="">
          {#each data.headers as header, i (i)}
              {#if header.Name == "search" }
                  <input name="{ header.Name }" value="{ header.Input.Value }" type="text" placeholder="{ header.Input.Placeholder }" class="sm:rounded-lg md:rounded-md outline-0 px-3 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider w-full" />
              {:else}
                  <input type="hidden" name="{ header.Name }" value="{ header.Input.Value }" />
              {/if}
          {/each}
        </form>
      </div>
    </div>
    <div class="-my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
      <div class="py-2 align-middle inline-block min-w-full sm:px-6 lg:px-8">
        <div class="shadow overflow-hidden border-b border-gray-200 sm:rounded-lg">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                {#each data.headers as header, i (i)}
                    {#if header.Input.Visible}
                    <th scope="col" class="text-left text-xs font-medium text-gray-500 tracking-wider">
                      <form method="GET">
                          <input
                            class="bg-transparent outline-0 px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                            name="search[{ header.Name }]"
                            placeholder="{ header.Input.Placeholder }"

                            disabled={ !header.Input.CanSearch }
                            value={ header.Input.CanSearch ? header.Input.Value : "" }
                          />

                          {#each data.headers as header2, i (i)}
                              {#if header.Name != header2.Name }
                                  <input type="hidden" name="{ header2.Name }" value="{ header2.Input.Value }" />
                              {/if}
                          {/each}
                      </form>
                    </th>
                    {/if}
                {/each}
                <th scope="col" class="text-center text-xs font-medium text-gray-500 uppercase tracking-wider">Action</th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              {#if data.paginate.total < 1 }
                <tr>
                  <td colspan="5" class="px-6 py-4 whitespace-nowrap text-gray-800 text-sm text-center">
                  No records found.
                  </td>
                </tr>
              {/if}

              {#if data.paginate.items}
                {#each data.paginate.items as record, idx (idx)}
                <tr>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <div class="flex items-center">
                      <div class="flex-shrink-0 h-10 w-10">
                        <img class="h-10 w-10 rounded-full" src="https://abs.twimg.com/sticky/default_profile_images/default_profile_normal.png" alt="">
                      </div>
                      <div class="ml-4">
                        <div class="text-sm font-medium text-gray-900">{ record.name }</div>
                      </div>
                    </div>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <div class="text-sm text-gray-500">{ record.email }</div>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <div class="text-sm text-gray-500">{ record.created_at }</div>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap">
                    <div class="text-sm text-gray-500">{ record.updated_at }</div>
                  </td>
                  <td class="px-6 py-4 whitespace-nowrap text-center text-sm font-medium">
                    <a href="/users/{ record.id }" class="text-indigo-600 hover:text-indigo-900 px-2">View</a>
                    <form class="px-2 inline-block" action="/users?_method=delete" method="POST">
                      {@html csrfField}
                      <input type="hidden" name="id" value="{record.id}">
                      <button class="text-red-600 hover:text-red-900">Delete</button>
                    </form>
                  </td>
                </tr>
                {/each}
              {/if}

            </tbody>
          </table>
        </div>
      </div>
    </div>
</div>
