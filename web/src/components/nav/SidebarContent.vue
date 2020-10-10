<template>
    <div>
        <b-row class="text-center">
            <b-col>
                <b-avatar size="6rem" variant="info" src="https://placekitten.com/300/300"></b-avatar>
            </b-col>
            <div class="w-100"></div>
            <b-col>{{ currentUser.userdata.name + " " + currentUser.userdata.lastname }}</b-col>
        </b-row>
    </div>
</template>

<script>
    export default {
        name: "SidebarContent",
        computed: {
            currentUser() {
                return this.$store.getters.currentUser;
            }
        }
    }
</script>

<style scoped>

</style>