<template>
  <b-container class="mt-4 mb-4">
    <div class="">
      <b-alert v-if="this.errors.error" variant="danger" show="">{{
        this.errors.error
      }}</b-alert>
    </div>
    <div class="d-flex justify-content-between">
      <template v-if="isPublic">
        <h4>Public Tags</h4>
        <b-button @click="loadTags('private')">Private tags</b-button>
      </template>
      <template v-else>
        <h4>Tags You're a Member of</h4>
        <b-button @click="loadTags('public')">Public tags</b-button>
      </template>
    </div>
    <div class="holder mt-5">
      <b-row class="grid mb-2">
        <div class="ml-5 d-flex justify-content-center ">
          <h5 class="grey">logo</h5>
        </div>
        <div class="ml-3">
          <h5 class="grey">Name</h5>
        </div>
        <div class="ml-3">
          <h5 class="grey">Members</h5>
        </div>
      </b-row>
      <template v-for="iter in tags">
        <b-row :key="iter.id" class="card mr-0 ml-0 mb-4 grid">
          <div class="mt-auto ml-4 mb-auto">
            <b-avatar
              variant="info"
              size="5rem"
              :src="getLogo(iter.avatar)"
              class="name mr-3"
            ></b-avatar>
          </div>
          <div class="d-flex align-items-center ml-3">
            <h5 class="name">
              <a :href="'/tags/' + iter.id + '/tag'"> {{ iter.name }}</a>
            </h5>
          </div>
          <div class="d-flex align-items-center" v-if="tags">
            <div class="ml-2 d-flex">
              <template v-for="(liter, lindex) in iter.members">
                <template v-if="lindex < 3">
                  <a
                    :key="lindex"
                    :href="'/user/' + liter.user.id + '/profile'"
                  >
                    <b-avatar
                      variant="info"
                      size="2rem"
                      :src="getLogo(liter.user.avatar)"
                      :key="lindex"
                      class="name mr-3"
                    ></b-avatar>
                  </a>
                </template>
                <template
                  v-if="
                    lindex == iter.members.length - 1 &&
                      plusMembersCount(iter.members.length) != 0
                  "
                >
                  <p :key="lindex">
                    +{{ plusMembersCount(iter.members.length) }}
                  </p>
                </template>
              </template>
            </div>
          </div>
          <div class="d-flex justify-content-center align-items-center">
            <b-button
              :href="'/tags/' + iter.id + '/tag'"
              variant="outline-primary"
              >Visit</b-button
            >
          </div>
        </b-row>
      </template>
    </div>
    <template v-if="!isBusy && tags"> </template>
    <template v-if="isBusy">
      <div class="d-flex justify-content-center mt-5">
        <b-spinner label="Spinning"></b-spinner>
      </div>
    </template>
  </b-container>
</template>

<script>
export default {
  name: "TagsList",
  data() {
    return {
      errors: {},
      isBusy: false,
      tags: null,
      isPublic: false
    };
  },
  methods: {
    loadTags(url) {
      this.isBusy = !this.isBusy;
      this.errors = [];
      this.tags = null;
      const vue = this;
      this.isPublic = !this.isPublic;
      window.axios
        .get("/api/tags/" + url)
        .then(response => {
          this.tags = response.data.tags;
          this.isBusy = !this.isBusy;
        })
        .catch(function(errors) {
          vue.errors = { error: errors.response.data.error };
        });
    },
    getLogo(item) {
      let images = process.env.VUE_APP_API + "/static" + item;
      return images;
    },
    goToUrl(url) {
      console.log(url);
      this.$router.push(url);
    },
    plusMembersCount(items) {
      if (items < 4) {
        return 0;
      } else {
        return items - 3;
      }
    }
  },
  mounted() {
    this.loadTags("public");
  }
};
</script>

<style scoped>
.card {
  background: white;
  min-height: 100px;
  border-radius: 10px;
}

.grid {
  display: grid;
  grid-template-columns: 120px 150px auto 90px;
}

.name {
  color: cadetblue;
  cursor: pointer;
}

.grey {
  color: #9e9e9e;
}
</style>
