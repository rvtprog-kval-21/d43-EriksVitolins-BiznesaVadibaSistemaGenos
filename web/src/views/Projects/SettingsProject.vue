<template>
    <b-container class="mt-3">
        <div class="d-flex justify-content-between">
            <h3>Settings</h3>
            <b-button variant="outline-primary" @click="this.$router.push(`/projects/${this.$route.params.id}/see`)">Back</b-button>
        </div>
        <div class="borders d-flex justify-content-between mt-5" v-if="this.user.is_owner">
            <h4>Archive Project</h4>
            <b-button variant="danger" @click="deleteProject() ">Archive</b-button>
        </div>
    </b-container>
</template>

<script>
    export default {
        name: "Settings.vue",
        data() {
            return {
                project: null,
                user: {},
                aboutIsOpened: false,
                isLoading: false
            };
        },
        methods: {
            getUser() {
                const users = this.project.members;
                this.user = users.find(e => e.user.id === this.currentUser.id);
                console.log(this.user)
                if (this.user) {
                    if (this.user.is_admin || !this.user.is_owner) {
                        return
                    }
                }
                this.$router.push(`/projects/${this.$route.params.id}/see`);
            },
            getProject() {
                this.isLoading = true
                window.axios
                    .get(`api/projects/get/${this.$route.params.id}/item/`)
                    .then(res => {
                        this.project = res.data.project;
                        this.getUser();
                        this.isLoading = false
                    })
                    .catch(() => {
                        this.$router.push("/projects");
                    });
            },
            makeToast(text, variant) {
                this.$bvToast.toast(text, {
                    autoHideDelay: 5000,
                    variant: variant,
                    title: "Notification"
                });
            },
            deleteProject() {
                window.axios
                    .get(`api/projects/remove/${this.$route.params.id}/project/`)
                    .then(() => {
                        this.$router.push("/projects");
                    })
                    .catch((rej) => {
                        this.makeToast(rej.response.data.error, "danger");
                    });
            },
        },
        computed: {
            currentUser() {
                return this.$store.getters.currentUser;
            },
        },
        async created() {
            await this.getProject();
        }
    }
</script>

<style scoped>
.borders{
    padding-top: 5px;
    padding-bottom: 5px;
    border-bottom: solid #607d8b 1px;
    border-top: solid #607d8b 1px;
}
</style>