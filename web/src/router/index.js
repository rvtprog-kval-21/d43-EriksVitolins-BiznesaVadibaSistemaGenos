import Vue from "vue";
import VueRouter from "vue-router";

const ProfileComponent = () => import("../views/Profile/ProfileComponent");
const AddUsers = () => import("../views/Admin/AddUsers");
const UserList = () => import("../views/Admin/UserList");
const Login = () => import("../views/Auth/Login");
const Home = () => import("../views/Home");
const NotFound = () => import("../views/404/NotFound");
const TagsList = () => import("../views/Tags/TagsList");
const TagsProfile = () => import("../views/Tags/TagsProfile");

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
    path: "/user/:id/profile",
    component: ProfileComponent,
    meta: {
      requireAdmin: true
    }
  },
  {
    path: "/tags",
    component: TagsList,
    meta: {
      requireAdmin: true
    }
  },
  {
    path: "/tags/:id/tag",
    component: TagsProfile,
    meta: {
      requireAdmin: true
    }
  },
  {
    path: "*",
    component: NotFound
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
