<template>
  <b-container class="mt-4">
    <div
      v-if="currentUser.role === 'admin'"
      class="d-flex justify-content-between"
    >
      <div></div>
      <b-button @click="$router.push('/projects/create')" variant="outline-success">Create a new project</b-button>
    </div>
    <div>
      <div class="row-list">
        <div></div>
        <h4 class="temp">Name</h4>
        <h4 class="temp">Users</h4>
      </div>
      <hr>
      <div class="list">
        <template v-for="(iter, index) in this.projects">
            <div :key="index" class="row-list">
              <div>
                <b-avatar size="4rem" :src="getImgUrl(iter.avatar)"></b-avatar>
              </div>
              <div class="name">
                <h4>{{iter.name}}</h4>
              </div>
              <div class="d-flex users">
                <template v-for="(liter,index) in iter.members">
                  <b-avatar :src="getImgUrl(liter.user.avatar)" :key="index"></b-avatar>
                </template>
              </div>
              <div>
                <b-button variant="outline-primary">Visit</b-button>
              </div>
            </div>
          <hr :key="index">
        </template>
      </div>
    </div>
  </b-container>
</template>

<script>
export default {
  name: "Projects",
  computed: {
    currentUser() {
      return this.$store.getters.currentUser;
    }
  },
  data() {
    return {
      projects: [],
    };
  },
  methods: {
    getProjects() {
      window.axios
              .get("api/projects/all/")
              .then(res => {
                this.projects = res.data.projects;
              })
              .catch(rej => {
                console.log(rej.response.data.error);
              });
    },
    getImgUrl(image) {
      let images = process.env.VUE_APP_API + "/static" + image;
      return images;
    },
  },
  created() {
    this.getProjects()
  }
};
</script>

<style scoped lang="scss">
  .container{
    height: 100%;
    .list{
      margin-top: 25px;
    }
  }
  .row-list{
    display: grid;
    align-items: center;
    grid-template-columns: 70px auto auto 50px;
    grid-template-rows: 60px;
    .users{
      overflow-x: auto;
      align-items: center;
    }
    .temp{
      h4{
        color: #9e9e9e;
      }
    }
    .name{
      border-right:solid 1px grey ;
      margin-right: 20px;
    }
  }
</style>
