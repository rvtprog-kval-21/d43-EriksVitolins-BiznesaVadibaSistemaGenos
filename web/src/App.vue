<template>
  <div id="app" class="">
    <link
      href="https://fonts.googleapis.com/css2?family=Noto+Sans+TC:wght@300&display=swap"
      rel="stylesheet"
    />
    <div id="nav" v-if="currentUser">
      <HeaderTest></HeaderTest>
    </div>
    <router-view />
  </div>
</template>

<script>
//import HeaderMain from "./components/nav/HeaderMain";
import HeaderTest from "./components/nav/HeaderTest";
export default {
  name: "App",
  components: { HeaderTest },
  computed: {
    currentUser() {
      return this.$store.getters.currentUser;
    }
  }
};
</script>

<style lang="scss">
html,
body {
  margin: 0;
  height: 100%;
}

p {
  margin: 0;
}

a {
  text-decoration: none;
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
  background-attachment: fixed;
  background-size: cover;
  background-repeat: repeat;
  background-position: center center;
  display: flex;
  height: 100%;
  background-color:#eceff1;
}

#nav {
  width: 100%;
  height: 70px;
}
@media screen and (min-width: 1000px) {
  #app {
    height: 100%;
  }
}
@media screen and (max-width: 750px) {
  #app {
    flex-direction: column;
  }
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
