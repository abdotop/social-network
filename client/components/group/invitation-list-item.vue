<script setup lang="ts">
const props = defineProps({
    group: {
        type: Object,
        required: true
    },
    memberId: {
        type: String,
        required: true
    },
    joined: {
        type: Boolean,
        default: false
    }
});

const status = ref("invited")

const { acceptInvitation, declineInvitation } = useGroupRequest()


const handleAccept = async () => {
    const response = await acceptInvitation(props.memberId)

    console.log(response);

    if (response.data) {
        status.value = 'accepted'
    }

}

const handleReject = async () => {
    const response = await declineInvitation(props.memberId)

    console.log(response);

    if (response.data) {
        status.value = 'rejected'
    }

}


</script>

<template>
    <UCard class="card w-60 bg-gray-100 dark:bg-gray-950 p-0 dark:bg-dark2 dark:border-slate-800"
        :ui="{ header: { padding: 'p-0 rounded-t-lg overflow-hidden' } }">
        <NuxtLink :href="'/groups/' + props.group.ID">
            <div class="card-media h-24 rounded-t-lg">
                <NuxtImg :src="'http://localhost:8081/' + props.group.BannerURL" class="" alt="" />
            </div>
        </NuxtLink>
        <div class="card-body z-10 relative w-full">
            <NuxtLink :href="'/groups/' + props.group.ID">
                <h5 class="card-title text-xl h-min text-black font-bold text-ellipsis w-full">
                    {{ props.group.Title }}
                </h5>
            </NuxtLink>
            <div class="flex mb-4 text-sm mt-2">
                <div class="md:block hidden" />
                <div>16k members</div>
            </div>
            <div class="flex-row flex gap-2" v-if="!joined">
                <UButton v-if="status ==='invited'" @click="handleAccept"
                    class="items-center bg-blue-500 justify-center flex-1 w-full">
                    accept
                </UButton>
                <UButton v-if="status ==='invited'" @click="handleReject"
                    class="items-center bg-blue-500 justify-center flex-1 w-full">
                    reject
                </UButton>
                <div v-if="status ==='accepted'">
                    Accepted
                </div>
                <div v-if="status ==='rejected'">
                    Rejected
                </div>
            </div>
            <NuxtLink v-else :href="'/groups/' + props.group.ID">
                <UButton class="items-center bg-blue-500 justify-center w-full">
                    View
                </UButton>
            </NuxtLink>
        </div>
    </UCard>
</template>