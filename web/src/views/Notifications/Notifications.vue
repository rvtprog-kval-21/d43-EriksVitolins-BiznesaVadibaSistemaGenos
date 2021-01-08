<template>
  <b-container>
    <h4 class="mt-3 mb-4">Notifications:</h4>
    <div>
      <div class="mb-4" v-for="(iter, index) in notifications" :key="index">
        <div>
          <div class="d-flex justify-content-between">
            <div class="d-flex avi">
              <b-avatar
                size="3rem"
                :src="getImgUrl(iter.author.avatar)"
              ></b-avatar>
              <p class="ml-3">
                {{ iter.author.name + " " + iter.author.last_name }}
              </p>
              <p class="ml-4 grey mb-0">{{ iter.topic }}</p>
              <p class="ml-4 grey mb-0">{{ getDate(iter.published) }}</p>
            </div>
            <b-icon
              @click="dismissNotification(iter.id)"
              scale="2"
              icon="x-circle"
              class="icon-button"
            ></b-icon>
          </div>
          <VueShowdown
            class="border mt-2"
            :markdown="iter.content"
            flavor="github"
            :options="{ emoji: true }"
          ></VueShowdown>
        </div>
      </div>
    </div>
  </b-container>
</template>

<script>
export default {
  name: "Notifications",
  data() {
    return {
      notifications: []
    };
  },
  methods: {
    dismissNotification(id) {
      window.axios
        .get(
          process.env.VUE_APP_API +
            `/api/notifications/update/${id}/notification/seen`
        )
        .then(() => {
          this.makeToast("Notification dismissed", "success")
          this.getNotifications();
        })
        .catch(error => {
          console.log(error);
        });
    },
    makeToast(text, variant) {
      this.$bvToast.toast(text, {
        autoHideDelay: 5000,
        variant: variant,
        title: "Notification"
      });
    },
    getNotifications() {
      window.axios
        .get(
          process.env.VUE_APP_API + "/api/notifications/get/existing/current"
        )
        .then(response => {
          this.notifications = response.data.notifications;
        })
        .catch(error => {
          console.log(error);
        });
    },
    getDate(date) {
      date = new Date(date);
      const day = date.getDate();
      const year = date.getFullYear();
      const month = date.toLocaleString("default", { month: "long" });
      return day + " " + month + " " + year;
    },
    getImgUrl(image) {
      let images = process.env.VUE_APP_API + "/static" + image;
      return images;
    }
  },
  created() {
    this.getNotifications();
  }
};
</script>

<style scoped>
.avi {
  align-items: center;
}
.grey {
  color: #9e9e9e;
}
.icon-button:hover {
  cursor: pointer;
  color: #dc3545 !important;
}
</style>
