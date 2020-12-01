import Vue from "vue";
import VueRouter from "vue-router";

const Login = () => import("../views/general/Login");
const NotFound = () => import("../views/general/NotFound");

const Blog = () => import("../views/general/Blog/Blog");
const AdminPanelBlog = () => import("../views/general/Blog/AdminPanelBlog");
const CreatorPanelBlog = () => import("../views/general/Blog/CreatorPanelBlog");
const IndividualBlog = () => import("../views/general/Blog/IndividualBlog");

const Tracking = () => import("../views/Tracking/Tracking");
const AdminPanelTracking = () => import("../views/Tracking/AdminPanelTracking");
const ManagerPanelTracking = () => import("../views/Tracking/ManagerPanelTracking");

const Projects = () => import("../views/Projects/Projects");
const CreateProjects = () => import("../views/Projects/CreateProjects");

const Submissions = () => import("../views/general/Submissions");

const ProfileComponent = () => import("../views/Profile/ProfileComponent");
const AddUsers = () => import("../views/Admin/AddUsers");
const UserList = () => import("../views/Admin/UserList");
const Home = () => import("../views/Home");
const TagsList = () => import("../views/Tags/TagsList");
const TagsProfile = () => import("../views/Tags/TagsProfile");
const Timetable = () => import("../views/TimeTable/TimeTables");
Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Login",
    component: Login
  },
  {
    path: "/blog",
    component: Blog
  },
  {
    path: "/blog/admin",
    component: AdminPanelBlog
  },
  {
    path: "/projects",
    component: Projects
  },
  {
    path: "/projects/create",
    component: CreateProjects
  },
  {
    path: "/blog/creator",
    component: CreatorPanelBlog
  },
  {
    path: "/blog/:id/view",
    component: IndividualBlog
  },
  {
    path: "/tracking",
    component: Tracking
  },
  {
    path: "/tracking/admin",
    component: AdminPanelTracking
  },
  {
    path: "/tracking/manager",
    component: ManagerPanelTracking
  },
  {
    path: "/home",
    component: Home
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
    path: "/timetable",
    component: Timetable,
    meta: {
      requireAdmin: true
    }
  },
  {
    path: "/submissions",
    component: Submissions,
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
