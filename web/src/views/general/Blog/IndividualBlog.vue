<template>
    <div class="container-blog">
        <template v-if="blog.photo">
            <img class="banner" :src="getImgUrl(this.blog.photo)" alt="">
        </template>
        <div class="d-flex justify-content-between mt-3">
            <div class="info">
                <template v-if="blog.topic">
                    <p>{{blog.topic}}</p>
                </template>
                <b-icon icon="dot"></b-icon>
                <template v-if="blog.publish_at">
                    <p>{{getDate(blog.publish_at)}}</p>
                </template>
            </div>
            <div class="d-flex">
                <b-avatar :src="getImgUrl(blog.user.avatar)"></b-avatar>
                <div class="user">
                    <p>{{blog.user.name + " " + blog.user.last_name}}</p>
                    <p class="grey">{{blog.user.title}}</p>
                </div>
            </div>
        </div>
       <template v-if="blog.title">
           <h3 class="d-flex mt-3 justify-content-center">
               <VueShowdown
                       :markdown="blog.title"
                       flavor="github"
                       :options="{ emoji: true }"
               ></VueShowdown>
           </h3>
       </template>
        <template v-if="blog.content">
            <VueShowdown
                    :markdown="blog.content"
                    flavor="github"
                    :options="{ emoji: true }"
            ></VueShowdown>
        </template>
    </div>
</template>

<script>
    export default {
        name: "IndividualBlog",
        data() {
            return {
                blog: null
            }
        },
        methods: {
            getBlog() {
                window.axios
                    .get("api/blog/get/" + this.$route.params.id +"/")
                    .then(res => {
                        this.blog = res.data.blog;
                    })
                    .catch(rej => {
                        this.makeToast(rej.response.data.error, "danger");
                    });
            },
            getImgUrl(image) {
                let images = process.env.VUE_APP_API + "/static" + image;
                return images;
            },
            makeToast(text, variant) {
                this.$bvToast.toast(text, {
                    autoHideDelay: 5000,
                    variant: variant,
                    title: "Notification"
                });
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
            this.getBlog()
        },
    }
</script>

<style scoped lang="scss">
    .container-blog{
        max-width: 800px;
        margin: auto;
        .banner{
            max-height: 400px;
            background-size: cover;
            background: #49599a;
            width: 100%;
        }
        .info{
            margin-top: 5px;
            display: flex;
            font-weight: bolder;
            align-items: center;
            p{
                margin-bottom: 0;
            }
        }
        .user {
            margin-left: 10px;
            p{
                margin-bottom: 0;
                font-size: small;
            }
        }
    }
    .grey{
        color: #9e9e9e;
    }
</style>