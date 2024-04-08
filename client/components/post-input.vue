<template lang="">
    <div class="hidden lg:p-20 uk- open" id="create-status" uk-modal="">

        <div
            class="tt relative overflow-hidden mx-auto bg-primary shadow-xl rounded-lg md:w-[520px] w-full dark:bg-dark2">

            <div class="text-center py-4 border-b mb-0 dark:border-slate-700">
                <h2 class="text-sm font-medium text-black"> Create Post </h2>

                <!-- close button -->
                <button type="button" class="button-icon absolute top-0 right-0 m-2.5 uk-modal-close">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                        stroke="currentColor" class="w-6 h-6">
                        <path stroke-linecap="round" stroke-linejoin="round" d="M6 18L18 6M6 6l12 12" />
                    </svg>
                </button>
            </div>

            <form action="" ref="post_form" @submit="handleSubmitPost">

                <div class="space-y-5 mt-3 p-2">
                    <textarea
                        class="w-full !text-black placeholder:!text-black !bg-white !border-transparent focus:!border-transparent focus:!ring-transparent !font-normal !text-xl   dark:!text-white dark:placeholder:!text-white dark:!bg-slate-800"
                        name="post-content-text" id="" rows="6" placeholder="What do you have in mind?" required ></textarea>
                </div>
 
                <div class="flex items-center gap-2 text-sm py-2 px-4 font-medium flex-wrap">
                    <button type="button"
                        class="flex items-center gap-1.5 bg-sky-50 text-sky-600 rounded-full py-1 px-2 border-2 border-sky-100 dark:bg-sky-950 dark:border-sky-900"
                        @click="openFileInput">
                        <i class='bx bx-image text-base' ></i>
                        Image
                    </button>
                    <input type="file" id="photo-input" name="photo" accept="image/*" style="display: none"
                        ref="fileInput">
                </div>
                
                <UAlert v-if="showAlert" icon="ic:baseline-warning" class= "custom-alert" variant="solid" title="Error!" :description = WarningMessage> </UAlert>

                <div  class="p-5 flex justify-between items-center">
                    <div v-if="path == '/' || path== '/profile'">
                        <button
                            class="inline-flex items-center py-1 px-2.5 gap-1 font-medium text-sm rounded-full bg-slate-50 border-2 border-slate-100 group aria-expanded:bg-slate-100 aria-expanded: dark:text-white dark:bg-slate-700 dark:border-slate-600"
                            type="button">
                            {{privacyState}}
                            <i class='bx bx-chevron-down text-base duration-500 group-aria-expanded:rotate-180' ></i>
                        </button>

                        <div  class="p-2 bg-white rounded-lg shadow-lg text-black font-medium border border-slate-100 w-60 dark:bg-slate-700"
                            uk-drop="offset:10;pos: bottom-left; reveal-left;animate-out: true; animation: uk-animation-scale-up uk-transform-origin-bottom-left ; mode:click">

                            <!-- <form>
                        </form> -->
                            <label>
                                <input type="radio" name="radio-status" ref="publicCheck"
                                value = "public" class="peer appearance-none hidden" checked @change="()=>{resetSelectMenu();this.privacyState='Public' }"/>
                                <div
                                    class=" relative flex items-center justify-between cursor-pointer rounded-md p-2 px-3 hover:bg-secondery peer-checked:[&_.active]:block dark:bg-dark3">
                                    <div class="text-sm"> public </div>
                                    <i class='bx bxs-check-circle hidden active absolute -translate-y-1/2 right-2 text-2xl text-blue-600 uk-animation-scale-up'></i>
                                </div>
                            </label>
                            <label>
                                <input type="radio" name="radio-status" ref="privateCheck"
                                    value = "private" class="peer appearance-none hidden"  @change="()=>{
                                      resetSelectMenu()
                                      this.privacyState='Private' }"/>
                                <div
                                    class=" relative flex items-center justify-between cursor-pointer rounded-md p-2 px-3 hover:bg-secondery peer-checked:[&_.active]:block dark:bg-dark3">
                                    <div class="text-sm"> private </div>
                                    <i class='bx bxs-check-circle hidden active absolute -translate-y-1/2 right-2 text-2xl text-blue-600 uk-animation-scale-up'></i>
                                </div>
                            </label>
                            <label>
                              <input type="radio" name="radio-status" ref="privateCheck"
                                    value = "almost private" class="peer appearance-none hidden"  @change="()=>{
                                      this.showSelectUser = true 
                                      this.privacyState = 'almost private'
                                    }"/>
                                <div
                                    class=" relative flex items-center justify-between cursor-pointer rounded-md p-2 px-3 hover:bg-secondery peer-checked:[&_.active]:block dark:bg-dark3">
                                    <div class="text-sm"> private </div>
                                    <i class='bx bxs-check-circle hidden active absolute -translate-y-1/2 right-2 text-2xl text-blue-600 uk-animation-scale-up'></i>
                                </div>
                            </label>

                        </div>
                    </div>

                    <div v-if="showSelectUser &&(path == '/' || path== '/profile') " >
                      <UISelecUser />
                    </div>
                    
                    <div class="flex items-center gap-2">
                        <button ref= "creatPostButon" type="submit" class="button bg-blue-500 text-white py-2 px-12 text-[14px]">
                            Create</button>
                    </div>
                </div>
            </form>

        </div>


    </div>

