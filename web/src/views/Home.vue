<template>
  <div class="content d-flex mt-2">
    <div class="col-10">
      <p>Hello</p>
    </div>
    <div class="col-2">
      <div class="p-5">
        <template v-for="(iter, index) in onlineUsers">
          <div :key="index" class="profile mb-2 d-flex align-items-center">
            <b-avatar
              size="3rem"
              badge
              badge-variant="success"
              :src="getImgUrl(iter.avatar)"
            ></b-avatar>
            <h6 class="ml-3">{{ iter.name + " " + iter.last_name }}</h6>
          </div>
        </template>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Home",
  data() {
    return {
      onlineUsers: null
    };
  },
  mounted() {
    this.getOnlineUsers();
  },
  methods: {
    getOnlineUsers() {
      const vue = this;
      window.axios.get("/api/usersonline").then(res => {
        vue.onlineUsers = res.data.users;
      });
      setInterval(() => {
        window.axios.get("/api/usersonline").then(res => {
          vue.onlineUsers = res.data.users;
        });
      }, 60 * 1000);
    },
    getImgUrl(avatar) {
      let images = process.env.VUE_APP_API + "/static" + avatar;
      return images;
    },
    profile() {
      this.$router.push("/user/" + this.currentUser.id + "/profile");
    }
  },
  computed: {
    currentUser() {
      return this.$store.getters.currentUser;
    }
  }
};
</script>

<style scoped></style>
