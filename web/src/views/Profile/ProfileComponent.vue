<template>
  <div class="background">
    <div class="content">
      <alertComponent v-if="this.errors.error" class="alert-danger">
        <p>{{ this.errors.error }}</p>
      </alertComponent>
      <alertComponent v-if="this.alerts.message" class="alert-success">
        <p>{{ this.alerts.message }}</p>
      </alertComponent>
      <div class="user" v-if="!showAdminSettings">
        <div class="backgroundHeader">
          <img :src="getImgUrl()" alt="" />
        </div>
        <div class="card">
          <div class="name">
            <p v-if="getName && getLastName">
              {{ this.user.userdata.name + " " + this.user.userdata.lastname }}
            </p>
            <p v-else>...</p>
          </div>
          <div class="main">
            <div class="category">
              <div class="header">
                <p>Personal info</p>
                <hr />
              </div>
              <div class="info">
                <div class="holder">
                  <p class="title">About:</p>
                  <p v-if="getAbout" class="result">
                    {{ this.user.userdata.about }}
                  </p>
                  <p v-else class="result">Nan</p>
                </div>

                <div class="holder">
                  <p class="title">Phone:</p>
                  <p v-if="getAbout" class="result">
                    {{ this.user.userdata.about }}
                  </p>
                  <p v-else class="result">Nan</p>
                </div>
                <div class="holder">
                  <p class="title">Birthday:</p>
                  <p v-if="getAbout" class="result">
                    {{ this.user.userdata.about }}
                  </p>
                  <p v-else class="result">Nan</p>
                </div>
                <div class="holder">
                  <p class="title">Nameday:</p>
                  <p v-if="getAbout" class="result">
                    {{ this.user.userdata.about }}
                  </p>
                  <p v-else class="result">Nan</p>
                </div>
              </div>
            </div>
            <div class="category">
              <div class="header">
                <p>Work info</p>
                <hr />
              </div>
              <div class="info">
                <div class="holder">
                  <p class="title">Role:</p>
                  <p class="result">{{ this.user.role }}</p>
                </div>
                <div class="holder">
                  <p class="title">Title:</p>
                  <p v-if="getTitle" class="result">
                    {{ this.user.userdata.title }}
                  </p>
                  <p v-else class="result">Nan</p>
                </div>
                <div class="holder">
                  <p class="title">Tags:</p>
                  <p class="result">{{ this.user.created_at }}</p>
                </div>
                <div class="holder">
                  <p class="title">Location:</p>
                  <p class="result">{{ this.user.created_at }}</p>
                </div>
              </div>
            </div>
            <div class="category" v-if="currentUser.role === 'admin'">
              <div class="header">
                <p>Admin info</p>
                <hr />
              </div>
              <div class="info">
                <div class="holder">
                  <p class="title">Created_at:</p>
                  <p class="result">{{ this.user.created_at }}</p>
                </div>
                <div class="holder">
                  <p class="title">Updated_at:</p>
                  <p class="result">{{ this.user.updated_at }}</p>
                </div>
              </div>
            </div>
          </div>
          <div class="button-holder">
            <div class="buttons">
              <button v-if="currentUser.id === user.id">Change settings</button>
              <button
                @click="showAdminSettings = !showAdminSettings"
                v-if="currentUser.role === 'admin'"
              >
                Admin settings
              </button>
            </div>
          </div>
        </div>
      </div>
      <div class="admin" v-if="showAdminSettings">
        <div class="header-admin">
          <p v-if="getName && getLastName">
            {{ user.userdata.name + " " + user.userdata.lastname }}
          </p>
          <svg
            @click="showAdminSettings = !showAdminSettings"
            xmlns="http://www.w3.org/2000/svg"
            class="icon icon-tabler icon-tabler-arrow-left"
            width="36"
            height="36"
            viewBox="0 0 24 24"
            stroke-width="2"
            stroke="#3F51B5"
            fill="none"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <path stroke="none" d="M0 0h24v24H0z" fill="none" />
            <line x1="5" y1="12" x2="19" y2="12" />
            <line x1="5" y1="12" x2="11" y2="18" />
            <line x1="5" y1="12" x2="11" y2="6" />
          </svg>
        </div>

        <div class="card">
          <div class="holder">
            <div class="title">
              <p>Lock Account!</p>
            </div>
            <button @click="lock" v-if="this.user.locked === 0">Lock</button>
            <button v-else @click="lock">Unlock</button>
          </div>
        </div>

        <div class="card">
          <div class="holder">
            <div class="title">
              <p>Init password reset!</p>
            </div>
            <button @click="resetPassword">Reset</button>
          </div>
        </div>

        <div class="card" v-if="user.temp_password">
          <div class="holder">
            <div class="title">
              <p>Temp password!</p>
            </div>
            <p v-if="showPassword">{{ user.temp_password }}</p>
            <p v-if="!showPassword">************</p>
            <button v-if="!showPassword" @click="showPassword = !showPassword">
              Show
            </button>
            <button v-if="showPassword" @click="showPassword = !showPassword">
              Hide
            </button>
          </div>
        </div>
        <div class="card big">
          <div class="holder">
            <div class="title">
              <p>Add tags!</p>
            </div>
            <div class="input">
              <vue-tags-input
                v-model="tag"
                :tags="tags"
                @tags-changed="newTags => (tags = newTags)"
                :add-only-from-autocomplete="true"
                :autocomplete-items="filteredTags"
              />
            </div>
            <button>Save</button>
          </div>
        </div>

        <div class="card big last">
          <div class="holder">
            <div class="title">
              <p>Add Location!</p>
            </div>
            <div class="input">
              <vue-tags-input
                v-model="location"
                :tags="locations"
                @tags-changed="newTags => (locations = newTags)"
                :add-only-from-autocomplete="true"
                :autocomplete-items="filteredLocations"
              />
            </div>
            <button>Save</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import VueTagsInput from "@johmun/vue-tags-input";
