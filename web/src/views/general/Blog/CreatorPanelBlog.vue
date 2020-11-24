<template>
  <b-container class="mt-3 body">
    <div class="d-flex justify-content-end">
      <a href="/blog"><b-button variant="outline-info">Back</b-button></a>
    </div>
    <hr />
    <div class="mt-3 mb-3">
      <div class="d-flex justify-content-end">
        <b-button
          v-if="!isCreating && !isEditing"
          @click="isCreating = true"
          variant="outline-primary"
          >Create a Post</b-button
        >
        <b-button
          v-else
          @click="[(isCreating = false), (isEditing = false)]"
          variant="outline-primary"
          >Home</b-button
        >
      </div>
      <template v-if="isCreating">
        <CreateBlog></CreateBlog>
      </template>
      <template v-if="isEditing">
        <CreateBlog :editID="idBlog" :isEdit="true"></CreateBlog>
      </template>
      <template v-if="!isEditing && !isCreating">
        <vue-good-table
          class="mt-3"
          :search-options="{ enabled: true }"
          :columns="columns"
          :rows="blogs"
          :pagination-options="{
            enabled: true,
            mode: 'records',
            perPage: 5,
            position: 'top',
            perPageDropdown: [10, 15, 20],
            dropdownAllowAll: false,
            setCurrentPage: 2,
            nextLabel: 'next',
            prevLabel: 'prev',
            rowsPerPageLabel: 'Rows per page',
            ofLabel: 'of',
            pageLabel: 'page', // for 'pages' mode
            allLabel: 'All'
          }"
        >
          <template slot="table-row" slot-scope="props">
            <div
              v-if="props.column.field === 'view'"
              class="table-row d-flex justify-content-center"
            >
              <b-button @click="goToBlog(props.formattedRow.id)" variant="outline-info">View</b-button>
            </div>
            <div
              v-else-if="props.column.field === 'edit'"
              class="d-flex justify-content-center table-row"
            >
              <b-button @click="editArticle(props.formattedRow.id)" variant="info">Edit</b-button>
            </div>
            <div v-else-if="props.column.field === 'count'"
                    class="d-flex justify-content-center table-row">
              <p>{{getCount(props.formattedRow.id)}}</p>
            </div>
            <div
              v-else-if="props.column.field === 'deleted_at'"
              class="table-row d-flex justify-content-center"
            >
              <b-button v-if="props.formattedRow.deleted_at === '0001-01-01T00:00:00Z'" @click="deleteBlog(props.formattedRow.id)" variant="danger">Delete</b-button>
              <b-button @click="unDeleteBlog(props.formattedRow.id)" variant="outline-success" v-else>Undelete</b-button>
            </div>
            <div v-else class="table-row d-flex justify-content-center">
              {{ props.formattedRow[props.column.field] }}
            </div>
          </template>
        </vue-good-table>
      </template>
    </div>
  </b-container>
</template>

<script>
import CreateBlog from "../../../components/blog/CreateBlog";
export default {
  name: "CreatorPanelBlog",
  components: { CreateBlog },
  data() {
    return {
      isCreating: false,
      isEditing: false,
      idBlog: 0,
      blogs: [],
      count: {},
      columns: [
        {
          label: "Id",
          field: "id",
          type: "number"
        },
        {
          label: "Title",
          field: "title"
        },
        {
          label: "Topic",
          field: "topic"
        },
        {
          label: "View",
          field: "view"
        },
        {
          label: "Edit",
          field: "edit"
        },
        {
          label: "Delete",
          field: "deleted_at",
        },
        {
          label: "View",
          field: "count",
        },
      ]
    };
  },
  methods: {
    getBlogs() {
      window.axios
        .get("api/blog/owner")
        .then(res => {
          this.blogs = res.data.blogs;
          this.getCounts()
        })
        .catch(rej => {
          this.makeToast(rej.response.data.error, "danger");
        });
    },
    getCounts() {
     for (let iter = 0; iter < this.blogs.length; iter++) {
       const cond = `id${this.blogs[iter].id}`
       window.axios
               .get(`api/blog/get/${this.blogs[iter].id}/count`)
               .then(res => {
                 this.count[cond] = res.data.count
               })
     }
    },
    getCount(id) {
      const cond = `id${id}`
      return this.count[cond]
    },
    goToBlog(field) {
      this.$router.push("/blog/" + field + "/view")
    },
    deleteBlog(field) {
      window.axios
      .get(`api/blog/owner/get/${field}/delete`)
              .then(rej => {
                this.getBlogs()
                this.makeToast(rej.data.message, "success");
              })
              .catch(rej => {
                this.makeToast(rej.response.data.error, "danger");
              });
    },
    unDeleteBlog(field) {
      window.axios
              .get(`api/blog/owner/get/${field}/undelete`)
              .then(rej => {
                this.getBlogs()
                this.makeToast(rej.data.message, "success");
              })
              .catch(rej => {
                this.makeToast(rej.response.data.error, "danger");
              });
    },
    editArticle(field){
      this.idBlog = field
      this.isEditing = true
    },
    makeToast(text, variant) {
      this.$bvToast.toast(text, {
        autoHideDelay: 5000,
        variant: variant,
        title: "Notification"
      });
    }
  },
  mounted() {
    this.getBlogs();
  }
};
</script>

<style>
.container {
  min-height: 800px;
}
</style>
