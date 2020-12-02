<template>
    <div class="body">
       <div class="right">
           <div class="header d-flex flex-column align-items-center justify-content-center mt-4 mb-4">
               <b-avatar size="8rem" :src="getImgUrl(project.avatar)"></b-avatar>
               <h3 class="mt-2 mb-2">{{project.name}}</h3>
           </div>
       </div>
        <div class="left">Test</div>
    </div>
</template>

<script>
    export default {
        name: "ProjectProfile",
        data() {
            return {
                project: null,
            };
        },
        methods: {
            getProject() {
                window.axios
                    .get(`api/projects/get/${this.$route.params.id}/item/`)
                    .then(res => {
                        this.project = res.data.project;
                    })
                    .catch(() => {
                       this.$router.push("/projects")
                    });
            },
            getImgUrl(image) {
                let images = process.env.VUE_APP_API + "/static" + image;
                return images;
            },
        },
        created() {
            this.getProject()
        }
    }
</script>

<style scoped lang="scss">
.body{
    height: 100%;
    display: flex;
    .right{
        width: 70%;
    };
    .left{
        width: 30%;
    }
}
</style>