<template>
  <div class="mt-4 form">
    <div class="mb-3 d-flex justify-content-between">
      <h3>Submissions:</h3>
      <b-button variant="outline-primary" @click="$router.push('/tracking')"
        >Submissions</b-button
      >
    </div>
    <vue-good-table
      class="mb-4"
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
          v-if="props.column.field === 'author'"
          class="table-row d-flex justify-content-center"
        >
          <p>{{ props.row.user.name + " " + props.row.user.last_name }}</p>
        </div>
        <div
          v-if="props.column.field === 'accept'"
          class="table-row d-flex justify-content-center"
        >
          <b-button
            @click="closeSubmission(props.row.id)"
            variant="outline-info"
            v-if="!props.formattedRow.accepted"
            >Close</b-button
          >
          <b-button
            @click="openSubmission(props.row.id)"
            variant="outline-warning"
            v-if="props.formattedRow.accepted"
            >Open</b-button
          >
        </div>
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
</template>

<script>
export default {
  name: "ManagerPanelTracking",
  data() {
    return {
      errors: {},
      columns: [
        {
          label: "Author",
          field: "author"
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
          label: "Closed",
          field: "is_confirmed",
          sortable: true,
          filterOptions: {
            enabled: true, // enable filter for this column
            filterValue: "False", // initial populated value for this filter
            filterDropdownItems: ["True", "False"] // dropdown (with selected values) instead of text input
          }
        },
        {
          label: "Attachments",
          field: "tracked_attachments",
          sortable: false
        },
        {
          label: "Accept",
          field: "accept"
        }
      ],
      isManager: false,
      submissions: []
    };
  },
  methods: {
    getIsManager() {
      window.axios
        .get("api/manager/check/ismanager/")
        .then(res => {
          this.isManager = res.data.isManager;
          if (this.isManager == false) {
            this.$router.push("/tracking");
          }
        })
        .catch(rej => {
          console.log(rej.response.data.error);
        });
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
        .get("api/tracking/manager/list/")
        .then(res => {
          this.submissions = res.data.submissions;
        })
        .catch(() => {
          this.makeToast(
            "There was an error getting your submissions",
            "danger"
          );
        });
    },
    closeSubmission(id) {
      window.axios
        .get(`api/tracking/manager/item/${id}/close`)
        .then(() => {
          this.getSubmissions();
        })
        .catch(() => {
          this.makeToast("There was an error closing submission", "danger");
        });
    },
    openSubmission(id) {
      window.axios
        .get(`api/tracking/manager/item/${id}/open`)
        .then(() => {
          this.getSubmissions();
        })
        .catch(() => {
          this.makeToast("There was an error opening submission", "danger");
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
  width: 90%;
  margin: auto;
  h5 {
    margin: 0;
  }
}
.container {
  height: 100%;
}
</style>
