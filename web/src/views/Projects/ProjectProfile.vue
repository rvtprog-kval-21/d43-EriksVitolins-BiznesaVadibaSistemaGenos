<template>
  <div class="body">
    <div class="right">
      <div
        class="header d-flex flex-column align-items-center justify-content-center mt-4 mb-4"
      >
        <template  v-if="project">
          <template>
            <b-avatar size="8rem" :src="getImgUrl(project.avatar)"></b-avatar>
          </template>
          <h3 class="mt-2 mb-2">{{ project.name }}</h3>
          <div v-if="user" class="w-100 p-4">
            <div @click="aboutIsOpened = !aboutIsOpened" class="d-flex about flex-column align-items-center justify-content-center mb-2">
              <h5 class="mb-0">About</h5>
              <b-icon v-if="!aboutIsOpened" icon="arrow-down"></b-icon>
                <b-icon v-if="aboutIsOpened" icon="arrow-up"></b-icon>
            </div>
              <VueShowdown
                      v-if="aboutIsOpened"
                      class="border"
                      :markdown="project.about"
                      flavor="github"
                      :options="{ emoji: true }"
              ></VueShowdown>
          </div>
        </template>
      </div>
    </div>
    <div class="left">
      <div class="users" v-if="project">
        <template v-for="(iter, index) in project.members">
          <div class="user" :key="index">
            <template v-if="iter">
              <b-dropdown variant="outline-none" class="item">
                <template class="wow" #button-content>
                  <div class="info" v-if="iter">
                    <b-avatar
                            size="3rem"
                            :src="getImgUrl(iter.user.avatar)"
                    ></b-avatar>
                    <h6 v-if="iter.user.name">{{ iter.user.name + " " + iter.user.last_name }}</h6>
                  </div>
                </template>
                <b-dropdown-item :href="`/user/${iter.user.id}/profile`"
                >Profile</b-dropdown-item
                >
                <template v-if="!isLoading">
                  <b-dropdown-item @click="makeAdmin(iter.user.id)" v-if="user.is_owner && getXUser(iter.user.id).is_admin == false && getXUser(iter.user.id).is_owner == false && currentUser.id != iter.user.id" href="#">Make Admin</b-dropdown-item>
                  <b-dropdown-item @click="unmakeAdmin(iter.user.id)" v-else-if="user.is_owner && getXUser(iter.user.id).is_admin == true && getXUser(iter.user.id).is_owner == false &&  currentUser.id != iter.user.id" href="#">Take Away Admin</b-dropdown-item>
                  <b-dropdown-item @click="kickUser(iter.user.id)" v-if="(user.is_owner || user.is_admin) && getXUser(iter.user.id).is_admin == false && getXUser(iter.user.id).is_owner == false &&  currentUser.id != iter.user.id" href="#">Kick</b-dropdown-item>
                  <b-dropdown-item @click="leaveProject()" v-if="currentUser.id == iter.user.id" href="#">Leave</b-dropdown-item>
                </template>
              </b-dropdown>
            </template>
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
      user: {},
      aboutIsOpened: false,
      isLoading: false
    };
  },
  computed: {
    currentUser() {
      return this.$store.getters.currentUser;
    },
    isTheSameUser(id) {
      if (this.currentUser) {
        console.log(this.currentUser.id)
        return this.currentUser.id == id
      } else {
        return false
      }
    }
  },
  methods: {
    getUser() {
      const users = this.project.members;
      this.user = users.find(e => e.user.id === this.currentUser.id);
    },
    getXUser(id) {
      const users = this.project.members;
      const user = users.find(e => e.user.id === id);
      return user
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
    makeAdmin(id) {
      window.axios
              .get(`api/projects/add/${this.$route.params.id}/admin/?id=${id}`)
              .then(() => {
                this.getProject()
                this.makeToast(`User was made an admin`, "success");
              })
              .catch((rej) => {
                this.makeToast(rej.response.data.error, "danger");
              });
    },
    unmakeAdmin(id) {
      window.axios
              .get(`api/projects/remove/${this.$route.params.id}/admin/?id=${id}`)
              .then(() => {
                this.getProject()
                this.makeToast(`Admin was taken away from the user`, "success");
              })
              .catch((rej) => {
                this.makeToast(rej.response.data.error, "danger");
              });
    },
    kickUser(id) {
      window.axios
              .get(`api/projects/kick/${this.$route.params.id}/member/?id=${id}`)
              .then(() => {
                this.getProject()
                this.makeToast(`User was kicked from the project`, "success");
              })
              .catch((rej) => {
                this.makeToast(rej.response.data.error, "danger");
              });
    },
    leaveProject() {
      window.axios
              .get(`api/projects/leave/${this.$route.params.id}/member/`)
              .then(() => {
                this.$router.push("/projects");
              })
              .catch((rej) => {
                this.makeToast(rej.response.data.error, "danger");
              });
    },
    makeToast(text, variant) {
      this.$bvToast.toast(text, {
        autoHideDelay: 5000,
        variant: variant,
        title: "Notification"
      });
    },
    getImgUrl(image) {
      let images = process.env.VUE_APP_API + "/static" + image;
      return images;
    }
  },
  async created() {
    await this.getProject();
  }
};
</script>

<style scoped lang="scss">
.body {
  height: 100%;
  display: flex;
  .right {
    width: 70%;
    overflow: auto;
      .about{
          cursor: pointer;
          &:hover{
              color: #17a2b8;
          }
      }
  }
  .left {
    width: 30%;
    overflow: auto;
    .users {
      padding: 30px;
      .user {
        .item {
          .info {
            display: flex;
            margin: 0 0 10px 0;
            cursor: pointer;
            align-items: center;
            h6 {
              margin-left: 10px;
            }
          }
        }
        &:hover {
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
      .btn {
        display: flex;
        justify-content: space-between;
        min-width: 400px;
        align-items: center;
      }
    }
  }
}
</style>
