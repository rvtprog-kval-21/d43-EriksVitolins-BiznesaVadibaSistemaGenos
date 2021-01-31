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
            <b-button variant="success" v-if="!sameUser()">Follow</b-button>
            <b-button variant="outline-primary" @click="$router.push('/user/' + user.id + '/settings')" v-if="sameUser()"
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
            <p class="m-2">{{user.about}}</p>
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
        <b-col  v-if="this.currentUser.role == 'admin'" class="col-lg-7 mt-5 col-md-12 p-4 bg-white card-universal">
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
                ><b-button @click="resetPassword(user.email)">Reset password</b-button></b-col
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
      newEmail: ""
    };
  },
  mounted() {
    this.getUser();
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
        .post("/api/admin/user/" + this.user.ID + "/lock")
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
        .post("/api/admin/user/" + this.user.ID + "/unlock")
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
    resetPassword(email) {
      this.errors = [];
      this.alerts = [];
      const vue = this;
      window.axios
        .post("/api/admin/user/reset/password", {"email": email})
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
.mother-load{
  min-height: 1000px;
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
  .border{
    border: 1px black solid;
  }
</style>
