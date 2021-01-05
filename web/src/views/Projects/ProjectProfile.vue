<template>
  <div class="body">
    <div class="right">
      <b-button
        v-if="this.user.is_admin || this.user.is_owner"
        class="mt-3 ml-3"
        variant="outline-primary"
        @click="goToSettings()"
        >Settings</b-button
      >
      <div
        class="header d-flex flex-column align-items-center justify-content-center mt-4 mb-4"
      >
        <template v-if="project">
          <template>
            <b-avatar size="8rem" :src="getImgUrl(project.avatar)"></b-avatar>
          </template>
          <h3 class="mt-2 mb-2">{{ project.name }}</h3>
          <div v-if="user" class="w-100 p-4">
            <div
              @click="aboutIsOpened = !aboutIsOpened"
              class="d-flex about flex-column align-items-center justify-content-center mb-2"
            >
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
          <div class=" p-4 w-100">
            <h5 class="d-flex">Announcement Board:</h5>
            <div class="w-100">
              <div class="w-100 bg-white a-board">
                <div class="loadMore mt-2 mr-2 d-flex justify-content-end">
                  <b-button @click="getAnn()" v-if="!loadingNew" pill  variant="outline-info">
                    <b-icon icon="arrow-down-circle" font-scale="2"></b-icon>
                  </b-button>
                  <b-icon v-else icon="arrow-clockwise" animation="spin" font-scale="2"></b-icon>
                </div>
                <div class="d-flex flex-column">
                  <div class="ann-row ml-3 mb-3" v-for="(iter, index) in annc" :key="index">
                      <div class="d-flex align-items-end">
                        <b-avatar size="3rem" :src="getImgUrl(iter.author.avatar)"></b-avatar>
                        <h6> {{
                          iter.author.name +
                          " " +
                          iter.author.last_name
                          }}</h6>
                        <p>{{getDate(iter.published)}}</p>
                      </div>
                      <div class="bubble mt-3">
                        <VueShowdown
                                :markdown="iter.content"
                                flavor="github"
                                :options="{ emoji: true }"
                        ></VueShowdown>
                      </div>
                  </div>
                </div>
              </div>
              <div
                v-if="this.user.is_admin || this.user.is_owner"
                class="mt-3 d-flex form-row"
              >
                <b-form-textarea
                  class="w-50"
                  id="textarea"
                  v-model="message"
                  placeholder="Message"
                  rows="3"
                  max-rows="6"
                ></b-form-textarea>
                <vSelect
                  multiple
                  class="w-25"
                  v-model="messageTags"
                  label="name"
                  :options="project.tags"
                />
                <b-button @click="createAnn()" variant="success" class="w-25"
                  >Announce</b-button
                >
              </div>
            </div>
          </div>
        </template>
      </div>
    </div>
    <div class="left">
      <div class="users" v-if="project">
        <template v-for="(iter, index) in project.members">
          <div class="user" :key="index">
            <template v-if="iter">
              <b-dropdown variant="outline-none" class="item item-drowpdown">
                <template class="wow" #button-content>
                  <div class="info" v-if="iter">
                    <b-avatar
                      size="3rem"
                      :src="getImgUrl(iter.user.avatar)"
                    ></b-avatar>
                    <h6 v-if="iter.user.name">
                      {{
                        iter.user.name +
                          " " +
                          iter.user.last_name +
                          " ( " +
                          iter.tag.name +
                          " )"
                      }}
                    </h6>
                  </div>
                </template>
                <b-dropdown-item :href="`/user/${iter.user.id}/profile`"
                  >Profile</b-dropdown-item
                >
                <template v-if="!isLoading">
                  <vSelect
                    v-if="user.is_owner || user.is_admin"
                    class=" ml-auto mr-auto w-75"
                    v-model="iter.tag"
                    label="name"
                    :options="project.tags"
                    @input="updateTag(iter.user.id, iter.tag)"
                  />
                  <b-dropdown-item
                    @click="makeAdmin(iter.user.id)"
                    v-if="
                      user.is_owner &&
                        getXUser(iter.user.id).is_admin == false &&
                        getXUser(iter.user.id).is_owner == false &&
                        currentUser.id != iter.user.id
                    "
                    href="#"
                    >Make Admin</b-dropdown-item
                  >
                  <b-dropdown-item
                    @click="unmakeAdmin(iter.user.id)"
                    v-else-if="
                      user.is_owner &&
                        getXUser(iter.user.id).is_admin == true &&
                        getXUser(iter.user.id).is_owner == false &&
                        currentUser.id != iter.user.id
                    "
                    href="#"
                    >Take Away Admin</b-dropdown-item
                  >
                  <b-dropdown-item
                    @click="kickUser(iter.user.id)"
                    v-if="
                      (user.is_owner || user.is_admin) &&
                        getXUser(iter.user.id).is_admin == false &&
                        getXUser(iter.user.id).is_owner == false &&
                        currentUser.id != iter.user.id
                    "
                    href="#"
                    >Kick</b-dropdown-item
                  >
                  <b-dropdown-item
                    @click="leaveProject()"
                    v-if="currentUser.id == iter.user.id"
                    href="#"
                    >Leave</b-dropdown-item
                  >
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
import vSelect from "vue-select";
import "vue-select/dist/vue-select.css";

