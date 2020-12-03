<template>
    <div class="body">
       <div class="right">
           <div class="header d-flex flex-column align-items-center justify-content-center mt-4 mb-4">
               <b-avatar size="8rem" :src="getImgUrl(project.avatar)"></b-avatar>
               <h3 class="mt-2 mb-2">{{project.name}}</h3>
           </div>
       </div>
        <div class="left">
            <div class="users">
                <template v-for="(iter, index) in project.members">
                    <div class="user" :key="index">
                        <b-dropdown variant="outline-none" class="item">
                            <template class="wow" #button-content>
                                <div class="info">
                                    <b-avatar size="3rem" :src="getImgUrl(iter.user.avatar)"></b-avatar>
                                    <h6>{{iter.user.name + " " + iter.user.last_name}}</h6>
                                </div>
                            </template>
                            <b-dropdown-item :href="`/user/${iter.user.id}/profile`">Profile</b-dropdown-item>
                            <b-dropdown-item href="#">Another item</b-dropdown-item>
                        </b-dropdown>
                    </div>
                </template>
            </div>
        </div>
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
        overflow: auto;
    };
    .left{
        width: 30%;
        overflow: auto;
        .users{
            padding: 30px;
            .user{
              .item{
                  .info{
                      display: flex;
                      margin: 0 0 10px 0;
                      cursor: pointer;
                      align-items: center;
                      h6{
                          margin-left: 10px;
                      }
                  }
              }
                &:hover{
                    border: solid 1px cornflowerblue;
                }
            }
        }
    }
}
</style>

<style lang="scss">
    .users {
        .user {
            .item {
                .btn{
                    display: flex;
                    justify-content: space-between;
                    min-width: 400px;
                    align-items: center;
                }

            }
        }
    }

</style>