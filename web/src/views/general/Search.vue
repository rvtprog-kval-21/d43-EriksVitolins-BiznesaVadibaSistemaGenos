<template>
  <b-container class="mt-3">
    <div>
      <h4>Users: {{ users.length }}</h4>
    </div>
    <hr />
    <template v-for="(iter, index) in users">
      <div class="row-bottom" :key="index">
        <div class="d-flex align-content-center justify-content-between">
          <div class="d-flex align-content-center ">
            <b-avatar size="3rem" :src="getImgUrl(iter.avatar)"></b-avatar>
            <h6 class="m-0">{{ iter.name + " " + iter.last_name }}</h6>
          </div>
          <b-button
            @click="$router.push(`/user/${iter.id}/profile`)"
            variant="outline-primary"
            >Visit</b-button
          >
        </div>
        <div class="border">
          <p class="ml-3 mt-3">{{ iter.about }}</p>
        </div>
        <hr />
      </div>
    </template>
  </b-container>
</template>

<script>
export default {
  name: "Search.vue",
  data() {
    return {
      search: "",
      users: []
    };
  },
  methods: {
    getUsers() {
      window.axios
        .post("api/users/search", { search: this.search })
        .then(res => {
          this.users = res.data.users;
        });
    },
    getImgUrl(avatar) {
      let images = process.env.VUE_APP_API + "/static" + avatar;
      return images;
    }
  },
  created() {
    this.search = this.$route.params.search;
    console.log(this.search);
    this.getUsers();
  }
};
</script>

<style scoped>
.m-0 {
  margin: auto 10px auto 10px !important;
}

.border {
  margin: 10px 0 0 40px;
}
</style>
