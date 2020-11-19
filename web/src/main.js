import Vue from "vue";
import AppComp from "./App.vue";
import router from "./router";
import store from "./store";
import { initialize } from "./helpers/middleware";
import { BootstrapVue, IconsPlugin } from "bootstrap-vue";
import "bootstrap/dist/css/bootstrap.css";
import "bootstrap-vue/dist/bootstrap-vue.css";
import VueGoodTablePlugin from "vue-good-table";

// import the styles
import "vue-good-table/dist/vue-good-table.css";

Vue.use(VueGoodTablePlugin);
Vue.use(BootstrapVue);
Vue.use(IconsPlugin);
require("./bootstrap");
initialize(store, router);

new Vue({
  router,
  store,
  render: h => h(AppComp)
}).$mount("#app");