import alertComponent from "../../components/alerts/alert-component";

export default {
  name: "ProfileComponent",
  computed: {
    currentUser() {
      return this.$store.getters.currentUser;
    }
  },
  components: {
    VueTagsInput,
    alertComponent
  },
  data() {
    return {
      user: {},
      showAdminSettings: false,
      errors: {},
      alerts: {},
      accountLocked: false,
      tag: "",
      tags: [],
      filteredTags: [
        "Automation",
        "Developer",
        "Hr",
        "Manual",
        "Finances",
        "Project Lead"
      ],
      location: "",
      locations: [],
      filteredLocations: [
        "Riga",
        "Ventspils",
        "Liepaja",
        "Jelgava",
        "Daugavpils"
      ],
      showPassword: false
    };
  },
  mounted() {
    this.getUser();
  },
  methods: {
    lock: function() {
      this.errors = [];
      this.alerts = [];
      const vue = this;
      window.axios
        .post("/api/auth/user/lock", { id: this.user.id })
        .then(response => {
          this.alerts = { message: response.data.message };
          this.getUser();
        })
        .catch(function(errors) {
          vue.errors = { error: errors.response.data.message };
        });
    },
    getName: function() {
      return this.user.userdata.name.length;
    },
    getLastName: function() {
      return this.user.userdata.lastname.length;
    },
    getTitle: function() {
      return this.user.userdata.title.length;
    },
    getAbout: function() {
      return this.user.userdata.about.length;
    },
    getUser: function() {
      this.errors = {};
      const vue = this;
      window.axios
        .get("/api/user/" + this.$route.params.id)
        .then(res => {
          this.user = res.data.data;
        })
        .catch(function(rej) {
          vue.errors = { error: rej.response.data.message };
        });
    },
    getImgUrl() {
      let images =
        process.env.VUE_APP_API +
        "/storage/avatars/" +
        this.user.id +
        "/avatar.png";
      return images;
    },
    resetPassword() {
      this.errors = [];
      this.alerts = [];
      const vue = this;
      window.axios
        .post("/api/auth/user/passreset", { id: this.user.id })
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
.background {
  background-color: #f5f5f5;
  width: 100%;
  display: flex;
  justify-content: center;
  overflow: auto;
}

.content {
  width: 80%;
  display: flex;
  flex-direction: column;
  align-items: center;

  .user {
    width: 100%;
  }
  .backgroundHeader {
    width: 100%;
    height: 200px;
    display: flex;
    justify-content: center;
    align-items: center;
  }
  img {
    width: 150px;
    height: 130px;
  }
  .card {
    width: 100%;
    background: white;
    overflow: auto;
    .name {
      display: flex;
      justify-content: center;
      margin-top: 10px;
      p {
        font-size: 28px;
        font-weight: bolder;
      }
    }
    .main {
      margin-top: 15px;
      .header {
        width: 80%;
        display: flex;
        flex-direction: column;
        margin-left: auto;
        align-items: center;
        margin-right: auto;
        margin-bottom: 20px;
        hr {
          border: 0;
          clear: both;
          display: block;
          width: 44%;
          background-color: grey;
          height: 1px;
          margin-bottom: 10px;
        }
      }
      .category {
        .info {
          display: grid;
          grid-template-columns: auto auto;
          .holder {
            margin-left: 15px;
            margin-bottom: 20px;
            display: flex;
            .title {
              font-weight: bolder;
            }
            p {
              font-size: 18px;
            }
            .result {
              margin-left: 10px;
            }
          }
        }
      }
    }
    .button-holder {
      display: flex;
      justify-content: center;
      .buttons {
        width: 90%;
        display: grid;
        grid-template-columns: auto;
        grid-template-rows: 50px;
        grid-gap: 40px;
        margin-bottom: 40px;
        margin-top: 10px;
        button {
          background-color: #e57373;
          border: none;
          color: white;
          cursor: pointer;
          &:hover {
            background-color: #af4448;
          }
        }
      }
    }
  }
  .admin {
    width: 100%;
    .header-admin {
      display: flex;
      justify-content: space-between;
      p {
        color: #636b6f;
      }

      svg:hover {
        stroke: red;
        cursor: pointer;
      }
    }

    .card {
      margin-top: 30px;
      min-height: 100px;
      &.big {
        min-height: 200px;
      }
      &.last {
        margin-bottom: 40px;
      }
      .holder {
        padding: 20px;
        display: flex;
        justify-content: space-between;
        align-items: center;
        p {
          color: #636b6f;
        }
        button {
          background-color: #e57373;
          border: none;
          color: white;
          cursor: pointer;
          padding: 15px 30px 15px 30px;
          &:hover {
            background-color: #af4448;
          }
        }
      }
    }
  }
}
</style>
