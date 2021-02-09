<template>
  <div class="content d-flex mt-2">
    <div v-if="isInit" class="col-10">
      <div class="top-row d-flex">
        <div class="blog pt-3 mr-0">
          <div class="blog-top pl-3 d-flex flex-column">
            <h5>Blogs:</h5>
          </div>
          <hr />
          <div>
            <div
              class="blog-row d-flex"
              v-for="(iter, index) in this.blogs"
              :key="index"
            >
              <div
                class="ml-2 mr-2 d-flex justify-content-center align-items-center"
              >
                <b-icon
                  v-if="!iter.viewed"
                  font-scale="2"
                  variant="danger"
                  icon="dot"
                ></b-icon>
              </div>
              <div>
                <div
                  @click="$router.push('/blog/' + iter.blog.id + '/view')"
                  class="d-flex align-items-center"
                >
                  <p class="font-weight-bolder">{{ iter.blog.topic }}</p>
                  <b-icon class="mr-2 mb-3 ml-2" icon="dot"></b-icon>
                  <VueShowdown
                    class="mb-0"
                    :markdown="iter.blog.title"
                    flavor="github"
                    :options="{ emoji: true }"
                  ></VueShowdown>
                </div>
                <div>
                  <VueShowdown
                    :markdown="iter.blog.headtext"
                    flavor="github"
                    :options="{ emoji: true }"
                  ></VueShowdown>
                </div>
              </div>
            </div>
            <h4 class="d-flex justify-content-center" v-if="blogs.length < 1">
              No blogs available
            </h4>
          </div>
        </div>
        <div class="pt-3 ml-3 calendar">
          <div class="d-flex">
            <h5 class="ml-3 mb-4">Calendars:</h5>
            <h5 class="ml-2">{{ events.length }}</h5>
          </div>
          <hr />
          <h4 class="d-flex justify-content-center" v-if="events.length < 1">
            No Events available
          </h4>
          <template v-for="(item, index) in events">
            <div :key="index" class="mb-3">
              <b-card :title="item.title" img-alt="Image" img-top>
                <template v-for="(iter, index) in item.members">
                  <div
                    class="d-flex align-items-center"
                    :key="index"
                    v-if="iter.isOwner == true"
                  >
                    <b-avatar
                      class="mr-1"
                      size="2rem"
                      :src="getImgUrl(iter.user.avatar)"
                    ></b-avatar>
                    <p class="mb-0 text-muted">
                      {{ iter.user.name + " " + iter.user.last_name }}
                    </p>
                  </div>
                </template>
                <hr />
                <b-card-text>
                  {{ item.description }}
                </b-card-text>
                <template v-if="item.members.length > 1">
                  <p>Participants:</p>
                  <div class="d-flex">
                    <template v-for="(iter, index) in item.members">
                      <div
                        class="d-flex align-items-center"
                        :key="index"
                        v-if="iter.isOwner == false"
                      >
                        <b-avatar
                          class="mr-1"
                          size="2rem"
                          :src="getImgUrl(iter.user.avatar)"
                        ></b-avatar>
                      </div>
                    </template>
                  </div>
                </template>
                <template #footer>
                  <small class="text-muted"
                    >{{ item.startDate + " ~ " + item.endDate }}}</small
                  >
                </template>
              </b-card>
            </div>
          </template>
        </div>
      </div>
      <div class="top-row d-flex mt-3">
        <div class="random mr-0">
          Test
        </div>
        <div class="follow-field">
          <h5 class="ml-3 mb-4">Announcements from people You follow:</h5>
          <hr />
          <h4 class="d-flex justify-content-center" v-if="annc.length < 1">
            No Announcements available
          </h4>
          <template v-for="(iter, index) in annc">
            <div :key="index" class="m-row">
              <div class="d-flex mt-2 mb-2">
                <div class="w-100">
                  <div class="d-flex align-content-center">
                    <div class="cursor" @click="goToProfile(iter.user.id)">
                      <b-avatar
                        class="avatar ml-2"
                        :src="getImgUrl(iter.user.avatar)"
                        variant="primary"
                        size="2rem"
                        text="EV"
                      ></b-avatar>
                    </div>
                    <p class="mb-0 ml-2 grey">
                      {{ new Date(iter.created_at) }}
                    </p>
                  </div>
                  <p class="ml-2 mt-2 mb-0">{{ iter.message }}</p>
                </div>
              </div>
              <hr class="mt-0" />
            </div>
          </template>
        </div>
      </div>
    </div>
    <div v-if="isInit" class="col-2">
      <div class="users">
        <template v-for="(iter, index) in users">
          <div class="user" :key="index">
            <b-dropdown variant="outline-none" class="item">
              <template class="wow" #button-content>
                <div class="info">
                  <b-avatar
                    size="3rem"
                    badge
                    badge-variant="success"
                    :src="getImgUrl(iter.avatar)"
                  ></b-avatar>
                  <h6>{{ iter.name + " " + iter.last_name }}</h6>
                </div>
              </template>
              <b-dropdown-item :href="`/user/${iter.id}/profile`"
                >Profile</b-dropdown-item
              >
              <b-dropdown-item href="#">Another item</b-dropdown-item>
            </b-dropdown>
          </div>
        </template>
      </div>
    </div>
    <div v-if="!isInit" class="w-100">
      <SetupWizard :is-init="isInit"></SetupWizard>
    </div>
  </div>
