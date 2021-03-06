<template>
  <b-container class="mt-3">
    <div class="d-flex justify-content-between">
      <h3>Settings</h3>
      <b-button variant="outline-primary" @click="goBack()">Back</b-button>
    </div>
    <div
      class="borders d-flex justify-content-between mt-5"
      v-if="this.user.is_owner"
    >
      <h4>Archive Project</h4>
      <b-button variant="danger" @click="deleteProject()">Archive</b-button>
    </div>
    <div
      class="borders d-flex justify-content-between mt-5"
      v-if="this.user.is_owner"
    >
      <h4>Project name:</h4>
      <b-form-input
        class="w-75"
        v-model="name"
        placeholder="Enter project name"
      ></b-form-input>
      <b-button @click="saveName()" variant="success" v-if="name != ''"
        >Save</b-button
      >
    </div>
    <div class="borders d-flex justify-content-between mt-5">
      <h5>Project Tags:</h5>
      <v-select class="w-75" multiple v-model="tags" taggable push-tags />
      <b-button @click="saveTags()" variant="success" v-if="tags.length > 0"
        >Save</b-button
      >
    </div>
    <div class="borders d-flex justify-content-between mt-5">
      <h4>Project avatar:</h4>
      <b-form-file
        v-model="avatar"
        :state="Boolean(avatar)"
        class="w-75"
        placeholder="Choose a photo or drop it here..."
        drop-placeholder="Drop photo here..."
        accept="image/*"
      ></b-form-file>
      <b-button variant="success" @click="saveAvatar()" v-if="avatar != null"
        >Save</b-button
      >
    </div>
    <div class="borders d-flex justify-content-between mt-5">
      <h5>Invite to Project:</h5>
      <vSelect
        class="w-75"
        multiple
        v-model="invitees"
        label="email"
        :options="options"
      />
      <b-button
        variant="success"
        @click="inviteUsers()"
        v-if="invitees.length > 0"
        >Add</b-button
      >
    </div>
    <div class="borders mt-5">
      <div class="d-flex justify-content-between editor">
        <h5>Project About:</h5>
        <quill-editor
          class="w-75"
          :content="about"
          :options="editorOption"
          @change="onEditorChange($event)"
        />
        <b-button variant="success" @click="saveAbout()" v-if="about != ''"
          >Save</b-button
        >
      </div>
    </div>
  </b-container>
</template>

<script>
import "vue-select/dist/vue-select.css";
import vSelect from "vue-select";

export default {
  name: "Settings.vue",
  components: { vSelect },
  data() {
    return {
      project: null,
      user: {},
      aboutIsOpened: false,
      isLoading: false,
      name: "",
      avatar: null,
      about: "",
      tags: [],
      options: [],
      invitees: [],
      editorOption: {
        // Some Quill options...
      }
    };
  },
  methods: {
    getUser() {
      const users = this.project.members;
      this.user = users.find(e => e.user.id === this.currentUser.id);
      console.log(this.user);
      if (this.user) {
        if (this.user.is_admin || !this.user.is_owner) {
          return;
        }
      }
      this.$router.push(`/projects/${this.$route.params.id}/see`);
    },
    getProject() {
      this.isLoading = true;
      this.tags = [];
      window.axios
        .get(`api/projects/get/${this.$route.params.id}/item/`)
        .then(res => {
          this.project = res.data.project;
          this.name = this.project.name;
          this.about = this.project.about;
          for (let i = 0; i < this.project.tags.length; i++) {
            this.tags.push(this.project.tags[i].name);
          }
          this.getUser();
          this.isLoading = false;
        })
        .catch(() => {
          this.$router.push("/projects");
        });
    },
    makeToast(text, variant) {
      this.$bvToast.toast(text, {
        autoHideDelay: 5000,
        variant: variant,
        title: "Notification"
      });
    },
    onEditorChange({ quill, html, text }) {
      console.log("editor change!", quill, html, text);
      this.about = html;
    },
    goBack() {
      this.$router.push(`/projects/${this.$route.params.id}/see`);
    },
    deleteProject() {
      window.axios
        .get(`api/projects/remove/${this.$route.params.id}/project/`)
        .then(() => {
          this.$router.push("/projects");
        })
        .catch(rej => {
          this.makeToast(rej.response.data.error, "danger");
        });
    },
    getUsers() {
      window.axios
        .get(`api/projects/get/${this.$route.params.id}/nonmembers/`)
        .then(res => {
          this.options = res.data;
        });
    },
    inviteUsers() {
      window.axios
        .post(`api/projects/invite/${this.$route.params.id}/users/`, {
          users: this.invitees
        })
        .then(() => {
          this.makeToast("Users were added", "success");
          this.invitees = [];
          this.getProject();
          this.getUsers();
        })
        .catch(rej => {
          this.makeToast(rej.response.data.error, "danger");
        });
    },
    saveName() {
      window.axios
        .post(`api/projects/change/${this.$route.params.id}/name/`, {
          name: this.name
        })
        .then(() => {
          this.makeToast("Name changed", "success");
          this.getProject();
        })
        .catch(rej => {
          this.makeToast(rej.response.data.error, "danger");
        });
    },
    saveTags() {
      window.axios
        .post(`api/projects/change/${this.$route.params.id}/tags/`, {
          tags: this.tags
        })
        .then(() => {
          this.makeToast("Tags changed", "success");
          this.getProject();
        })
        .catch(rej => {
          this.makeToast(rej.response.data.error, "danger");
        });
    },
    saveAbout() {
      window.axios
        .post(`api/projects/change/${this.$route.params.id}/about/`, {
          about: this.about
        })
        .then(() => {
          this.makeToast("About changed", "success");
          this.getProject();
        })
        .catch(rej => {
          this.makeToast(rej.response.data.error, "danger");
        });
    },
    saveAvatar() {
      let formData = new FormData();
      formData.append("avatar", this.avatar);
      window.axios
        .post(
          `api/projects/change/${this.$route.params.id}/avatar/`,
          formData,
          {
            headers: {
              "Content-Type": "multipart/form-data"
            }
          }
        )
        .then(() => {
          this.makeToast("Project avatar changed successfully", "success");
          this.getProject();
        })
        .catch(rej => {
          this.makeToast(rej.response.data.error, "danger");
        });
    }
  },
  computed: {
    currentUser() {
      return this.$store.getters.currentUser;
    }
  },
  async created() {
    await this.getProject();
    this.getUsers();
  }
};
</script>

<style scoped>
.borders {
  padding-top: 5px;
  padding-bottom: 5px;
  border-bottom: solid #607d8b 1px;
  border-top: solid #607d8b 1px;
}
.editor {
  min-height: 440px;
}
</style>
