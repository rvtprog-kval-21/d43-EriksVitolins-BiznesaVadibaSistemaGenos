import Vue from "vue";
import VueRouter from "vue-router";
import Login from "../views/Auth/Login";
import Home from "../views/Home";
import AddUsers from "../views/Admin/AddUsers";
import UserList from "../views/Admin/UserList";
import ProfileComponent from "../views/Profile/ProfileComponent";
import Location from "../views/Tags/Location";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    component: Login
  },
  {
    path: "/home",
    component: Home,
    meta: {
      requireAuth: true
    }
  },
  {
    path: "/admin/userAdd",
    component: AddUsers,
    meta: {
      requireAdmin: true
    }
  },
  {
    path: "/admin/users",
    component: UserList,
    meta: {
      requireAdmin: true
    }
  },
  {
    path: "/user/:id",
    component: ProfileComponent
  },
  {
    path: "/locations",
    component: Location
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
