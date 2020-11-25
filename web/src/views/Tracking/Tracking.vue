<template>
    <b-container class="mt-3">
        <template v-if="currentUser.role === 'admin' || isManager">
            <div class="d-flex justify-content-end">
                <a href="/tracking/admin">
                    <b-button
                            class="mr-3"
                            v-if="currentUser.role === 'admin'"
                            variant="outline-info"
                    >Admin Panel</b-button
                    >
                </a>
                <a href="/tracking/manager"
                ><b-button v-if="isManager" variant="outline-info"
                >Creator Panel</b-button
                ></a
                >
            </div>
        </template>
    </b-container>
</template>

<script>
    export default {
        name: "Tracking",
        computed: {
            currentUser() {
                return this.$store.getters.currentUser;
            }
        },
        data() {
            return {
                errors: {},
                isManager: true,
                blogs: []
            };
        },
        methods: {
            getIsOwner() {
                window.axios
                    .get("api/manager/check/ismanager/")
                    .then(res => {
                        this.isManager = res.data.isManager;
                    })
                    .catch(rej => {
                        console.log(rej.response.data.error)
                    });
            },
            getBlogs() {
                window.axios
                    .get("api/blog/all/")
                    .then(res => {
                        this.blogs = res.data.blogs;
                    })
                    .catch(rej => {
                        console.log(rej.response.data.error)
                    });
            },
            getImgUrl(image) {
                let images = process.env.VUE_APP_API + "/static" + image;
                return images;
            },
            getDate(date) {
                date = new Date(date);
                const day = date.getDate();
                const month = date.toLocaleString("default", { month: "long" });
                const year = date.getFullYear()
                return day + " " + month + " " + year;
            },
        },
        created() {
        }
    }
</script>

<style scoped lang="scss">
    .blog{
        width: 100%;
        cursor: pointer;
        height: 400px;
        margin-bottom: 80px;
        .image{
            width: 60%;
            img{
                width: 100%;
                height: 100%;
            }
        }
        .content{
            display: flex;
            justify-content: space-between;
            flex-direction: column;
            padding-left: 10px;
            .info{
                display: flex;
                font-weight: bolder;
                align-items: center;
            }
            p{
                margin-bottom: 0;
            }
        }
    }
</style>