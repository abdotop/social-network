<template>
  <main id="site__main" class="2xl:ml-[--w-side]  xl:ml-[--w-side-sm] p-2.5 h-[calc(100vh-var(--m-top))] mt-[--m-top] overflow-y-auto">

    <div class="max-w-[1065px] mx-auto">

      <!-- Cover  -->
      <div class="bg-white shadow lg:rounded-b-2xl lg:-mt-10 dark:bg-dark2">

        <!-- cover -->
        <Cover :nickname="nickname" :data="data" />

        <div v-if="(data.followStatus == 'accepted' && !data.isPublic) || data.isPublic" class="flex items-center justify-between border-t border-gray-100 px-2 dark:border-slate-700">

          <nav
            class="flex gap-0.5 rounded-xl overflow-hidden -mb-px text-gray-500 font-medium text-sm overflow-x-auto dark:text-white">
            <button @click="navigate('posts')" class="inline-block py-3 leading-8 px-3.5"
              :class="{ 'border-b-2 border-blue-600 text-blue-600': status === 'posts' }">
              Timeline
            </button>
            <button @click="navigate('about')" class="inline-block py-3 leading-8 px-3.5"
              :class="{ 'border-b-2 border-blue-600 text-blue-600': status === 'about' }">
              About
            </button>
          </nav>

          <div
            class="flex items-center gap-1 text-sm p-3 bg-secondery py-2 mr-2 rounded-xl max-lg:hidden dark:bg-white/5">
            <i class='bx bx-search text-lg md hydrated'></i>
            <input placeholder="Search .." class="!bg-transparent">
          </div>

        </div>


      </div>

      <div v-if="(data.followStatus == 'accepted' && !data.isPublic) || data.isPublic" class="flex 2xl:gap-12 gap-10 mb-8 max-lg:flex-col" id="js-oversized">

        <!-- feed story -->
        <div v-if="status === 'posts'">
          <FeedStory />
        </div>
        <!-- <post-input idUIselect="select1" /> -->
        <!-- <div class="bg-white shadow lg:rounded-b-2xl lg:-mt-10 dark:bg-dark2"> -->
        <div v-if="status === 'about'">
          <Detail :status="about" :data="data"/>
        </div>
        <!-- </div> -->

      </div>

    </div>

  </main>
</template>
<script>
import Cover from '~/components/Profile/Cover.vue';
import Detail from '~/components/Profile/Detail.vue';
import FeedStory from '~/components/Profile/FeedStory.vue';

export default {
  components: {
    Cover,
    FeedStory,
    Detail
  },
  data() {
    return {
      data: {
        firstname: "",
        lastname: "",
        follow: "",
        following: "",
        avatar: "",
        followStatus: "",
        isPublic: true
      },
      nickname:"",
      status: "posts"
    }
  },
  async mounted() {
    this.nickname = this.$route.params.nickname;
    const response = await getUser(this.nickname, "get")
    if (response && response.status === 200) {
      this.data = response.body;
    } else {
      console.error('Failed to get user data');
      navigateTo("/404")
    }
    this.data.follow = formatFollowersCount(this.data.follow);
    this.data.following = formatFollowersCount(this.data.following);
  },
  methods: {
    async navigate(to) {
      if (to === this.status) return;
      this.status = to;
    }
  }
}
</script>