export function initialize(store, router) {
  router.beforeEach((to, from, next) => {
    if (to.name !== "Login" && !isAuthed(store)) next({ name: "Login" });
    if (to.path === "/" && isAuthed(store)) next("/home");
    else next();
  });
  window.axios.defaults.baseURL = process.env.VUE_APP_API;

  window.axios.interceptors.request.use(request => {
    if (store.state.currentUser) {
      const token = store.state.currentUser.access_token;
      if (token) {
        request.headers.common["Authorization"] = "Bearer " + token;
      }
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
}

function isAuthed(store) {
  if (store.getters.currentUser) {
    return true;
  }
  return false;
}
