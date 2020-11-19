<template>
  <div>
    <b-alert
      variant="danger"
      class="text-center mb-0"
      v-if="errors.error"
      show
      >{{ errors.error }}</b-alert
    >
    <b-alert
      variant="success"
      class="text-center mb-0"
      v-if="alerts.alert"
      show
      >{{ alerts.alert }}</b-alert
    >
    <template v-if="tag.name">
      <div class="header d-flex align-items-center">
        <div class="container-fluid">
          <span class="mask bg-gradient opacity-8"></span>
          <div class="d-flex align-items-center container-fluid">
            <b-row class="w-75">
              <b-col cols="8">
                <h1 class="mb-4">
                  Welcome to
                  <span class="title">{{ tag.name }}</span>
                  profile! :)
                </h1>
                <p class="mb-5 mt-4">{{ tag.about }}</p>
              </b-col>
            </b-row>
          </div>
        </div>
      </div>
      <div
        class="body flex-wrap container-fluid d-flex justify-content-between"
      >
        <div class="bg-white p-0 profile col-md-12 col-lg-4 mt-5">
          <div
            class="bg-dark h-30 d-flex align-items-end background justify-content-center"
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
            <b-button @click="joinTag()" variant="success" v-if="!isMember"
              >Join</b-button
            >
            <b-button
              @click="[(settingsAreOpended = true), (userListOpened = false)]"
              variant="outline-primary"
              v-if="isMember && !settingsAreOpended"
              >Settings</b-button
            >
            <b-button
              @click="[(settingsAreOpended = false), (userListOpened = true)]"
              v-if="isMember && !userListOpened"
              variant="outline-primary"
              >Users</b-button
            >
            <b-button
              @click="[(settingsAreOpended = false), (userListOpened = false)]"
              v-if="isMember && (settingsAreOpended || userListOpened)"
              variant="outline-primary"
              >Feed</b-button
            >
          </div>
          <div class="d-flex justify-content-center mt-3">
            <h4>{{ tag.name }}</h4>
          </div>
          <div class="d-flex grey justify-content-center mt-1">
            <h6>{{ tag.about }}</h6>
          </div>
          <div v-if="isMember" class="d-flex grey justify-content-center mt-1">
            <h6>User count: {{ getMemberCount() }}</h6>
          </div>
        </div>
        <div
          class="col-lg-7 mt-5 col-md-12 p-4 bg-white card-universal"
          v-if="isMember"
        >
          <template v-if="settingsAreOpended">
            <div class="d-flex justify-content-between">
              <h4>Settings:</h4>
              <b-button
                @click="[(settingsAreOpended = false), (userListOpened = true)]"
                variant="outline-primary"
                >UserList</b-button
              >
              <b-button
                @click="
                  [(settingsAreOpended = false), (userListOpened = false)]
                "
                variant="outline-primary"
                >Feed</b-button
              >
            </div>
            <div class="d-flex mt-5">
              <div class="w-50">
                <h5>Invite User:</h5>
              </div>
              <div class="w-50">
                <h5>Audit:</h5>
                <p class="mt-3 mb-3">See the audit of the tag</p>
                <a :href="'/tags/' + this.$route.params.id + '/tag/timeline'">
                  <b-button variant="outline-info">Audit</b-button>
                </a>
              </div>
            </div>
            <div class="d-flex mt-5" v-if="isAdmin">
              <div class="w-50">
                <h5>Edit Name:</h5>
                <b-form-input
                  class="mt-3 mb-3 w-75"
                  v-model="newName"
                  placeholder="Enter your new name"
                ></b-form-input>
                <b-button @click="setNewName()" variant="success"
                  >Save</b-button
                >
              </div>
              <div class="w-50">
                <h5>Edit About:</h5>
                <b-form-textarea
                  class="mt-3 mb-3"
                  id="textarea-auto-height"
                  placeholder="New About Text"
                  rows="3"
                  v-model="newAbout"
                  max-rows="4"
                ></b-form-textarea>
                <b-button @click="setNewAbout()" variant="success"
                  >Save</b-button
                >
              </div>
            </div>
            <div class="d-flex mt-5" v-if="isAdmin">
              <template v-if="!tag.is_public">
                <div class="w-50">
                  <h5>Make Public:</h5>
                  <p class="mt-3 mb-3">This will let anyone join the Tag</p>
                  <b-button @click="makePublic" variant="warning"
                    >Change</b-button
                  >
                </div>
              </template>
              <template v-else>
                <div class="w-50">
                  <h5>Make Private:</h5>
                  <p class="mt-3 mb-3">
                    This will let someone join the Tag with a invite only
                  </p>
                  <b-button @click="makePrivate" variant="warning"
                    >Change</b-button
                  >
                </div>
              </template>
              <div class="w-50" v-if="isAdmin">
                <h5>Change Avatar</h5>
                <b-form-file
                  v-model="newAvatar"
                  :state="Boolean(newAvatar)"
                  placeholder="Choose a avatar or drop it here..."
                  drop-placeholder="Drop file here..."
                  accept="image/*"
                ></b-form-file>
                <div class="mt-3">
                  Selected file: {{ newAvatar ? newAvatar.name : "" }}
                </div>
                <b-button variant="success" @click="setNewAvatar()"
                  >Save</b-button
                >
              </div>
            </div>
            <div class="d-flex mt-5">
              <div class="w-50" v-if="!isOwner">
                <h5>Leave Group:</h5>
                <p>Will leave this group</p>
                <b-button variant="danger" v-b-modal.leave-tag>Leave</b-button>
              </div>
              <div class="w-50" v-if="isOwner">
                <h5>Delete Group:</h5>
                <h6 class="mt-3 mb-3">This will delete the tag completely</h6>
                <b-button variant="danger" v-b-modal.delete-tag
                  >Delete Tag</b-button
                >
              </div>
            </div>
          </template>
          <template v-else-if="userListOpened">
            <div class="d-flex justify-content-between">
              <h4>Users:</h4>
              <b-button
                @click="
                  [(settingsAreOpended = false), (userListOpened = false)]
                "
                variant="outline-primary"
                >Feed</b-button
              >
              <b-button
                @click="[(settingsAreOpended = true), (userListOpened = false)]"
                variant="outline-primary"
                >Settings</b-button
              >
            </div>
            <div class="table p-3">
              <div class="table-top grid mt-4 mb-4">
                <div><h6 class="grey">Avatar</h6></div>
                <div><h6 class="grey">Name</h6></div>
                <div><h6 class="grey">Is admin</h6></div>
                <div><h6 class="grey">Is Owner</h6></div>
                <div v-if="isAdmin || isOwner"><h6 class="grey">Kick</h6></div>
                <div v-if="isOwner"><h6 class="grey">Make Admin</h6></div>
              </div>
              <template v-for="(iter, index) in tag.members">
                <div :key="index" class="row grid flex-column">
                  <div>
                    <b-avatar :src="getAvatarUrl(iter.user.avatar)"></b-avatar>
                  </div>
                  <div>
                    <p>{{ iter.user.name + " " + iter.user.last_name }}</p>
                  </div>
                  <div>
                    <p>{{ iter.is_admin }}</p>
                  </div>
                  <div>
                    <p>{{ iter.is_owner }}</p>
                  </div>
                  <div>
                    <b-button size="sm" variant="danger height-80"
                      >Kick</b-button
                    >
                  </div>
                  <div>
                    <b-button size="sm" variant="warning height-80"
                      >Make admin</b-button
                    >
                  </div>
                </div>
              </template>
            </div>
          </template>
          <template v-else>
            <div class="d-flex justify-content-between">
              <h4>Activity Feed:</h4>
              <b-button
                @click="[(settingsAreOpended = false), (userListOpened = true)]"
                variant="outline-primary"
                >Users</b-button
              >
              <b-button
                @click="[(settingsAreOpended = true), (userListOpened = false)]"
                variant="outline-primary"
                >Settings</b-button
              >
            </div>
          </template>
        </div>
      </div>
    </template>
    <b-modal id="delete-tag" title="Delete this tag?" hide-footer>
      <p class="my-4">Do you wanna delete this tag?</p>
      <div class="d-flex justify-content-between">
        <b-button
          class="w-40"
          variant="danger"
          @click="[$bvModal.hide('delete-tag'), deleteTag()]"
          >Yes</b-button
        >
        <b-button
          class="w-40"
          variant="primary"
          @click="$bvModal.hide('delete-tag')"
          >No</b-button
        >
      </div>
    </b-modal>
    <b-modal id="leave-tag" title="Delete this tag?" hide-footer>
      <p class="my-4">Do you wanna leave this tag?</p>
      <div class="d-flex justify-content-between">
        <b-button
          class="w-40"
          variant="danger"
          @click="[$bvModal.hide('leave-tag'), leaveGroup()]"
          >Yes</b-button
        >
        <b-button
          class="w-40"
          variant="primary"
          @click="$bvModal.hide('leave-tag')"
          >No</b-button
        >
      </div>
    </b-modal>
  </div>