</template>

<script>
import SetupWizard from "../components/SetupWizard";
export default {
  components: {
    SetupWizard
  },
  name: "Home",
  data() {
    return {
      onlineUsers: [],
      blogs: [],
      events: [],
      isInit: true,
      annc: [],
      users: []
    };
  },
  methods: {
    getOnlineUsers() {
      window.axios.get("/api/usersonline").then(res => {
        this.onlineUsers = res.data.users;
        this.getUsers();
      });
    },
    getUsers() {
      if (this.onlineUsers) {
        const users = [];
        for (const item in this.onlineUsers) {
          users.push(item);
        }
        console.log(users);
        window.axios
          .post("/api/users/get/multiple", { users: users })
          .then(res => {
            console.log(122);
            this.users = res.data.users;
          });
      }
    },
    getEvents() {
      window.axios.post("api/calendar/get/home/events/").then(res => {
        this.events = res.data.events;
      });
    },
    getImgUrl(avatar) {
      let images = process.env.VUE_APP_API + "/static" + avatar;
      return images;
    },
    goToProfile(id) {
      this.$router.push("/user/" + id + "/profile");
    },
    profile() {
      this.$router.push("/user/" + this.currentUser.id + "/profile");
    },
    getAnnc: function() {
      window.axios.post("/api/user/announcements/gets/follow").then(res => {
        if (res.data.data) {
          this.annc = res.data.data;
        }
      });
    },
    getBlogs() {
      window.axios
        .get("api/blog/home/limited")
        .then(res => {
          this.blogs = res.data.blogs;
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
    checkInit: function() {
      window.axios.get("/api/account/check/").then(res => {
        this.isInit = res.data.isInit;
      });
    }
  },
  computed: {
    currentUser() {
      return this.$store.getters.currentUser;
    }
  },
  created() {
    this.getOnlineUsers();
    this.getBlogs();
    this.getEvents();
    this.getAnnc();
    this.checkInit();
  }
};
</script>

<style scoped lang="scss">
.top-row {
  height: 500px;
  .blog {
    width: 320px;
    background: white;
    height: 100%;
    border-radius: 10px;
    .blog-row {
      cursor: pointer;
      margin-bottom: 20px;
      &:hover {
        background-color: #76d275;
      }
    }
  }
}

.cursor {
  cursor: pointer;
}
.calendar {
  width: 100%;
  overflow: auto;
  border-radius: 10px;
}
.users {
  padding: 30px 30px 30px 0;
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
.random {
  width: 320px;
}
.follow-field {
  width: 100%;
  overflow: auto;
  max-height: 400px;
}
</style>

<style lang="scss">
.users {
  .user {
    .item {
      .btn {
        display: flex;
        justify-content: space-between;
        min-width: 300px;
        align-items: center;
      }
    }
  }
}
.grey {
  color: #b0bec5;
}
.border {
  border: 1px black solid;
}
</style>
