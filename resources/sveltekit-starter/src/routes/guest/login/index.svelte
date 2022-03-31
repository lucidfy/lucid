<script>
  import Gorvel from '$src/lib/guest/Gorvel.svelte'

  export let title = 'auth'
</script>

<Gorvel {title}>
  <div class="h-screen flex raleway">
    <div
      class="flex w-1/2 bg-gradient-to-tr from-gray2-200 to-gray-200 i justify-around items-center"
    >
      <div>
        <h1 class="text-gray2-700 font-thin text-6xl tracking-widest">Gorvel</h1>
      </div>
    </div>
    <div class="flex w-11/12 justify-center items-center bg-gray2-900">
      <form class="w-6/12">
        <h1 class="text-2xl mb-1">Hello Again!</h1>
        <p class="text-sm font-normal text-gray-400 mb-7">Welcome Back</p>
        <div class="flex items-center border-2 py-2 px-3 rounded mb-4">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-5 w-5 text-gray-400"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M16 12a4 4 0 10-8 0 4 4 0 008 0zm0 0v1.5a2.5 2.5 0 005 0V12a9 9 0 10-9 9m4.5-1.206a8.959 8.959 0 01-4.5 1.207"
            />
          </svg>
          <input
            class="pl-2 focus:outline-none border-none w-full bg-transparent"
            type="email"
            name="email"
            placeholder="Email Address"
          />
        </div>
        <div class="flex items-center border-2 py-2 px-3 rounded">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-5 w-5 text-gray-400"
            viewBox="0 0 20 20"
            fill="currentColor"
          >
            <path
              fill-rule="evenodd"
              d="M5 9V7a5 5 0 0110 0v2a2 2 0 012 2v5a2 2 0 01-2 2H5a2 2 0 01-2-2v-5a2 2 0 012-2zm8-2v2H7V7a3 3 0 016 0z"
              clip-rule="evenodd"
            />
          </svg>
          <input
            class="pl-2 focus:outline-none border-none w-full bg-transparent"
            type="text"
            name="password"
            placeholder="Password"
          />
        </div>
        <button
          type="submit"
          class="
            block w-full mt-5
            uppercase text-xs
            bg-transparent hover:bg-white
            text-white hover:text-gray2-700
            font-light
            py-2 px-4
            border border-white-500 hover:border
            rounded
          "
          >Login</button
        >
        <!-- <span class="text-sm ml-2 hover:text-blue-500 cursor-pointer">Forgot Password ?</span> -->
      </form>
    </div>
  </div>
</Gorvel>
