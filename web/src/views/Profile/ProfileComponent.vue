<template>
  <div class="mother-load">
    <template v-if="user.email">
      <div class="header d-flex align-items-center">
        <div class="container-fluid">
          <span class="mask bg-gradient opacity-8"></span>
          <div class="d-flex align-items-center container-fluid">
            <b-row class="w-75">
              <b-col cols="8">
                <h1 class="mb-4">
                  Welcome to
                  <span class="title">{{
                    user.name + " " + user.last_name
                  }}</span>
                  profile! :)
                </h1>
                <p class="mb-5 mt-4">{{ user.about }}</p>
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
            v-bind:style="{ backgroundImage: 'url(' + geBackgroundUrl() + ')' }"
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
            <b-button variant="success" v-if="!sameUser()">Message</b-button>
            <b-button
              variant="success"
              @click="followUser()"
              v-if="!sameUser() && !isBeingFollowed"
              >Follow</b-button
            >
            <b-button
              variant="warning"
              @click="unfollowUser()"
              v-if="!sameUser() && isBeingFollowed"
              >Unfollow</b-button
            >
            <b-button
              variant="outline-primary"
              @click="$router.push('/user/' + user.id + '/settings')"
              v-if="sameUser()"
              >Settings</b-button
            >
          </div>
          <div class="d-flex justify-content-center mt-3">
            <h4>{{ user.name + " " + user.last_name }}</h4>
          </div>
          <div class="d-flex grey justify-content-center mt-1">
            <h6>{{ user.title }}</h6>
          </div>
          <div class="d-flex grey justify-content-center mt-1">
            <h5>{{ user.email }} | {{ user.phone_number }}</h5>
          </div>
          <h5 class="pt-3 pl-4">About:</h5>
          <div class="border ml-4 mr-4">
            <p class="m-2">{{ user.about }}</p>
          </div>
          <div class="pt-2 pl-4 pr-4 pb-4">
            <div class="d-flex justify-content-center grey">
              <h5>Projects</h5>
            </div>
            <template v-for="(iter, index) in projects">
              <a :href="'/projects/' + iter.id + '/see'" :key="index">
                <b-avatar
                  variant="info"
                  size="4rem"
                  :src="getLogo(iter.avatar)"
                  class="name mr-3"
                ></b-avatar>
              </a>
            </template>
          </div>
          <div class="ml-4 mr-4 grey d-flex justify-content-between">
            <h6 class="">
              <span class="mr-2">Birthday:</span> {{ getDate(user.birthday) }}
            </h6>
            <h6 class="">
              <span class="mr-2">NameDay:</span> {{ getDate(user.name_day) }}
            </h6>
          </div>
        </b-col>
        <b-col class="col-lg-7 mt-5 col-md-12 p-4 bg-white card-universal">
          <div class="annc-top border">
            <h4 class="d-flex justify-content-center" v-if="annc.length < 1">
              No Announcements available
            </h4>
            <template v-for="(iter, index) in annc">
              <div :key="index" class="m-row">
                <div class="d-flex mt-2 mb-2">
                  <div class="w-75">
                    <div class="d-flex align-content-center">
                      <b-avatar
                        class="avatar ml-2"
                        :src="getImgUrl()"
                        variant="primary"
                        size="2rem"
                        text="EV"
                      ></b-avatar>
                      <p class="mb-0 ml-2 grey">
                        {{ new Date(iter.created_at) }}
                      </p>
                    </div>
                    <p class="ml-2 mt-2 mb-0">{{ iter.message }}</p>
                  </div>
                  <div
                    class="d-flex justify-content-center align-content-center"
                  >
                    <b-button
                      v-if="sameUser()"
                      @click="deleteAnnouncement(iter.id)"
                      variant="outline-danger"
                      >Delete</b-button
                    >
                  </div>
                </div>
                <hr class="mt-0" />
              </div>
            </template>
          </div>
          <div class="d-flex justify-content-between mt-4" v-if="sameUser()">
            <div class="w-65">
              <b-form-input
                v-model="message"
                placeholder="Enter message"
              ></b-form-input>
            </div>
            <div class="w-25">
              <b-button
                class="w-100"
                @click="newAnnouncement()"
                variant="outline-success"
                >Announce</b-button
              >
            </div>
          </div>
        </b-col>
        <b-col
          v-if="this.currentUser.role == 'admin'"
          class="col-lg-7 mt-5 col-md-12 p-4 bg-white card-universal"
        >
          <div class="d-flex justify-content-center">
            <b-alert v-if="this.errors.error" variant="danger" show="">{{
              this.errors.error
            }}</b-alert>
            <b-alert v-if="this.alerts.message" variant="success" show="">{{
              this.alerts.message
            }}</b-alert>
          </div>
          <template v-if="this.currentUser.role == 'admin'">
            <h6 class="category-title">Admin section:</h6>
            <b-row>
              <b-col
                v-if="user.created_at != '0001-01-01T00:00:00Z'"
                cols="6"
                class="pl-5 pr-5 pt-5"
                >Created on: {{ new Date(user.created_at) }}</b-col
              >
              <b-col
                v-if="user.updated_at != '0001-01-01T00:00:00Z'"
                cols="6"
                class="pl-5 pr-5 pt-5"
                >Updated on: {{ new Date(user.updated_at) }}</b-col
              >
              <b-col cols="6" class="pl-5 pr-5 pt-5"
                ><b-button @click="resetPassword(user.email)"
                  >Reset password</b-button
                ></b-col
              >
              <b-col cols="6" class="pl-5 pr-5 pt-5">
                <b-button
                  v-if="user.deleted_at == '0001-01-01T00:00:00Z'"
                  variant="danger"
                  @click="lock()"
                  >Lock Account</b-button
                >
                <b-button v-else variant="danger" @click="unlock()"
                  >Unlock Account</b-button
                >
              </b-col>
              <b-col cols="6" class="pl-5 pr-5 pt-5">
                <p>New Email:</p>
                <b-form-input v-model="newEmail" type="email"></b-form-input>
                <b-button
                  @click="setNewEmail()"
                  variant="outline-primary"
                  class="mt-3"
                  >Save</b-button
                >
              </b-col>
            </b-row>
          </template>
        </b-col>
      </div>
    </template>
  </div>
