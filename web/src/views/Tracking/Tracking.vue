<template>
  <b-container class="mt-3 container">
    <b-button
        class="mr-3"
        variant="outline-info"
        v-if="creatorIsOpended"
        @click="creatorIsOpended = false"
    >View submissions</b-button
    >
    <template v-if="currentUser.role === 'admin' || isManager">
      <div class="d-flex justify-content-end">
        <b-button
          class="mr-3"
          variant="outline-info"
          v-if="!creatorIsOpended"
          @click="creatorIsOpended = true"
          >Create submissions</b-button
        >
        <a href="/tracking/admin">
          <b-button
            class="mr-3"
            v-if="currentUser.role === 'admin'"
            variant="outline-info"
            >Admin Panel</b-button
          >
        </a>
        <a href="/tracking/manager"
          ><b-button v-if="isManager" variant="outline-info"
            >Manager Panel</b-button
          ></a
        >
      </div>
    </template>
    <div v-if="creatorIsOpended" class="mt-4 form">
      <h3>Submissions:</h3>
      <div class="mt-4">
        <h5 class="mb-2">Subject:</h5>
        <b-form-input
          class="w-75"
          v-model="form.subject"
          placeholder="Enter Subject"
        ></b-form-input>
      </div>

      <div class="mt-4">
        <h5 class="mb-2">Description:</h5>
        <b-form-textarea
          v-model="form.description"
          placeholder="Enter Description"
          rows="3"
          max-rows="6"
        ></b-form-textarea>
      </div>
      <div class="mt-4">
        <h5 class="mb-2">Attachments:</h5>
        <b-form-file
          class="w-75"
          multiple
          v-model="form.files"
          :state="Boolean(form.files)"
          placeholder="Choose a file or drop it here..."
          drop-placeholder="Drop file here..."
        ></b-form-file>
      </div>
      <div class="d-flex justify-content-between mt-5">
        <div></div>
        <b-button @click="goSubmit" variant="success">Submit</b-button>
      </div>
    </div>
    <div v-else class="mt-4 form">
      <h3>Submissions:</h3>
      <vue-good-table
        :search-options="{ enabled: true }"
        :pagination-options="{
          enabled: true,
          mode: 'records',
          perPage: 10,
          position: 'top',
          nextLabel: 'next',
          prevLabel: 'prev',
          rowsPerPageLabel: 'Rows per page',
          ofLabel: 'of',
          pageLabel: 'page', // for 'pages' mode
          allLabel: 'All'
        }"
        :columns="columns"
        :rows="submissions"
      >
        <template slot="table-row" slot-scope="props">
          <div
            v-if="props.column.field === 'tracked_attachments'"
            class="table-row d-flex justify-content-center"
          >
            <template v-if="props.formattedRow.tracked_attachments.length > 0">
              <b-button
                @click="downloadFiles(props.formattedRow.tracked_attachments)"
                variant="outline-info"
                >Download</b-button
              >
            </template>
            <template v-else>
              <p>No attachments</p>
            </template>
          </div>
          <div v-else class="table-row d-flex justify-content-center">
            {{ props.formattedRow[props.column.field] }}
          </div>
        </template>
      </vue-good-table>
    </div>
  </b-container>
</template>

<script>
export default {
  name: "Tracking",
  computed: {
    currentUser() {
      return this.$store.getters.currentUser;
    }
  },
  data() {
    return {
      errors: {},
      columns: [
        {
          label: "ID",
          field: "id",
          sortable: true,
          type: "number"
        },
        {
          label: "Subject",
          field: "subject"
        },
        {
          label: "Description",
          field: "description",
          sortable: true
        },
        {
          label: "Date submitted",
          field: "submit_date",
          sortable: true
        },
        {
          label: "Accepted",
          field: "is_confirmed",
          sortable: true,
          filterOptions: {
            enabled: true, // enable filter for this column
            filterValue: "True", // initial populated value for this filter
            filterDropdownItems: ["True", "False"] // dropdown (with selected values) instead of text input
          }
        },
        {
          label: "Attachments",
          field: "tracked_attachments",
          sortable: false
        }
      ],
      isManager: false,
      blogs: [],
      submissions: [],
      creatorIsOpended: true,
      formID: 0,
      form: {
        subject: "",
        description: "",
        files: null
      }
    };
  },
  methods: {
    getIsManager() {
      window.axios
        .get("api/manager/check/ismanager/")
        .then(res => {
          this.isManager = res.data.isManager;
        })
        .catch(rej => {
          console.log(rej.response.data.error);
        });
    },
    async goSubmit() {
      if (this.form.subject != "" && this.form.description != "") {
        let formData = new FormData();
        formData.append("subject", `${this.form.subject}`);
        formData.append("description", `${this.form.description}`);
        await window.axios
          .post("api/tracking/add/", formData, {
            headers: {
              "Content-Type": "multipart/form-data"
            }
          })
          .then(res => {
            this.formID = res.data.id;
            this.makeToast("Submitted successfully", "success");
            this.form.subject = "";
            this.form.description = "";
          })
          .catch(rej => {
            this.makeToast(rej.response.data.error, "danger");
          });

        for (let iter = 0; iter < this.form.files.length; iter += 1) {
          let formData = new FormData();
          console.log(this.form.files[iter]);
          console.log(this.formID);
          formData.append("file", this.form.files[iter]);
          formData.append("id", this.formID);
          window.axios
            .post("api/tracking/add/attachment/", formData, {
              headers: {
                "Content-Type": "multipart/form-data"
              }
            })
            .then(() => {
              console.log("saved");
            })
            .catch(rej => {
              this.makeToast(rej.response.data.error, "danger");
            });
        }
        this.formID = 0;
        this.form.files = null;
        this.getSubmissions();
      } else {
        this.makeToast("Subject or description wasn't filled", "danger");
      }
    },
    downloadFiles(files) {
      for (let iter = 0; iter < files.length; iter += 1) {
        window.axios
          .get("static" + files[iter].path, { responseType: "blob" })
          .then(res => {
            let filename;
            if (res.data.type == "application/pdf") {
              filename = "attachment.pdf";
            } else {
              filename = "attachment.png";
            }
            const blob = new Blob([res.data]);
            const link = document.createElement("a");
            link.href = URL.createObjectURL(blob);
            link.download = filename;
            link.click();
            URL.revokeObjectURL(link.href);
          })
          .catch(rej => {
            this.makeToast(rej.response.data.error, "danger");
          });
      }
    },
    makeToast(text, variant) {
      this.$bvToast.toast(text, {
        autoHideDelay: 5000,
        variant: variant,
        title: "Notification"
      });
    },
    getSubmissions() {
      window.axios
        .get("api/tracking/user/list")
        .then(res => {
          this.submissions = res.data.submissions;
        })
        .catch(() => {
          this.makeToast(
            "There was an error getting your submissions",
            "danger"
          );
        });
    }
  },
  created() {
    this.getIsManager();
    this.getSubmissions();
  }
};
</script>

<style scoped lang="scss">
.form {
  h5 {
    margin: 0;
  }
}
.container {
  height: 100%;
}
</style>
