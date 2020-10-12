import Vue from "vue";
import VueRouter from "vue-router";
import Login from "../views/Auth/Login";
import Home from "../views/Home";
import AddUsers from "../views/Admin/AddUsers";
//import UserList from "../views/Admin/UserList";
import ProfileComponent from "../views/Profile/ProfileComponent";
import Location from "../views/Tags/Location";
import LocationList from "../views/Tags/LocationList";
import UserListTest from "../views/Admin/UserListTest";

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
    component: UserListTest,
    meta: {
      requireAdmin: true
    }
  },
  {
    path: "/user/:id/profile",
    component: ProfileComponent
  },
  {
    path: "/locations",
    component: Location
  },
  {
    path: "/location/:id",
    component: LocationList
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
