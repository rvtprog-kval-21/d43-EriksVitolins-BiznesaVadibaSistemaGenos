<template>
  <div id="app" class="">
    <link
      href="https://fonts.googleapis.com/css2?family=Noto+Sans+TC:wght@300&display=swap"
      rel="stylesheet"
    />
    <div id="nav" v-if="currentUser">
      <HeaderMain></HeaderMain>
    </div>
    <router-view />
  </div>
</template>

<script>
import HeaderMain from "./components/nav/HeaderMain";
export default {
  name: "App",
  components: { HeaderMain },
  computed: {
    currentUser() {
      return this.$store.getters.currentUser;
    }
  },
  methods: {
    pingServer() {
      let currentUser = this.currentUser;
      if (currentUser.id) {
        console.log(1);
        window.axios.get("/api/ping");
      }
      setInterval(function() {
        if (currentUser.id) {
          window.axios.get("/api/ping");
        }
      }, 180 * 1000);
    }
  },
  mounted() {
    this.pingServer();
  }
};
</script>

<style lang="scss">
html,
body {
  margin: 0;
  height: 100%;
  background-attachment: fixed;
  background-size: cover;
  background-repeat: repeat;
  background-position: center center;
  background-color: #eceff1;
  height: 100%;
}

p {
  margin: 0;
}

a {
  text-decoration: none;
}

.ql-syntax {
  background-color: #23241f;
  color: #f8f8f2;
  overflow: visible;
}

hr {
  border: none;
}

.wrapper {
  height: 100%;
  width: 100%;
  display: flex;
}

.container {
  width: 100%;
  overflow: auto;
}

#app {
  background-color: #eceff1;
  height: 100%;
}

#nav {
  width: 100%;
}

html,
body,
p,
a,
li,
ul,
div,
button {
  font-family: "Noto Sans TC", sans-serif;
}
</style>
