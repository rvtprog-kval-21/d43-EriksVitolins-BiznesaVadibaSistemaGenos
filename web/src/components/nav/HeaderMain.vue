<template>
  <div class="header">
    <button class="btn-svg" v-if="!isOpened" @click="isOpened = !isOpened">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="icon icon-tabler icon-tabler-menu-2"
        width="30"
        height="30"
        viewBox="0 0 24 24"
        stroke-width="2"
        stroke="#607D8B"
        fill="none"
        stroke-linecap="round"
        stroke-linejoin="round"
      >
        <path stroke="none" d="M0 0h24v24H0z" />
        <line x1="4" y1="6" x2="20" y2="6" />
        <line x1="4" y1="12" x2="20" y2="12" />
        <line x1="4" y1="18" x2="20" y2="18" />
      </svg>
    </button>
    <div
      class="navbar"
      v-if="(isOpened && smallerThanBreakPoint) || !smallerThanBreakPoint"
    >
      <button
        class="btn-svg"
        v-if="isOpened && smallerThanBreakPoint"
        @click="isOpened = !isOpened"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="icon icon-tabler icon-tabler-menu-2"
          width="30"
          height="30"
          viewBox="0 0 24 24"
          stroke-width="2"
          stroke="white"
          fill="none"
          stroke-linecap="round"
          stroke-linejoin="round"
        >
          <path stroke="none" d="M0 0h24v24H0z" />
          <line x1="4" y1="6" x2="20" y2="6" />
          <line x1="4" y1="12" x2="20" y2="12" />
          <line x1="4" y1="18" x2="20" y2="18" />
        </svg>
      </button>
      <div class="content">
        <router-link
          class="nav-row"
          tag="li"
          active-class="active"
          to="/home"
          exact
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="icon icon-tabler icon-tabler-home-2"
            width="18"
            height="18"
            viewBox="0 0 24 24"
            stroke-width="2"
            stroke="#ffffff"
            fill="none"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <path stroke="none" d="M0 0h24v24H0z" />
            <polyline points="5 12 3 12 12 3 21 12 19 12" />
            <path d="M5 12v7a2 2 0 0 0 2 2h10a2 2 0 0 0 2 -2v-7" />
            <rect x="10" y="12" width="4" height="4" />
          </svg>
          <a>Home</a>
        </router-link>
        <NavMenuDropDown title="Profile">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="icon icon-tabler icon-tabler-user"
            width="18"
            height="18"
            viewBox="0 0 24 24"
            stroke-width="2"
            stroke="#ffffff"
            fill="none"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <path stroke="none" d="M0 0h24v24H0z" />
            <circle cx="12" cy="7" r="4" />
            <path d="M6 21v-2a4 4 0 0 1 4 -4h4a4 4 0 0 1 4 4v2" />
          </svg>
          <div slot="nav-menu-content">
            <router-link
              class="nav-row nav-links"
              tag="li"
              active-class="active"
              :to="'/user/' + currentUser.id"
              exact
            >
              <a>My profile</a>
            </router-link>
            <router-link
              class="nav-row nav-links"
              tag="li"
              active-class="active"
              to="/profile/settings"
              exact
            >
              <a>Settings</a>
            </router-link>
          </div>
        </NavMenuDropDown>
        <AdminNavItems v-if="currentUser.role === 'admin'"></AdminNavItems>
        <NavMenuDropDown title="Tags">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="icon icon-tabler icon-tabler-traffic-cone"
            width="18"
            height="18"
            viewBox="0 0 24 24"
            stroke-width="2"
            stroke="#ffffff"
            fill="none"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <path stroke="none" d="M0 0h24v24H0z" fill="none" />
            <line x1="4" y1="20" x2="20" y2="20" />
            <line x1="9.4" y1="10" x2="14.6" y2="10" />
            <line x1="7.8" y1="15" x2="16.2" y2="15" />
            <path d="M6 20l5 -15h2l5 15" />
          </svg>
          <div slot="nav-menu-content">
            <router-link
              class="nav-row nav-links"
              tag="li"
              active-class="active"
              to="/tags"
              exact
            >
              <a>Tags</a>
            </router-link>
            <router-link
              class="nav-row nav-links"
              tag="li"
              active-class="active"
              to="/locations"
              exact
            >
              <a>Location</a>
            </router-link>
          </div>
        </NavMenuDropDown>
        <router-link
          class="nav-row"
          tag="li"
          active-class="active"
          to="/home"
          exact
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="icon icon-tabler icon-tabler-mail"
            width="18"
            height="18"
            viewBox="0 0 24 24"
            stroke-width="2"
            stroke="#ffffff"
            fill="none"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <path stroke="none" d="M0 0h24v24H0z" />
            <rect x="3" y="5" width="18" height="14" rx="2" />
            <polyline points="3 7 12 13 21 7" />
          </svg>
          <a>Tērzēšana</a>
        </router-link>
        <div class="nav-row">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="icon icon-tabler icon-tabler-logout"
            width="18"
            height="18"
            viewBox="0 0 24 24"
            stroke-width="2"
            stroke="#ffffff"
            fill="none"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <path stroke="none" d="M0 0h24v24H0z" />
            <path
              d="M14 8v-2a2 2 0 0 0 -2 -2h-7a2 2 0 0 0 -2 2v12a2 2 0 0 0 2 2h7a2 2 0 0 0 2 -2v-2"
            />
            <path d="M7 12h14l-3 -3m0 6l3 -3" />
          </svg>
          <a href="#" @click.prevent="logout">Logout</a>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import AdminNavItems from "./AdminNavItems";
import NavMenuDropDown from "./NavMenuDropDown";
import { setAuthorization } from "../../helpers/middleware";

export default {
  name: "HeaderMain",
  components: { NavMenuDropDown, AdminNavItems },
  data() {
    return {
      isOpened: true,
      smallerThanBreakPoint: false
    };
  },
  created() {
    window.addEventListener("resize", this.handleResize);
    this.handleResize();
  },
  methods: {
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
    handleResize() {
      if (window.innerWidth < 1200) {
        this.isOpened = false;
        this.smallerThanBreakPoint = true;
      } else {
        this.isOpened = true;
        this.smallerThanBreakPoint = false;
      }
    }
  },
  computed: {
    currentUser() {
      return this.$store.getters.currentUser;
    }
  }
};
</script>

<style lang="scss">
.navbar {
  width: 300px;
  background-color: #636b6f;
  height: 100%;
}

.content {
  padding: 20px 0px 40px 0px;
}

.header {
  height: 100%;

  .nav-row {
    width: 100%;
    padding-top: 10px;
    padding-bottom: 10px;
    font-size: 18px;
    display: flex;
    align-items: center;

    a {
      color: white;
    }

    &:hover {
      background-color: #9e9e9e;

      a {
        color: #ffc107;
      }
    }

    &.nav-links a {
      margin-left: 40px;
    }
  }

  svg {
    margin-right: 10px;
    margin-left: 10px;
  }

  .header li {
    list-style: none;
  }

  .nav-row.router-link-exact-active.active {
    background-color: #b0bec5;
    cursor: default;
  }

  .nav-menu {
    justify-content: space-between;

    .nav-menu-name {
      display: flex;
    }

    &:hover {
      background-color: #9e9e9e;
      cursor: pointer;

      a {
        color: #ffc107;
      }

      .nav-menu-name {
        display: flex;
      }
    }
  }
}

.btn-svg {
  border: none;
  background-color: inherit;
  cursor: pointer;
}

@media screen and (max-width: 750px) {
  .navbar {
    width: 100%;
  }
}
</style>
