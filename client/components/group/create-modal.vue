<template lang="">
  <div id="create-group-overlay" class="hidden lg:p-20 " uk-modal="">
    <div
      class=" tt relative overflow-hidden mx-auto bg-white shadow-xl rounded-lg md:w-[520px] w-full dark:bg-dark2"
    >
      <div class="text-center py-4 border-b mb-0 dark:border-slate-700">
        <h2 class="text-sm font-medium text-black">
          Create Group
        </h2>

        <button type="button" class="button-icon absolute top-0 right-0 m-2.5 uk-modal-close">
          <svg
            xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
            stroke="currentColor" class="w-6 h-6"
          >
            <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>

      <form ref="create_group_form" action="">
        <div class="space-y-5 mt-3 p-2">
          <input
            id=""
            class="w-full !text-black placeholder:!text-black !bg-white !border-transparent focus:!border-transparent focus:!ring-transparent !font-normal !text-xl   dark:!text-white-100 dark:placeholder:!text-red-500 dark:!bg-slate-800" name="title" type="text" placeholder="name"
          >
          <textarea id="" class="resize-none w-full" cols="30" rows="10" name="description" placeholder="Description" />
        </div>


        <div class="p-5 flex justify-between items-center">
          <div class="flex items-center gap-2">
            <button type="submit" class="button bg-blue-500 text-white py-2 px-12 text-[14px]">
              Create
            </button>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
const uikit = useUIkit3()

export default {
  mounted() {
    this.$refs.create_group_form.addEventListener('submit', this.submitData)
  },
  methods: {
    async submitData(e) {
      e.preventDefault()
      const formData = new FormData(e.target)
      const { data, pending, error, refresh } = await useFetch('/api/groups/', {
        method: 'post',
        body: JSON.stringify(Object.fromEntries(formData.entries())),
        onResponse({ request, response, options }) {
          const data = JSON.parse(response._data)
          const gid = data.data.ID
          uikit?.modal("#create-group-overlay").hide()
          navigateTo(`/groups/${gid}`)
        },
      })
    }
  }
}



</script>

<style lang="">

</style>