</template>

<script>
export default {
  name: "ProfileComponent",
  computed: {
    currentUser() {
      return this.$store.getters.currentUser;
    }
  },
  data() {
    return {
      user: {},
      showAdminSettings: false,
      errors: {},
      alerts: {},
      accountLocked: false,
      tag: "",
      projects: {},
      isBeingFollowed: false,
      message: "",
      annc: [],
      newEmail: ""
    };
  },
  mounted() {
    this.getUser();
    this.getAnnc();
    this.isFollowingUser();
  },
  methods: {
    getLogo(item) {
      let images = process.env.VUE_APP_API + "/static" + item;
      return images;
    },
    getDate(date) {
      date = new Date(date);
      const day = date.getDate();
      const month = date.toLocaleString("default", { month: "long" });
      return day + " " + month;
    },
    lock: function() {
      this.errors = [];
      this.alerts = [];
      const vue = this;
      window.axios
        .post("/api/admin/user/settings/" + this.$route.params.id + "/lock")
        .then(response => {
          this.alerts = { message: response.data.data };
          this.getUser();
        })
        .catch(function(errors) {
          vue.errors = { error: errors.response.data.error };
        });
    },
    unlock: function() {
      this.errors = [];
      this.alerts = [];
      const vue = this;
      window.axios
        .post("/api/admin/user/settings/" + this.$route.params.id+ "/unlock")
        .then(response => {
          this.alerts = { message: response.data.data };
          this.getUser();
        })
        .catch(function(errors) {
          vue.errors = { error: errors.response.data.error };
        });
    },
    getUser: function() {
      this.errors = {};
      const vue = this;
      window.axios
        .get("/api/user/" + this.$route.params.id + "/profile")
        .then(res => {
          this.user = res.data.data;
          this.projects = res.data.projects;
        })
        .catch(function(rej) {
          vue.errors = { error: rej.response.data.error };
        });
    },
    getImgUrl() {
      let images = process.env.VUE_APP_API + "/static" + this.user.avatar;
      return images;
    },
    sameUser() {
      return this.$route.params.id == this.currentUser.id;
    },
    geBackgroundUrl() {
      let images = process.env.VUE_APP_API + "/static" + this.user.background;
      return images;
    },
    setNewEmail() {
      this.errors = [];
      this.alerts = [];
      const vue = this;
      window.axios
        .post("/api/admin/user/" + this.$route.params.id + "/newEmail", {
          id: this.$route.params.id,
          email: this.newEmail
        })
        .then(response => {
          this.alerts = { message: response.data.message };
          this.getUser();
        })
        .catch(function(errors) {
          vue.errors = { error: errors.response.data.error };
        });
    },
    followUser() {
      this.errors = [];
      this.alerts = [];
      const vue = this;
      window.axios
        .post("/api/user/follower/start", {
          following_id: this.$route.params.id
        })
        .then(response => {
          this.isBeingFollowed = true;
          this.alerts = { message: response.data.data };
        })
        .catch(function(errors) {
          vue.errors = { error: errors.response.data.error };
        });
    },
    isFollowingUser() {
      if (this.sameUser()) {
        return;
      }
      this.errors = [];
      const vue = this;
      window.axios
        .post("/api/user/follower/check", {
          following_id: this.$route.params.id
        })
        .then(response => {
          this.isBeingFollowed = response.data.isFollowing;
        })
        .catch(function(errors) {
          vue.errors = { error: errors.response.data.error };
        });
    },
    unfollowUser() {
      this.errors = [];
      this.alerts = [];
      const vue = this;
      window.axios
        .post("/api/user/follower/delete", {
          following_id: this.$route.params.id
        })
        .then(response => {
          this.isBeingFollowed = false;
          this.alerts = { message: response.data.message };
        })
        .catch(function(errors) {
          vue.errors = { error: errors.response.data.error };
        });
    },
    newAnnouncement() {
      this.errors = [];
      this.alerts = [];
      const vue = this;
      window.axios
        .post("/api/user/announcements/new/announcements", {
          message: this.message
        })
        .then(() => {
          this.message = "";
          this.getAnnc();
        })
        .catch(function(errors) {
          vue.errors = { error: errors.response.data.error };
        });
    },
    deleteAnnouncement(id) {
      this.errors = [];
      this.alerts = [];
      const vue = this;
      window.axios
        .post("/api/user/announcements/delete/announcement", {
          id: `${id}`
        })
        .then(() => {
          this.getAnnc();
        })
        .catch(function(errors) {
          vue.errors = { error: errors.response.data.error };
        });
    },
    getAnnc: function() {
      this.errors = {};
      const vue = this;
      window.axios
        .post("/api/user/announcements/get/user", { id: this.$route.params.id })
        .then(res => {
          this.annc = res.data.data;
        })
        .catch(function(rej) {
          vue.errors = { error: rej.response.data.error };
        });
    },
    resetPassword(email) {
      this.errors = [];
      this.alerts = [];
      const vue = this;
      window.axios
        .post("/api/admin/user/reset/password", { email: email })
        .then(response => {
          this.alerts = { message: response.data.message };
          this.getUser();
        })
        .catch(function(errors) {
          vue.errors = { error: errors.response.data.error };
        });
    }
  }
};
</script>

<style scoped lang="scss">
.annc-top {
  overflow: auto;
  height: 90%;
  max-height: 600px;
}
.mother-load {
  min-height: 1000px;
}

.w-65 {
  width: 65% !important;
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
    min-height: 700px;
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
.border {
  border: 1px black solid;
}

.grey {
  color: #b0bec5;
}
</style>
