<template>
    <b-container class="mt-3">
        <div class="d-flex justify-content-between">
            <h3>Settings</h3>
            <b-button variant="outline-primary" @click="goBack()">Back</b-button>
        </div>
        <div class="borders d-flex justify-content-between mt-5" v-if="this.user.is_owner">
            <h4>Archive Project</h4>
            <b-button variant="danger" @click="deleteProject() ">Archive</b-button>
        </div>
        <div class="borders d-flex justify-content-between mt-5" v-if="this.user.is_owner">
            <h4>Project name:</h4>
            <b-form-input
                    class="w-75"
                    v-model="name"
                    placeholder="Enter project name"
            ></b-form-input>
            <b-button @click="saveName()" variant="success" v-if="name != ''">Save</b-button>
        </div>
        <div class="borders d-flex justify-content-between mt-5">
                <h4>Project avatar:</h4>
                <b-form-file
                        v-model="avatar"
                        :state="Boolean(avatar)"
                        class="w-75"
                        placeholder="Choose a photo or drop it here..."
                        drop-placeholder="Drop photo here..."
                        accept="image/*"
                ></b-form-file>
            <b-button variant="success"  @click="saveAvatar()" v-if="avatar != null">Save</b-button>
        </div>
        <div class="borders mt-5">
        <div class="d-flex justify-content-between">
                <h5>Project About:</h5>
                <b-form-textarea
                        v-model="about"
                        placeholder="Project about(You can use markdown)"
                        rows="9"
                        class="w-75"
                        no-resize
                ></b-form-textarea>
            <b-button variant="success" @click="saveAbout()" v-if="about != ''" >Save</b-button>
            </div>
            <div class="mt-1 mb-4">
                <VueShowdown
                        :markdown="about"
                        flavor="github"
                        :options="{ emoji: true }"
                ></VueShowdown>
            </div>
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
                isLoading: false,
                name:"",
                avatar:null,
                about:"",
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
                        this.name = this.project.name
                        this.about = this.project.about
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
            goBack() {
                this.$router.push(`/projects/${this.$route.params.id}/see`)
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
            saveName() {
                window.axios
                    .post(`api/projects/change/${this.$route.params.id}/name/`, {name: this.name})
                    .then(() => {
                        this.makeToast("Name changed", "success");
                        this.getProject();
                    })
                    .catch((rej) => {
                        this.makeToast(rej.response.data.error, "danger");
                    });
            },
            saveAbout() {
                window.axios
                    .post(`api/projects/change/${this.$route.params.id}/about/`, {about: this.about})
                    .then(() => {
                        this.makeToast("About changed", "success");
                        this.getProject();
                    })
                    .catch((rej) => {
                        this.makeToast(rej.response.data.error, "danger");
                    });
            },
            saveAvatar() {
                let formData = new FormData();
                formData.append("avatar", this.avatar);
                window.axios
                    .post(`api/projects/change/${this.$route.params.id}/avatar/`,formData, {
                        headers: {
                            "Content-Type": "multipart/form-data"
                        }
                    })
                    .then(() => {
                        this.makeToast("Project avatar changed successfully", "success");
                        this.getProject();
                    })
                    .catch(rej => {
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