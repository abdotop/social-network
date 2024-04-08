<template>
    <div class="relative overflow-hidden w-full lg:h-60 h-40">
        <img src="assets/images/group/group-cover-4.jpg" alt="" class="h-full w-full object-cover inset-0">

        <!-- overly -->
        <div class="w-full bottom-0 absolute left-0 bg-gradient-to-t from -black/60 pt-10 z-10"></div>
    </div>

    <div class="lg:px-10 md:p-5 p-3">

        <div class="flex flex-col justify-center -mt-20">

            <div class="relative h-20 w-20 mb-4 z-10">
                <div
                    class="relative overflow-hidden rounded-full md:border-[2px] border-gray-100 shrink-0 dark:border-slate-900 shadow">

                    <img :src="'http://localhost:8081/'+data.avatar" class="h-full w-full object-cover inset-0" />
                </div>
            </div>

            <div class="flex lg:items-center justify-between max-md:flex-col max-md:gap-3">

                <div class="flex-1">
                    <h3 class="md:text-2xl text-lg font-bold text-black dark:text-white"> {{ data.firstname }} {{
                        data.lastname }} </h3>
                    <p class=" font-normal text-gray-500 mt-2 flex gap-2 dark:text-white/80">
                        <span> <b class="font-medium text-black dark:text-white">{{ data.follow }}</b>
                            Follow-up</span>
                        <span> • </span>
                        <span> <b class="font-medium text-black dark:text-white">{{ data.following }}</b> Follower(s) </span>
                    </p>
                </div>

                <div class="flex items-center gap-2 mt-1">
                    <a href="/messages" v-if="data.followStatus === 'accepted'"
                        class="button bg-primary text-white flex items-center gap-2  py-2 px-3.5 shadow max-lg:flex-1">
                        <!-- SVG and text here -->
                        <span class="text-sm"> Chat </span>
                    </a>
                    <button v-if="data.followStatus === 'accepted'"
                        class="button bg-secondery flex items-center gap-2  py-2 px-3.5 max-lg:flex-1" @click="follow">
                        <!-- SVG and text here -->
                        <span class="text-sm"> Unfollow </span>
                    </button>
                    <button v-if="data.followStatus === 'none'"
                        class="button bg-secondery flex items-center gap-2  py-2 px-3.5 max-lg:flex-1" @click="follow">
                        <!-- SVG and text here -->
                        <span class="text-sm"> Follow </span>
                    </button>
                    <button v-else-if="data.followStatus === 'requested'"
                        class="button bg-secondery flex items-center gap-2  py-2 px-3.5 max-lg:flex-1" @click="follow">
                        <!-- SVG and text here -->
                        <span class="text-sm"> Cancel </span>
                    </button>
                </div>

            </div>


        </div>
    </div>
</template>


<script>
export default {
    props: {
        nickname: {
            type: String,
            required: true
        },
        data: {
            type: Object,
            required: true
        }
    },
    methods: {
        async follow() {
            switch (this.data.followStatus) {
                case 'none':
                    const ok1 = await useFollow(this.nickname, "follow");
                    if (ok1.ok) {
                        this.data.followStatus = ok1.action
                    }
                    break;
                case 'requested':
                    const ok2 = await useFollow(this.nickname, "unfollow");
                    if (ok2.ok) {
                        this.data.followStatus = "none"
                    }
                    break;
                case 'accepted':
                    const ok3 = await useFollow(this.nickname, "unfollow");
                    if (ok3.ok) {
                        this.data.followStatus = "none"
                    }
                    break;
                default:
                    break;
            }
        }
    }
}
</script>