export default {
  name: "ProjectProfile",
  components: { vSelect },
  data() {
    return {
      project: null,
      test: [],
      user: {},
      aboutIsOpened: false,
      isLoading: false,
      message: "",
      messageTags: null,
      annc: [],
      loadingNew: false,
      currentPage: 0
    };
  },
  computed: {
    currentUser() {
      return this.$store.getters.currentUser;
    }
  },
  methods: {
    getDate(date) {
      date = new Date(date);
      const day = date.getDate();
      const year = date.getFullYear();
      const month = date.toLocaleString("default", {month: "long"});
      return day + " " + month + " " + year;
    },
    getUser() {
      const users = this.project.members;
      this.user = users.find(e => e.user.id === this.currentUser.id);
    },
    getXUser(id) {
      const users = this.project.members;
      const user = users.find(e => e.user.id === id);
      return user;
    },
    getProject() {
      this.isLoading = true;
      window.axios
        .get(`api/projects/get/${this.$route.params.id}/item/`)
        .then(res => {
          this.project = res.data.project;
          this.getUser();
          this.isLoading = false;
        })
        .catch(() => {
          this.$router.push("/projects");
        });
    },
    goToSettings() {
      this.$router.push(`/projects/${this.$route.params.id}/settings`);
    },
    makeAdmin(id) {
      window.axios
        .get(`api/projects/add/${this.$route.params.id}/admin/?id=${id}`)
        .then(() => {
          this.getProject();
          this.makeToast(`User was made an admin`, "success");
        })
        .catch(rej => {
          this.makeToast(rej.response.data.error, "danger");
        });
    },
    getAnn() {
      this.loadingNew = true
      window.axios
        .get(
          `api/projects/see/${this.$route.params.id}/current/announcement?currentPage=${this.currentPage}`
        )
        .then(res => {
          console.log(res.data.annc)
          const arr = res.data.annc
          this.annc = this.annc.concat(arr);
          this.currentPage +=1;
          if (res.data.annc.length < 1) {
            this.makeToast(`No more announcements available`, "warning");
          }
        })
        .catch(rej => {
          this.makeToast(rej.response.data.error, "danger");
        });
      this.loadingNew = false
    },
    createAnn() {
      window.axios
        .post(`api/projects/create/${this.$route.params.id}/new/announcement`, {
          content: this.message,
          tags: this.messageTags
        })
        .then(() => {
          this.getProject();
          this.message = null;
          this.messageTags = null;
          this.currentPage = 0;
          this.annc = [];
          this.getAnn();
          this.makeToast(`Admin was taken away from the user`, "success");
        })
        .catch(rej => {
          this.makeToast(rej.response.data.error, "danger");
        });
    },
    updateTag(id, tag) {
      window.axios.post(
        `api/projects/update/${this.$route.params.id}/tags/?id=${id}`,
        tag
      );
    },
    unmakeAdmin(id) {
      window.axios
        .get(`api/projects/remove/${this.$route.params.id}/admin/?id=${id}`)
        .then(() => {
          this.getProject();
          this.makeToast(`Admin was taken away from the user`, "success");
        })
        .catch(rej => {
          this.makeToast(rej.response.data.error, "danger");
        });
    },
    kickUser(id) {
      window.axios
        .get(`api/projects/kick/${this.$route.params.id}/member/?id=${id}`)
        .then(() => {
          this.getProject();
          this.makeToast(`User was kicked from the project`, "success");
        })
        .catch(rej => {
          this.makeToast(rej.response.data.error, "danger");
        });
    },
    leaveProject() {
      window.axios
        .get(`api/projects/leave/${this.$route.params.id}/member/`)
        .then(() => {
          this.$router.push("/projects");
        })
        .catch(rej => {
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
    this.getAnn()
  }
};
</script>

<style scoped lang="scss">
  .loadMore{
    position: -webkit-sticky;
    position: sticky;
    top: 0;
    padding-top: 10px;
  }

  .ann-row{
    p{
      color: #9e9e9e;
      font-size: small;
      margin-left:10px ;
    }

    h6{
      margin-bottom: auto;
      margin-top: auto;
      margin-left: 10px;
    }

    .bubble{
      position: relative;
      max-width: 30em;

      background-color: #fff;
      padding: 0.5em 1em;
      font-size: 1.25em;
      border-radius: 1rem;
      box-shadow:	0 0.125rem 0.5rem rgba(0, 0, 0, .3), 0 0.0625rem 0.125rem rgba(0, 0, 0, .2);
    }
    .bubble:after{
      content: '';
      position: absolute;
      width: 0;
      height: 0;
      bottom: 100%;
      left: 0.6em;
      border: .75rem solid transparent;
      border-top: none;

      border-bottom-color: #fff;
      filter: drop-shadow(0 -0.0625rem 0.0625rem rgba(0, 0, 0, .1));
    }
  }
.body {
  height: 100%;
  display: flex;
  .right {
    width: 70%;
    overflow: auto;
    .about {
      cursor: pointer;
      &:hover {
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
  .a-board {
    min-height: 200px;
    max-height: 400px;
    overflow: auto;
  }
}
</style>

<style lang="scss">
.dropdown-menu {
  width: 250px;
}
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
