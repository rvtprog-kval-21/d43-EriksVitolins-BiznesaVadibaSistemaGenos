<template>
    <div>
        <b-alert variant="danger" class="text-center mb-0" v-if="errors.error" show>{{ errors.error }}</b-alert>
        <b-alert variant="success" class="text-center mb-0" v-if="alerts.alert" show>{{ alerts.alert }}</b-alert>
        <template v-if="tag.name">
            <div class="header d-flex align-items-center">
                <div class="container-fluid">
                    <span class="mask bg-gradient opacity-8"></span>
                    <div class="d-flex align-items-center container-fluid">
                        <b-row class="w-75">
                            <b-col cols="8">
                                <h1 class="mb-4">
                                    Welcome to
                                    <span class="title">{{tag.name}}</span>
                                    profile! :)
                                </h1>
                                <p class="mb-5 mt-4">{{ tag.about }}</p>
                            </b-col>
                        </b-row>
                    </div>
                </div>
            </div>
            <div
                    class="body flex-wrap container-fluid d-flex justify-content-between"
            >
                <div class="bg-white p-0 profile col-md-12 col-lg-4 mt-5">
                    <div
                            class="bg-dark h-30 d-flex align-items-end background justify-content-center"
                    >
                        <b-avatar
                                class="avatar"
                                :src="getImgUrl()"
                                variant="primary"
                                size="8rem"
                                text="EV"
                        ></b-avatar>
                    </div>
                    <div class="d-flex p-4 justify-content-between">
                        <b-button @click="joinTag()" variant="success" v-if="!isMember">Join</b-button>
                        <b-button @click="settingsAreOpended = !settingsAreOpended" variant="outline-primary" v-if="isMember"
                        >Settings</b-button
                        >
                    </div>
                    <div class="d-flex justify-content-center mt-3">
                        <h4>{{ tag.name}}</h4>
                    </div>
                    <div class="d-flex grey justify-content-center mt-1">
                        <h6>{{ tag.about }}</h6>
                    </div>
                    <div  v-if="isMember" class="d-flex grey justify-content-center mt-1">
                        <h6>User count: {{ getMemberCount() }}</h6>
                    </div>
                </div>
                <div class="col-lg-7 mt-5 col-md-12 p-4 bg-white card-universal" v-if="isMember">
                    <template v-if="settingsAreOpended">
                        <div class="d-flex justify-content-between">
                            <h4>Settings:</h4>
                            <b-button @click="settingsAreOpended = !settingsAreOpended" variant="outline-primary">Feed</b-button>
                        </div>
                        <div class="d-flex mt-5">
                            <div class="w-50">
                                <h5>Invite User:</h5>
                            </div>
                            <div class="w-50">
                                <h5>User List:</h5>
                            </div>
                        </div>
                        <div class="d-flex mt-5">
                            <div class="w-50">
                                <h5>Edit Name:</h5>
                            </div>
                            <div class="w-50">
                                <h5>Edit About:</h5>
                            </div>
                        </div>
                        <div class="d-flex mt-5">
                           <template v-if="!tag.is_public">
                               <div class="w-50">
                                   <h5>Make Public:</h5>
                               </div>
                           </template>
                            <template v-else>
                                <div class="w-50">
                                    <h5>Make Private:</h5>
                                </div>
                            </template>
                            <div class="w-50">
                                <h5>Change Avatar</h5>
                            </div>
                        </div>
                        <div class="d-flex mt-5">
                            <div class="w-50">
                                <h5>Leave Group:</h5>
                            </div>
                            <div class="w-50">
                                <h5>Delete Group:</h5>
                                <h6 class="mt-1 mb-2">This will delete the tag completely</h6>
                                <b-button variant="danger" v-b-modal.delete-tag>Delete Tag</b-button>
                            </div>
                        </div>
                    </template>
                    <template v-else>
                        <div class="d-flex justify-content-between">
                            <h4>Activity Feed:</h4>
                            <b-button @click="settingsAreOpended = !settingsAreOpended" variant="outline-primary">Settings</b-button>
                        </div>
                    </template>
                </div>
            </div>
        </template>
        <b-modal id="delete-tag" title="Delete this tag?" hide-footer>
            <p class="my-4">Do you wanna delete this tag?</p>
            <div class="d-flex justify-content-between">
                <b-button class="w-40" variant="danger" @click="[$bvModal.hide('delete-tag'), deleteTag()]">Yes</b-button>
                <b-button class="w-40" variant="primary" @click="$bvModal.hide('delete-tag')">No</b-button>
            </div>
        </b-modal>
    </div>
</template>


<script>
    export default {
        name: "TagsProfile",
        data() {
            return {
                errors: {},
                alerts: {},
                tag: {},
                isMember: false,
                settingsAreOpended: false
            };
        },
        methods: {
            getImgUrl() {
                let images = process.env.VUE_APP_API + "/static" + this.tag.avatar;
                return images;
            },
            getTag(){
                this.errors = {};
                const vue = this;
                window.axios
                    .get("/api/tags/tag/" + this.$route.params.id + "/profile")
                    .then(res => {
                        this.tag = res.data.tag
                        this.isMember = res.data.is_member
                    })
                    .catch(function(rej) {
                        vue.errors = { error: rej.response.data.error };
                    });
            },
            getMemberCount() {
                return this.tag.members.length
            },
            joinTag() {
                this.errors = {};
                const vue = this;
                window.axios
                    .post("/api/tags/tag/" + this.$route.params.id + "/join")
                    .then(res => {
                        this.alerts = { alert: res.data.message}
                        this.getTag()
                    })
                    .catch(function(rej) {
                        vue.errors = { error: rej.response.data.error };
                    });
            },

            deleteTag() {
                this.errors = {};
                const vue = this;
                window.axios
                    .post("/api/tags/tag/" + this.$route.params.id + "/delete")
                    .then(() => {
                        this.$router.push("/tags");
                    })
                    .catch(function(rej) {
                        vue.errors = { error: rej.response.data.error };
                    });
            }
        },
        mounted() {
            this.getTag()
        }
    }
</script>

<style scoped lang="scss">
    .w-40{
        width: 40%;
    }
    .header {
        min-height: 400px;
        background-size: cover;
        background: #49599a;
        color: white;
        .about {
            font-size: larger;
        }
        .title {
            color: #fce4ec;
        }
    }
    .body {
        width: 90%;
        position: relative;
        top: -100px;
        .card-universal {
            border-radius: 5px;
            .category-title {
                color: #b0bec5;
                font-weight: lighter;
            }
        }
        .profile {
            min-height: 500px;
            .avatar {
                position: relative;
                bottom: -75px;
            }
            .grey {
                color: #b0bec5;
            }
        }
        .background {
            background-size: cover;
        }
        .h-30 {
            height: 35% !important;
        }
    }
</style>