</template>

<script>
export default {
  name: "TagsProfile",
  data() {
    return {
      errors: {},
      alerts: {},
      tag: {},
      isMember: false,
      settingsAreOpended: false,
      newName: "",
      newAbout: "",
      userListOpened: false,
      isAdmin: false,
      isOwner: false,
      newAvatar: null
    };
  },
  methods: {
    getImgUrl() {
      let images = process.env.VUE_APP_API + "/static" + this.tag.avatar;
      return images;
    },
    getAvatarUrl(image) {
      let images = process.env.VUE_APP_API + "/static" + image;
      return images;
    },
    getTag() {
      this.errors = {};
      const vue = this;
      window.axios
        .get("/api/tags/tag/" + this.$route.params.id + "/profile")
        .then(res => {
          this.tag = res.data.tag;
          this.isMember = res.data.is_member;
          this.checkAdminAndOwner();
        })
        .catch(function(rej) {
          vue.errors = { error: rej.response.data.error };
        });
    },
    getMemberCount() {
      return this.tag.members.length;
    },
    joinTag() {
      this.errors = {};
      this.alerts = {};
      const vue = this;
      window.axios
        .post("/api/tags/tag/" + this.$route.params.id + "/join")
        .then(res => {
          this.alerts = { alert: res.data.message };
          this.getTag();
        })
        .catch(function(rej) {
          vue.errors = { error: rej.response.data.error };
        });
    },

    deleteTag() {
      this.errors = {};
      this.alerts = {};
      const vue = this;
      window.axios
        .post("/api/tags/tag/" + this.$route.params.id + "/delete")
        .then(() => {
          this.$router.push("/tags");
        })
        .catch(function(rej) {
          vue.errors = { error: rej.response.data.error };
        });
    },
    setNewName() {
      this.errors = {};
      this.alerts = {};
      const vue = this;
      window.axios
        .post("/api/tags/tag/" + this.$route.params.id + "/newname", {
          name: this.newName
        })
        .then(res => {
          this.alerts = { alert: res.data.message };
          this.newName = "";
          this.getTag();
        })
        .catch(function(rej) {
          vue.errors = { error: rej.response.data.error };
        });
    },
    setNewAbout() {
      this.errors = {};
      this.alerts = {};
      const vue = this;
      window.axios
        .post("/api/tags/tag/" + this.$route.params.id + "/newabout", {
          about: this.newAbout
        })
        .then(res => {
          this.alerts = { alert: res.data.message };
          this.newAbout = "";
          this.getTag();
        })
        .catch(function(rej) {
          vue.errors = { error: rej.response.data.error };
        });
    },
    setNewAvatar() {
      this.errors = {};
      this.alerts = {};
      const vue = this;
      let formData = new FormData();
      formData.append("file", this.newAvatar);
      window.axios
        .post(
          "/api/tags/tag/" + this.$route.params.id + "/setavatar",
          formData,
          {
            headers: {
              "Content-Type": "multipart/form-data"
            }
          }
        )
        .then(res => {
          this.alerts = { alert: res.data.message };
          this.newAvatar = null;
          this.getTag();
        })
        .catch(function(rej) {
          vue.errors = { error: rej.response.data.error };
        });
    },
    makePublic() {
      this.errors = {};
      this.alerts = {};
      const vue = this;
      window.axios
        .get("/api/tags/tag/" + this.$route.params.id + "/makePublic")
        .then(res => {
          this.alerts = { alert: res.data.message };
          this.getTag();
        })
        .catch(function(rej) {
          vue.errors = { error: rej.response.data.error };
        });
    },

    makePrivate() {
      this.errors = {};
      this.alerts = {};
      const vue = this;
      window.axios
        .get("/api/tags/tag/" + this.$route.params.id + "/makePrivate")
        .then(res => {
          this.alerts = { alert: res.data.message };
          this.getTag();
        })
        .catch(function(rej) {
          vue.errors = { error: rej.response.data.error };
        });
    },
    leaveGroup() {
      this.errors = {};
      this.alerts = {};
      const vue = this;
      window.axios
        .post("/api/tags/tag/" + this.$route.params.id + "/leave")
        .then(() => {
          this.$router.push("/tags");
        })
        .catch(function(rej) {
          vue.errors = { error: rej.response.data.error };
        });
    },
    checkAdminAndOwner() {
      for (let iter = 0; iter < this.tag.members.length; iter++) {
        if (this.tag.members[iter].UserID == this.$route.params.id) {
          this.isAdmin = this.tag.members[iter].is_admin;
          this.isOwner = this.tag.members[iter].is_owner;
          return;
        }
      }
    }
  },
  mounted() {
    this.getTag();
  }
};
</script>

<style scoped lang="scss">
.w-40 {
  width: 40%;
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
    min-height: 500px;
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
.grey {
  color: #b0bec5;
}

.grid {
  display: grid;
  grid-template-columns: 100px auto auto auto auto auto;
  grid-template-rows: 50px;
  h6 {
    text-align: center;
  }

  p {
    text-align: center;
  }
  div {
    display: flex;
    justify-content: center;
  }

  border-bottom: 1px solid #b0bec5;
  margin-bottom: 20px;
}
</style>
