<template>
  <div>
    <b-navbar sticky toggleable="lg" type="dark" variant="info">
      <b-navbar-brand href="/home">Genos</b-navbar-brand>
      <b-button class="button-burger" v-b-toggle.sidebar-left>
        <b-icon icon="justify" variant="white"></b-icon>
      </b-button>
      <b-sidebar
        id="sidebar-left"
        bg-variant="dark"
        text-variant="light"
        title="Genos"
        left
      >
        <div class="py-2">
          <SidebarContent></SidebarContent>
        </div>
      </b-sidebar>

      <b-navbar-toggle class="button-burger" target="nav-collapse">
        <template v-slot:default="{ expanded }">
          <b-icon v-if="expanded" icon="chevron-bar-up"></b-icon>
          <b-icon v-else icon="chevron-bar-down"></b-icon>
        </template>
      </b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav>
        <b-navbar-nav class="ml-auto">
          <b-nav-form>
            <b-form-input
              v-model="searchTerm"
              size="sm"
              class="mr-sm-2"
              placeholder="Search"
            ></b-form-input>
            <b-button
              @click="$router.push('/search/' + searchTerm + '/')"
              size="sm"
              class="my-2 my-sm-0"
              type="submit"
              >Search</b-button
            >
          </b-nav-form>
        </b-navbar-nav>
      </b-collapse>
      <b-button
        @click="goToNotifications()"
        class="button-burger"
        title="Notifications"
      >
        <b-icon icon="bell" variant="white"> </b-icon>
        <b-badge v-if="notifications != 0" pill variant="danger">{{
          notifications
        }}</b-badge>
      </b-button>
      <b-button
        v-b-tooltip.hover
        title="Settings"
        @click="settings"
        class="button-burger"
      >
        <b-icon icon="gear" variant="white"></b-icon>
      </b-button>
      <b-button
        v-b-tooltip.hover
        title="Profile"
        @click="profile"
        class="button-burger"
      >
        <b-icon icon="person" variant="white"></b-icon>
      </b-button>
      <b-button
        v-b-tooltip.hover
        title="Logout"
        @click="logout"
        class="button-burger"
      >
        <b-icon icon="chevron-bar-right" variant="white"></b-icon>
      </b-button>
    </b-navbar>
  </div>
</template>

<script>
import SidebarContent from "./SidebarContent";
import { setAuthorization } from "../../helpers/middleware";
export default {
  components: { SidebarContent },
  name: "HeaderMain",
  data() {
    return {
      notifications: 0,
      searchTerm: ""
    };
  },
  methods: {
    goToNotifications() {
      this.$router.push("/notifications");
    },
    getNotificationCount() {
      window.axios
        .get(process.env.VUE_APP_API + "/api/notifications/get/count/current")
        .then(response => {
          console.log(response.data);
          this.notifications = response.data.count;
        })
        .catch(error => {
          console.log(error);
        });
    },
    logout() {
      window.axios
        .get(process.env.VUE_APP_API + "/api/auth/logout")
        .then(response => {
          setAuthorization(response.data.access_token);
        })
        .catch(error => {
          console.log(error);
        });
      this.$store.commit("logout");
      this.$router.push("/");
    },
    profile() {
      this.$router.push("/user/" + this.currentUser.id + "/profile");
    },
    settings() {
      this.$router.push("/user/" + this.currentUser.id + "/settings");
    }
  },
  computed: {
    currentUser() {
      return this.$store.getters.currentUser;
    }
  },
  created() {
    this.getNotificationCount();
  }
};
</script>

<style scoped lang="scss">
.button-burger {
  background-color: #17a2b8;
  border: none;
  border-radius: 0;
}
.button-burger:hover {
  background-color: #17a2b8;
  svg {
    color: #fc2f05 !important;
  }
}
.navbar-toggler {
  border: 0;
  outline: none !important;
  box-shadow: none;
  svg {
    color: white;
  }
}
</style>
