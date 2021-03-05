<template>
  <div>
    <b-row class="text-center">
      <b-col>
        <b-avatar size="6rem" variant="info" :src="getImgUrl()"></b-avatar>
      </b-col>
      <div class="w-100"></div>
      <b-col class="link" @click="profile"
        ><p>{{ currentUser.name + " " + currentUser.lastname }}</p></b-col
      >
    </b-row>
    <b-row class="row-link">
      <b-col>
        <router-link
          class="nav-row"
          tag="li"
          active-class="active"
          to="/home"
          exact
        >
          <a>Home</a>
        </router-link>
      </b-col>
    </b-row>
    <b-row class="row-link">
      <b-col>
        <router-link
          class="nav-row"
          tag="li"
          active-class="active"
          to="/blog"
          exact
        >
          <a>Blog</a>
        </router-link>
      </b-col>
    </b-row>
    <b-row class="row-link">
      <b-col>
        <router-link
          class="nav-row"
          tag="li"
          active-class="active"
          to="/tracking"
          exact
        >
          <a>Tracking</a>
        </router-link>
      </b-col>
    </b-row>
    <b-row class="row-link">
      <b-col>
        <router-link
          class="nav-row"
          tag="li"
          active-class="active"
          to="/projects"
          exact
        >
          <a>Projects</a>
        </router-link>
      </b-col>
    </b-row>
    <b-row class="row-link">
      <b-col>
        <router-link
          class="nav-row"
          tag="li"
          active-class="active"
          to="/calendar"
          exact
        >
          <a>Calendar</a>
        </router-link>
      </b-col>
    </b-row>
    <b-row class="row-link">
      <b-col>
        <router-link
          class="nav-row"
          tag="li"
          active-class="active"
          to="/chat/0"
          exact
        >
          <a>Chat <span v-if="counts !== 0">{{counts}}</span></a>
        </router-link>
      </b-col>
    </b-row>
    <b-row class="row-link">
      <b-col>
        <router-link
          class="nav-row"
          tag="li"
          active-class="active"
          to="/timetable"
          exact
        >
          <a>Timetable</a>
        </router-link>
      </b-col>
    </b-row>
    <b-row class="row-link" v-if="currentUser.role === 'admin'">
      <b-col class="nav-menu" v-b-toggle:Users>
        <div
          class="d-flex justify-content-between align-items-center nav-menu-inside"
        >
          <a href="#">Users</a>
          <b-icon class="when-closed" icon="plus-circle"></b-icon>
          <b-icon class="when-open" icon="x"></b-icon>
        </div>
      </b-col>
    </b-row>
    <b-row class="row-link child-level">
      <b-collapse id="Users" class="collapse-comp">
        <b-row class="row-link">
          <b-col>
            <router-link
              class="nav-row"
              tag="li"
              active-class="active"
              to="/admin/userAdd"
              exact
            >
              <a>Add</a>
            </router-link>
          </b-col>
        </b-row>
      </b-collapse>
    </b-row>
    <b-row class="row-link child-level">
      <b-collapse id="Users" class="collapse-comp">
        <b-row class="row-link">
          <b-col>
            <router-link
              class="nav-row"
              tag="li"
              active-class="active"
              to="/admin/users"
              exact
            >
              <a>List</a>
            </router-link>
          </b-col>
        </b-row>
      </b-collapse>
    </b-row>
  </div>
</template>

<script>
export default {
  name: "SidebarContent",
  data() {
    return {
      image: "",
      counts: 0
    };
  },
  mounted() {
    this.getImgUrl();
    this.GetUnreadCount()
  },
  methods: {
    getImgUrl() {
      let images =
        process.env.VUE_APP_API + "/static" + this.currentUser.avatar;
      return images;
    },
    profile() {
      this.$router.push("/user/" + this.currentUser.id + "/profile");
    },
    GetUnreadCount() {
      window.axios
          .get("api/chatting/get/unread/count")
          .then(res => {
            this.counts = res.data.count
          });
    },
  },
  computed: {
    currentUser() {
      return this.$store.getters.currentUser;
    }
  },
};
</script>

<style scoped lang="scss">
.link:hover {
  cursor: pointer;
  p {
    color: #00acc1;
  }
}

.collapse-comp {
  width: 100%;
}

.collapsed {
  .when-open {
    display: none;
  }
}
.not-collapsed {
  .when-closed {
    display: none;
  }
}

.row {
  margin: 0;
}
.row-link {
  width: 100%;
  font-size: 18px;
  .nav-menu {
    padding-top: 10px;
    padding-bottom: 10px;
  }
  .col {
    padding-left: 0;
    padding-right: 0;
    width: 100%;
    .active {
      background-color: #b0bec5;
    }
    .nav-row {
      padding-top: 10px !important;
      padding-bottom: 10px !important;
    }
  }
  li {
    list-style: none;
  }
  a {
    color: white;
    text-decoration: none;
    margin-left: 20px;
  }
  &.child-level {
    background-color: #607d8b;
    a {
      margin-left: 40px;
    }
  }
  &:hover {
    background-color: #9e9e9e;

    a {
      color: #ffc107;
    }
  }
  span{
    background: red;
    padding: 0px 5px;
  }
}
.nav-menu-inside {
  svg {
    margin-right: 20px;
  }
}
</style>
