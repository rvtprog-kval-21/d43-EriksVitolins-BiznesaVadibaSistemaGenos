<template>
  <div class="content d-flex mt-2">
    <div class="col-10">
      <div class="top-row">
        <div class="blog pt-3">
          <div class="blog-top pl-3 d-flex flex-column">
            <h5>Blogs:</h5>
          </div>
          <hr />
          <div>
            <div
              class="blog-row d-flex"
              v-for="(iter, index) in this.blogs"
              :key="index"
            >
              <div
                class="ml-2 mr-2 d-flex justify-content-center align-items-center"
              >
                <b-icon
                  v-if="!iter.viewed"
                  font-scale="2"
                  variant="danger"
                  icon="dot"
                ></b-icon>
              </div>
              <div>
                <div
                  @click="$router.push('/blog/' + iter.blog.id + '/view')"
                  class="d-flex align-items-center"
                >
                  <p class="font-weight-bolder">{{ iter.blog.topic }}</p>
                  <b-icon class="mr-2 mb-3 ml-2" icon="dot"></b-icon>
                  <VueShowdown
                    class="mb-0"
                    :markdown="iter.blog.title"
                    flavor="github"
                    :options="{ emoji: true }"
                  ></VueShowdown>
                </div>
                <div>
                  <VueShowdown
                    :markdown="iter.blog.headtext"
                    flavor="github"
                    :options="{ emoji: true }"
                  ></VueShowdown>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
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
      onlineUsers: null,
      blogs: []
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
    },
    getBlogs() {
      window.axios
        .get("api/blog/home/limited")
        .then(res => {
          this.blogs = res.data.blogs;
        })
        .catch(rej => {
          this.makeToast(rej.response.data.error, "danger");
        });
    },
    makeToast(text, variant) {
      this.$bvToast.toast(text, {
        autoHideDelay: 5000,
        variant: variant,
        title: "Notification"
      });
    }
  },
  computed: {
    currentUser() {
      return this.$store.getters.currentUser;
    }
  },
  created() {
    this.getBlogs();
  }
};
</script>

<style scoped lang="scss">
.top-row {
  height: 500px;
  .blog {
    width: 400px;
    background: white;
    height: 100%;
    border-radius: 10px;
    .blog-row {
      cursor: pointer;
      margin-bottom: 20px;
      &:hover {
        background-color: #76d275;
      }
    }
  }
}
</style>
