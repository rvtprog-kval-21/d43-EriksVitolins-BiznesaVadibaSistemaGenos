export function initialize(store, router) {
  router.beforeEach((to, from, next) => {
    const requiresAuth = to.matched.some(record => record.meta.requireAuth);
    const currentUser = store.state.currentUser;
    if (requiresAuth && !currentUser) {
      next("/");
    } else if (to.path === "/" && currentUser) {
      next("/home");
    } else {
      next();
    }
  });

  window.axios.defaults.baseURL = process.env.VUE_APP_API;

  window.axios.interceptors.request.use(request => {
    if (store.state.currentUser) {
      const token = store.state.currentUser.access_token;
      if (token) {
        request.headers.common["Authorization"] = "Bearer " + token;
      }
      //window.axios.headers.append('Access-Control-Allow-Origin', process.env.VUE_APP_API);
      // window.axios.headers.append('Access-Control-Allow-Credentials', 'true');
    }
    return request;
  });
  window.axios.interceptors.response.use(null, error => {
    if (error.response.status === 401 && store.state.currentUser) {
      store.commit("logout");
      store.commit("notAuthorized");
      router.push("/");
    }

    return Promise.reject(error);
  });
}

export function setAuthorization(token) {
  window.axios.defaults.headers.common.Authorization = `Bearer ${token}`;
  window.axios.defaults.headers.common.Accept = "application/json";
}
