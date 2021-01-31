<template>
  <b-container>
    <div class="d-flex justify-content-between mt-3">
      <h4>Settings</h4>
      <b-button variant="outline-primary" @click="$router.push('/user/' + user.id + '/profile')">Profile</b-button>
    </div>
    <hr>
    <div class="d-flex flex-column">
      <h5>Background:</h5>
      <b-form-file
              v-model="background"
              :state="Boolean(background)"
              class="mt-3 mb-3"
              placeholder="Choose a photo or drop it here..."
              drop-placeholder="Drop photo here..."
              accept="image/*"
      ></b-form-file>
      <b-button variant="outline-success">Save</b-button>
    </div>
    <hr>
    <div class="d-flex flex-column">
      <h5>Avatar:</h5>
      <b-form-file
              v-model="avatar"
              :state="Boolean(avatar)"
              class="mt-3 mb-3"
              placeholder="Choose a photo or drop it here..."
              drop-placeholder="Drop photo here..."
              accept="image/*"
      ></b-form-file>
      <b-button variant="outline-success">Save</b-button>
    </div>
    <hr>
    <div>
      <h5>Name:</h5>
      <div class="d-flex mb-4 mt-2">
        <b-form-input v-model="user.name" class="w-25" placeholder="Enter your name"></b-form-input>
        <b-form-input v-model="user.last_name" class="w-25 ml-5" placeholder="Enter your lastname"></b-form-input>
      </div>
      <b-button @click="saveName()" variant="outline-success">Save</b-button>
    </div>
    <hr>
    <div>
      <h5>Birthday and Nameday:</h5>
      <b-button variant="outline-success">Save</b-button>
    </div>
    <hr>
    <div>
      <h5>Password:</h5>
      <b-button variant="outline-success">Save</b-button>
    </div>
    <hr>
    <div>
      <h5>About:</h5>
      <b-button variant="outline-success">Save</b-button>
    </div>
    <hr>
    <div>
      <h5>Title:</h5>
      <div class="d-flex mb-4 mt-2">
        <b-form-input v-model="user.title" class="w-25" placeholder="Enter your title"></b-form-input>
      </div>
      <b-button @click="saveTitle()" variant="outline-success">Save</b-button>
    </div>
    <hr>
    <div>
      <h5>Phone Number:</h5>
      <div class="d-flex mb-4 mt-2">
        <b-form-input v-model="user.phone_number" class="w-25" placeholder="Enter your phone number"></b-form-input>
      </div>
      <b-button @click="saveNumber()" variant="outline-success">Save</b-button>
    </div>
    <hr>
  </b-container>
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
      newEmail: "",
      avatar: null,
      background: null,
    };
  },
  created() {
    this.getUser();
    this.sameUser();
  },
  methods: {
    getLogo(item) {
      let images = process.env.VUE_APP_API + "/static" + item;
      return images;
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
    saveName: function() {
      window.axios
              .post("/api/user/settings/new/name", this.user)
              .then(res => {
                this.makeToast(res.data.message, "success")
                this.getUser()
              })
              .catch(function(rej) {
                this.makeToast(rej.response.data.error, "danger")
              });
    },
    saveNumber: function() {
      if (8 < this.user.phone_number.length || 11 < this.user.phone_number.length) {
        this.makeToast("A valid phone number wasn't inputed", "danger")
        return
      }
      window.axios
              .post("/api/user/settings/new/number", this.user)
              .then(res => {
                this.makeToast(res.data.message, "success")
                this.getUser()
              })
              .catch(function(rej) {
                this.makeToast(rej.response.data.error, "danger")
              });
    },
    saveTitle: function() {
      window.axios
              .post("/api/user/settings/new/title", this.user)
              .then(res => {
                this.makeToast(res.data.message, "success")
                this.getUser()
              })
              .catch(function(rej) {
                this.makeToast(rej.response.data.error, "danger")
              });
    },
    makeToast(text, variant) {
      this.$bvToast.toast(text, {
        autoHideDelay: 5000,
        variant: variant,
        title: "Notification"
      });
    },
    getImgUrl() {
      let images = process.env.VUE_APP_API + "/static" + this.user.avatar;
      return images;
    },
    sameUser() {
      if (this.$route.params.id != this.currentUser.id) {
        this.$router.push("user/" + this.$route.params.id + "/settings")
      }
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
  }
};
</script>

<style scoped lang="scss">
  .grey {
  color: #b0bec5;
  }
  .border{
    border: 1px black solid;
  }
</style>
