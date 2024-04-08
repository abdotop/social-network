<script>

export default {
    async setup() {
        let userFollowers = await GetAllFollower()
        userFollowers = ref(userFollowers.body)
        return { userFollowers }
    },
    async mounted() {
        let response = await GetAllFollower()
        this.userFollowers = response.body
        console.log(this.userFollowers)
        // Initialisation de Select2 une fois que les données ont été récupérées
        $('.js-example-basic-multiple').select2({
            placeholder: 'Select Followers ..'
        })

    }
}
</script>

<template>
    <select class="js-example-basic-multiple" id="selecUser" style="width: 100%" name="states[]" multiple="multiple">
        <option v-for="follower  in userFollowers" :id="follower.id">{{ follower.firstname }} {{ follower.lastname }}
        </option>
    </select>
</template>


<style>
.select2-dropdown {
    background-color: rgb(30 41 59 / var(--tw-bg-opacity));
    border: none !important;
}

.select2-container--default .select2-selection--multiple {
    background-color: rgb(51 65 85 / 1);
    border: none !important;
    max-height: 8vh;
    overflow-y: auto;
    overflow-x: hidden;

}

.select2-container--default .select2-selection--multiple .select2-selection__choice {
    background-color: rgb(30 41 59 / var(--tw-bg-opacity));
    border: none !important;
}

.select2-container--default .select2-results__option--selected {
    background-color: rgb(51 65 85 / 1);
}
</style>
