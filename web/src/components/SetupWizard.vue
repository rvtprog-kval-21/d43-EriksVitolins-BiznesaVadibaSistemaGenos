<template>
  <form-wizard @on-complete="initUser()" shape="square" color="#3498db">
    <tab-content title="Start">
      <h4>We are very pleased with you joining our company</h4>
      <h4>But first you need to setup your profile</h4>
    </tab-content>
    <tab-content title="Profile details">
      <div>
        <h5>About:</h5>
        <b-form-textarea
          id="textarea"
          v-model="user.about"
          placeholder="Enter something..."
          rows="3"
          class="w-50 mt-2 mb-4"
          max-rows="6"
        ></b-form-textarea>
        <b-button @click="saveAbout()" variant="outline-success">Save</b-button>
      </div>
      <hr />
      <div>
        <h5>Title:</h5>
        <div class="d-flex mb-4 mt-2">
          <b-form-input
            v-model="user.title"
            class="w-25"
            placeholder="Enter your title"
          ></b-form-input>
        </div>
        <b-button @click="saveTitle()" variant="outline-success">Save</b-button>
      </div>
      <hr />
      <div>
        <h5>Phone Number:</h5>
        <div class="d-flex mb-4 mt-2">
          <b-form-input
            v-model="user.phone_number"
            class="w-25"
            placeholder="Enter your phone number"
          ></b-form-input>
        </div>
        <b-button @click="saveNumber()" variant="outline-success"
          >Save</b-button
        >
      </div>
      <hr />
    </tab-content>
    <tab-content title="Important dates">
      <div>
        <h5>Birthday and Nameday:</h5>
        <div class="d-flex mb-4 mt-2">
          <div>
            <label for="birthday-datepicker">Choose a Birthday</label>
            <b-form-datepicker
              id="birthday-datepicker"
              v-model="user.birthday"
              class="mb-2"
            ></b-form-datepicker>
          </div>
          <div class="ml-4">
            <label for="nameday-datepicker">Choose a Nameday</label>
            <b-form-datepicker
              id="nameday-datepicker"
              v-model="user.name_day"
              class="mb-2"
            ></b-form-datepicker>
          </div>
        </div>
        <b-button @click="saveBirthday()" variant="outline-success"
          >Save</b-button
        >
      </div>
      <hr />
    </tab-content>
    <tab-content title="Background and Avatar">
      <h5>Background:</h5>
      <b-form-file
        v-model="background"
        :state="Boolean(background)"
        class="mt-3 mb-3"
        placeholder="Choose a photo or drop it here..."
        drop-placeholder="Drop photo here..."
        accept="image/*"
      ></b-form-file>
      <b-button @click="saveBackground()" variant="outline-success"
        >Save</b-button
      >
      <hr />
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
        <b-button @click="saveAvatar()" variant="outline-success"
          >Save</b-button
        >
      </div>
      <hr />
    </tab-content>
    <tab-content title="End">
      <h3>Thank you for joining this company</h3>
    </tab-content>
  </form-wizard>
</template>

<script>
import { FormWizard, TabContent } from "vue-form-wizard";
import "vue-form-wizard/dist/vue-form-wizard.min.css";
export default {
  name: "SetupWizard",
  props: ["isInit"],
  components: {
    FormWizard,
    TabContent
  },
  data() {
    return {
      user: {},
      background: null,
      avatar: null
    };
  },
  methods: {
    onComplete() {
      console.log(1);
    },
    getUser: function() {
      this.errors = {};
      const vue = this;
      window.axios
        .get("/api/user/" + this.currentUser.id + "/profile")
        .then(res => {
          this.user = res.data.data;
        })
        .catch(function(rej) {
          vue.errors = { error: rej.response.data.error };
        });
    },
    initUser: function() {
      window.axios.get("/api/account/init/").then(() => {
        this.isInit = true;
        location.reload();
      });
    },
    saveNumber: function() {
      if (
        8 < this.user.phone_number.length ||
        11 < this.user.phone_number.length
      ) {
        this.makeToast("A valid phone number wasn't inputed", "danger");
        return;
      }
      window.axios
        .post("/api/user/settings/new/number", this.user)
        .then(res => {
          this.makeToast(res.data.message, "success");
          this.getUser();
        })
        .catch(function(rej) {
          this.makeToast(rej.response.data.error, "danger");
        });
    },
    saveTitle: function() {
      window.axios
        .post("/api/user/settings/new/title", this.user)
        .then(res => {
          this.makeToast(res.data.message, "success");
          this.getUser();
        })
        .catch(function(rej) {
          this.makeToast(rej.response.data.error, "danger");
        });
    },
    saveAbout: function() {
      window.axios
        .post("/api/user/settings/new/about", this.user)
        .then(res => {
          this.makeToast(res.data.message, "success");
          this.getUser();
        })
        .catch(function(rej) {
          this.makeToast(rej.response.data.error, "danger");
        });
    },
    saveBirthday: function() {
      if (!this.user.name_day.includes("T")) {
        this.user.name_day = this.user.name_day + "T22:34:12.785+02:00";
      } else if (!this.user.birthday.includes("T")) {
        this.user.birthday = this.user.birthday + "T22:34:12.785+02:00";
      }

      window.axios
        .post("/api/user/settings/new/birthday", this.user)
        .then(res => {
          this.makeToast(res.data.message, "success");
          this.getUser();
        })
        .catch(function(rej) {
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
    saveAvatar() {
      let formData = new FormData();
      formData.append("avatar", this.avatar);
      window.axios
        .post(`api/user/settings/new/avatar`, formData, {
          headers: {
            "Content-Type": "multipart/form-data"
          }
        })
        .then(() => {
          this.makeToast("User avatar changed successfully", "success");
          this.getProject();
        })
        .catch(rej => {
          this.makeToast(rej.response.data.error, "danger");
        });
    },
    saveBackground() {
      let formData = new FormData();
      formData.append("background", this.background);
      window.axios
        .post(`api/user/settings/new/background`, formData, {
          headers: {
            "Content-Type": "multipart/form-data"
          }
        })
        .then(() => {
          this.makeToast("User avatar changed successfully", "success");
          this.getProject();
        })
        .catch(rej => {
          this.makeToast(rej.response.data.error, "danger");
        });
    }
  },
  mounted() {
    this.getUser();
  },
  computed: {
    currentUser() {
      return this.$store.getters.currentUser;
    }
  }
};
</script>

<style scoped></style>
