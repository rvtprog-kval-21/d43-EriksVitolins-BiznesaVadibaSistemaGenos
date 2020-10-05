import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import { initialize } from "./helpers/middleware";

require("./bootstrap");
Vue.component("pagination", require("laravel-vue-pagination"));
initialize(store, router);

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
