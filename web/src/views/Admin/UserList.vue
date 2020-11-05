<template>
  <div class="overflow-auto">
    <div class="">
      <b-alert v-if="this.errors.error" variant="danger" show="">{{
        this.errors.error
      }}</b-alert>
    </div>
    <b-row class="filter ml-2">
      <b-col>
        <p>User count: {{ this.users.length }}</p>
      </b-col>
      <b-col>
        <b-form-group
          label="Filter"
          label-cols-sm="3"
          label-align-sm="right"
          label-size="sm"
          label-for="filterInput"
          class="mb-0"
        >
          <b-input-group size="sm">
            <b-form-input
              v-model="filter"
              type="search"
              id="filterInput"
              placeholder="Type to Search"
            ></b-form-input>
            <b-input-group-append>
              <b-button :disabled="!filter" @click="filter = ''"
                >Clear</b-button
              >
            </b-input-group-append>
          </b-input-group>
        </b-form-group>
      </b-col>
    </b-row>
    <div class="table-holder ml-3 mr-3">
      <b-table
        :filter="filter"
        :filter-included-fields="filterOn"
        primary-key="id"
        :fields="fields"
        :tbody-transition-props="transProps"
        responsive="lg"
        head-variant="light"
        :busy="isBusy"
        class="overflow-auto"
        striped
        hover
        :items="users"
      >
        <template v-slot:table-busy>
          <div class="text-center text-danger my-2">
            <b-spinner class="align-middle"></b-spinner>
            <strong>Loading...</strong>
          </div>
        </template>
        <template v-slot:cell(Avatar)="data">
          <b-avatar
            class="avatar"
            @click="toProfile(data.item.ID)"
            :src="getImgUrl(data.item)"
            variant="primary"
            size="2rem"
            text="EV"
          ></b-avatar>
        </template>
        <template v-slot:cell(name)="data">
          <p class="profile-link" @click="toProfile(data.item.ID)">
            {{ data.item.Name }} {{ data.item.LastName }}
          </p>
        </template>
      </b-table>
    </div>
  </div>
</template>

<script>
export default {
  name: "UserList",
  data() {
    return {
      transProps: {
        // Transition name
        name: "flip-list"
      },
      users: {},
      isBusy: false,
      errors: {},
      filter: null,
      filterOn: [],
      fields: [
        { key: "Avatar" },
        { key: "ID", sortable: true },
        { key: "Name", sortable: true },
        { key: "Email", sortable: true },
        { key: "Role", sortable: true },
        { key: "CreatedAt", sortable: true }
      ]
    };
  },
  mounted() {
    this.loadUsers();
  },
  methods: {
    loadUsers() {
      this.toggleBusy();
      this.errors = [];
      const vue = this;
      window.axios
        .post("/api/admin/users")
        .then(response => {
          this.users = response.data.data;
          this.toggleBusy();
        })
        .catch(function(errors) {
          vue.errors = { error: errors.response.data.error };
        });
    },
    toProfile(id) {
      this.$router.push("/user/" + id + "/profile");
    },
    toggleBusy() {
      this.isBusy = !this.isBusy;
    },
    getImgUrl(item) {
      let images = process.env.VUE_APP_API + "/static" + item.Avatar;
      return images;
    }
  }
};
</script>

<style scoped lang="scss">
.form-row {
  margin-right: 0;
}
.row {
  margin-right: 0;
}
.profile-link {
  color: #00acc1;
  cursor: pointer;
  &:hover {
    color: #af4448;
  }
}
.avatar {
  cursor: pointer;
}

.filter {
  margin-bottom: 20px;
}
</style>
