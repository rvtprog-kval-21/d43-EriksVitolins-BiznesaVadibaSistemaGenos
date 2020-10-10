<template>
    <div>
        <b-row class="text-center">
            <b-col>
                <b-avatar size="6rem" variant="info" :src="getImgUrl"></b-avatar>
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
        },
        methods: {
            getImgUrl() {
                let images =
                    process.env.VUE_APP_API +
                    "/storage/avatars/" +
                    this.user.id +
                    "/avatar.png";
                return images;
            },
        }
    }
</script>

<style scoped>

</style>