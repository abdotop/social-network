<template>
    <!-- <div class="w-full"> -->
    <div class="flex 2xl:gap-12 gap-10 mt-8 max-lg:flex-col" id="js-oversized">

        <div class="lg:space-y-4 lg:pb-8 max-lg:grid sm:grid-cols-2 max-lg:gap-6 uk-sticky"
            uk-sticky="media: 1024; end: #js-oversized; offset: 80" style="width: 100%;">

            <div class="box p-5 px-6">


                <div class="flex items-ce justify-between text-black dark:text-white">
                    <h3 class="font-bold text-lg"> Intro </h3>
                    <a v-if="data.followStatus === 'self'" href="/settings" class="text-sm text-blue-500">Edit</a>
                </div>

                <ul class="text-gray-700 space-y-4 mt-4 text-sm dark:text-white/80">
                    <li class="flex items-center gap-3">
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                            stroke="currentColor" class="w-6 h-6">
                            <path stroke-linecap="round" stroke-linejoin="round"
                                d="M19.5 14.25v-2.625a3.375 3.375 0 00-3.375-3.375h-1.5A1.125 1.125 0 0113.5 7.125v-1.5a3.375 3.375 0 00-3.375-3.375H8.25m0 12.75h7.5m-7.5 3H12M10.5 2.25H5.625c-.621 0-1.125.504-1.125 1.125v17.25c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125V11.25a9 9 0 00-9-9z">
                            </path>
                        </svg>
                        <div> Posts <span class="font-semibold text-black dark:text-white"> {{ data.NumberOfPosts }}
                            </span> </div>
                    </li>

                    <!-- <li class="flex items-center gap-3">
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                            stroke="currentColor" class="w-6 h-6">
                            <path stroke-linecap="round" stroke-linejoin="round"
                                d="M11.48 3.499a.562.562 0 011.04 0l2.125 5.111a.563.563 0 00.475.345l5.518.442c.499.04.701.663.321.988l-4.204 3.602a.563.563 0 00-.182.557l1.285 5.385a.562.562 0 01-.84.61l-4.725-2.885a.563.563 0 00-.586 0L6.982 20.54a.562.562 0 01-.84-.61l1.285-5.386a.562.562 0 00-.182-.557l-4.204-3.602a.563.563 0 01.321-.988l5.518-.442a.563.563 0 00.475-.345L11.48 3.5z">
                            </path>
                        </svg>
                        <div> Rating · <span class="font-semibold text-black dark:text-white"> 4.6</span> (59
                            reviews)
                        </div>
                    </li> -->
                    <li class="flex items-center gap-3">
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                            stroke="currentColor" class="w-6 h-6">
                            <path stroke-linecap="round" stroke-linejoin="round"
                                d="M12.75 19.5v-.75a7.5 7.5 0 00-7.5-7.5H4.5m0-6.75h.75c7.87 0 14.25 6.38 14.25 14.25v.75M6 18.75a.75.75 0 11-1.5 0 .75.75 0 011.5 0z">
                            </path>
                        </svg>
                        <div> Followed By <span class="font-semibold text-black dark:text-white"> {{ data.follow }}
                                <span v-if="data.follow > 1">People</span>
                                <span v-else>Person</span>
                            </span> </div>
                    </li>

                </ul>

            </div>
            <div class="box p-5 px-6">
                <div class="flex items-ce justify-between text-black dark:text-white">
                    <h3 class="font-bold text-lg"> Full Name </h3>
                    <a v-if="data.followStatus === 'self'" href="/settings" class="text-sm text-blue-500">Edit</a>
                </div>
                <div class="sm:px-4 p-2.5 pt-0">
                    <p class="font-normal"> {{ data.firstname }} {{ data.lastname }} </p>
                </div>
                <div class="flex items-ce justify-between text-black dark:text-white">
                    <h3 class="font-bold text-lg"> Nickname </h3>
                </div>
                <div class="sm:px-4 p-2.5 pt-0">
                    <p class="font-normal"> @{{ data.nickname }} </p>
                </div>
                <div class="flex items-ce justify-between text-black dark:text-white">
                    <h3 class="font-bold text-lg"> Email </h3>
                </div>
                <div class="sm:px-4 p-2.5 pt-0">
                    <p class="font-normal"> @{{ data.email }} </p>
                </div>
                <div class="flex items-ce justify-between text-black dark:text-white">
                    <h3 class="font-bold text-lg"> Birthday Date </h3>
                </div>
                <div class="sm:px-4 p-2.5 pt-0">
                    <p class="font-normal"> {{ data.birthday.split('T')[0] }} </p>
                </div>
                <div class="flex items-ce justify-between text-black dark:text-white">
                    <h3 class="font-bold text-lg"> About Me </h3>
                </div>
                <div class="sm:px-4 p-2.5 pt-0">
                    <p class="font-normal"> {{ data.about }} </p>
                </div>
            </div>

            <!-- group media  -->




            <!-- related pages  -->



            <!-- related pages  -->



        </div>
        <div class="uk-sticky-placeholder" style="height: 1213px; width: 400px; margin: 0px;" hidden="">
        </div>

    </div>
</template>

<script>
export default {
    props: {
        data: {
            type: Object,
            required: true
        }
    }
}
</script>