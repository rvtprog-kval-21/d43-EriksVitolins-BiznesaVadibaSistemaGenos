<template>
  <b-container class="container">
    <div class="d-flex justify-content-between mt-4">
      <div></div>
      <b-button variant="outline-primary" @click="$router.push('/projects')"
        >Projects</b-button
      >
    </div>
    <div class="">
      <div class="row-blog">
        <h5>Project name:</h5>
        <b-form-input
          class="w-75"
          v-model="form.name"
          placeholder="Enter project name"
        ></b-form-input>
      </div>
    </div>
    <div>
      <div class="row-blog">
        <h5>Project avatar:</h5>
        <b-form-file
          v-model="form.avatar"
          :state="Boolean(form.avatar)"
          class="w-75"
          placeholder="Choose a photo or drop it here..."
          drop-placeholder="Drop photo here..."
          accept="image/*"
        ></b-form-file>
      </div>
    </div>
    <div>
      <div class="row-blog">
        <h5>Project About:</h5>
        <quill-editor
          :content="form.about"
          :options="editorOption"
          @change="onEditorChange($event)"
        />
      </div>
    </div>
    <div>
      <div class="row-blog">
        <h5>Project Owners:</h5>
        <vSelect
          class="w-75"
          multiple
          v-model="form.owners"
          label="email"
          :options="options"
        />
      </div>
      <div class="mt-1 mb-4"></div>
    </div>
    <div class="d-flex justify-content-between">
      <div></div>
      <b-button @click="addProjects" variant="success">Save</b-button>
    </div>
  </b-container>
</template>

<script>
import "vue-select/dist/vue-select.css";
import vSelect from "vue-select";

export default {
  name: "CreateProjects",
  components: { vSelect },
  data() {
    return {
      form: {
        name: "",
        avatar: null,
        about: "",
        owners: []
      },
      tags: [],
      options: [],
      editorOption: {
        // Some Quill options...
      }
    };
  },
  methods: {
    addProjects() {
      let owners = "";
      if (this.form.owners.length > 0) {
        for (let iter = 0; iter < this.form.owners.length; iter++) {
          if (iter === 0) {
            owners += this.form.owners[iter].id;
          } else {
            owners += ";" + this.form.owners[iter].id;
          }
        }
      }
      if (this.form.name !== "" && owners !== "") {
        let formData = new FormData();
        formData.append("about", `${this.form.about}`);
        formData.append("name", `${this.form.name}`);
        formData.append("users", `${owners}`);
        formData.append("avatar", this.form.avatar);
        window.axios
          .post("api/projects/add/new", formData, {
            headers: {
              "Content-Type": "multipart/form-data"
            }
          })
          .then(() => {
            this.makeToast("Project created successfully", "success");
            this.$router.push("/projects");
          })
          .catch(rej => {
            this.makeToast(rej.response.data.error, "danger");
          });
      } else {
        this.makeToast("Name or owners field not filled", "danger");
      }
    },
    onEditorChange({ quill, html, text }) {
      console.log("editor change!", quill, html, text);
      this.form.about = html;
    },
    makeToast(text, variant) {
      this.$bvToast.toast(text, {
        autoHideDelay: 5000,
        variant: variant,
        title: "Notification"
      });
    },
    getUsers() {
      window.axios.post("api/users").then(res => {
        this.options = res.data.users;
      });
    },
    isAdmin() {
      if (this.currentUser.role === "admin") {
        this.$router.push("/projects");
      }
    }
  },
  created() {
    this.getUsers();
  }
};
</script>

<style scoped lang="scss">
.row-blog {
  display: flex;
  margin-top: 20px;
  align-items: center;
  justify-content: space-between;
  h5 {
    margin-right: 25px;
    margin-bottom: 0;
  }
}

.container {
  height: 100%;
}
</style>
