<template>
    <!-- <div class="w-full"> -->
    <div class="flex 2xl:gap-12 gap-10 mt-8 max-lg:flex-col" id="js-oversized">

        <div class="lg:space-y-4 lg:pb-8 max-lg:grid sm:grid-cols-2 max-lg:gap-6 uk-sticky"
            uk-sticky="media: 1024; end: #js-oversized; offset: 80" style="width: 100%;">

            <div v-if="data.followers.body.length != 0" class="box p-5 px-6" style="width: 100%">

                <h3 class="font-bold text-lg"> Followers </h3>
                <div v-for="friend in data.followers.body"
                    class="flex items-ce justify-between text-black dark:text-white">
                    <ul class="text-gray-700 space-y-4 mt-4 text-sm dark:text-white/80">
                        <li class="flex items-center gap-3">
                            <img class="w-8 h-8" :src="'http://localhost:8081' + '/uploads/default-avatar.png'" />
                            <nuxt-link :href="'/profile/' + friend.nickname"><span
                                    class="text-lg font-semibold text-black dark:text-white"> {{ friend.firstname }} {{
                                    friend.lastname }}
                                </span> </nuxt-link>
                        </li>
                    </ul>
                </div>
            </div>

            <div v-if="data.followees.body.length != 0" class="box p-5 px-6" style="width: 100%">

                <h3 class="font-bold text-lg"> Followees </h3>
                <div v-for="friend in data.followees.body"
                    class="flex items-ce justify-between text-black dark:text-white">
                    <ul class="text-gray-700 space-y-4 mt-4 text-sm dark:text-white/80">
                        <li class="flex items-center gap-3">
                            <img class="w-8 h-8" :src="'http://localhost:8081' + '/uploads/default-avatar.png'" />
                            <nuxt-link :href="'/profile/' + friend.nickname"><span
                                    class="text-lg font-semibold text-black dark:text-white"> {{ friend.firstname }} {{
                                    friend.lastname }}
                                </span> </nuxt-link>
                        </li>
                    </ul>
                </div>
            </div>

        </div>

    </div>
    <div class="uk-sticky-placeholder" style="height: 1213px; width: 400px; margin: 0px;" hidden="">
    </div>
    
</template>

<script>
export default {
    props: {
        data: {
            type: Object,
            required: true
        },
    },
    async mounted() {
        console.log(data)
    }
}

</script>