</template>
<script>
import usePostStore from "~/stores/usePostStore.js";

const postStore = usePostStore()

export default {
  setup() {
    let showAlert = ref(false)
    let showSelectUser = ref(false)
    let path = ref(location.pathname)
    let WarningMessage = ref("")
    let privacyState = ref("Privacy")
    return { showAlert, showSelectUser, path, WarningMessage, privacyState }
  },
  mounted() {
    $('.js-example-basic-multiple').on("select2:select", this.resetCheckbox)
  },
  methods: {
    openFileInput() {
      this.$refs.fileInput.click();
    },
    async handleSubmitPost(e) {
      e.preventDefault();
      let followersSelected = Array.from($('.js-example-basic-multiple').find(':selected'))
      let formdata = new FormData(e.target)
      if (formdata.get("radio-status") == "almost private" && !followersSelected.length) {
        this.WarningMessage = "bro please select some followers"
        this.showAlert = true
        setTimeout(() => this.showAlert = false, 2000)
        return
      }
      if (followersSelected.length == 0 && !formdata.get("radio-status") && (this.path == '/' || this.path == '/profile')) {
        this.WarningMessage = "Choose audiance for your post "
        this.showAlert = true
        setTimeout(() => this.showAlert = false, 2000)
        return
      }
      let jsonFormObject = {
        content: formdata.get("post-content-text"),
        privacy: !this.path.includes("groups") ? 
        (followersSelected.length > 0 ? "almost private" : formdata.get("radio-status"))
         : "group",
        followersSelectedID: followersSelected.length > 0 ? followersSelected.map((v) => v.id) : null,
        group_id: this.path.includes("groups") ? this.path.split('/')[2] : null,
      }
      console.log(jsonFormObject);
      if (formdata.get('photo').name) {
        let body = new FormData()
        body.append('file', formdata.get('photo'))
        let response = await fetch("/api/upload", {
          method: "POST",
          body: body,
        }).then(response => response.json()).catch(err => ({ errors: err }))
        if (response.data) jsonFormObject.image_url = response.data
      }
      try {
        let response = await fetch("/api/post/insert", {
          method: "POST",
          body: JSON.stringify(jsonFormObject),
          headers: {
            'Content-Type': 'application/json'
          }
        }).then(response => response.json())
        console.log(response)
        postStore.addPost(response.body.data)
        UIkit.modal("#create-status").hide();
      } catch (err) {
        console.error(err)
      }
    },  
    resetCheckbox() {
      this.$refs.publicCheck.checked = false
      this.$refs.privateCheck.checked = false

    },
    resetSelectMenu() {
      $('.js-example-basic-multiple').val([]).change()
      this.showSelectUser = false
      console.log(this.showSelectUser);
    },
  },
  watch: {
    '$route'(to, from) {
      this.path = to.path
    },
  },
}
</script>
<style>
.custom-alert {
  background-color: red;
}
</style>