<template>
    <div>
        <b-alert variant="danger" class="text-center" v-if="errors.error" show>{{ errors.error }}</b-alert>
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
                <b-col class="bg-white p-0 profile col-md-12 col-lg-4 mt-5">
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
                        <b-button variant="success" v-if="!isMember">Join</b-button>
                        <b-button variant="outline-primary" v-if="isMember"
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
                </b-col>
            </div>
        </template>
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
            }
        },
        mounted() {
            this.getTag()
        }
    }
</script>

<style scoped lang="scss">
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
