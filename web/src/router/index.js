import Vue from "vue";
import VueRouter from "vue-router";
import Search from "@/views/general/Search";

const Login = () => import("../views/general/Login");
const NotFound = () => import("../views/general/NotFound");

const Blog = () => import("../views/general/Blog/Blog");
const AdminPanelBlog = () => import("../views/general/Blog/AdminPanelBlog");
const CreatorPanelBlog = () => import("../views/general/Blog/CreatorPanelBlog");
const IndividualBlog = () => import("../views/general/Blog/IndividualBlog");

const Tracking = () => import("../views/Tracking/Tracking");
const AdminPanelTracking = () => import("../views/Tracking/AdminPanelTracking");
const ManagerPanelTracking = () =>
  import("../views/Tracking/ManagerPanelTracking");

const Projects = () => import("../views/Projects/Projects");
const CreateProjects = () => import("../views/Projects/CreateProjects");
const ProjectProfile = () => import("../views/Projects/ProjectProfile");
const SettingsProject = () => import("../views/Projects/SettingsProject");

const Submissions = () => import("../views/general/Submissions");

const ProfileComponent = () => import("../views/Profile/ProfileComponent");
const SettingsComponent = () => import("../views/Profile/SettingsComponent");
const AddUsers = () => import("../views/Admin/AddUsers");
const UserList = () => import("../views/Admin/UserList");
const Home = () => import("../views/Home");
const Timetable = () => import("../views/TimeTable/TimeTables");
const Notifications = () => import("../views/Notifications/Notifications");

const Chatting = () => import("../views/Chatting/Chatting");

const Calendar = () => import("../views/Calendar/Calendar");
Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Login",
    component: Login
  },
  {
    path: "/blog",
    component: Blog,
    meta: {
      requireAdmin: true
    }
  },
  {
    path: "/notifications",
    component: Notifications,
    meta: {
      requireAdmin: true
    }
  },
  {
    path: "/chat",
    component: Chatting,
    meta: {
      requireAdmin: true
    }
  },
  {
    path: "/calendar",
    component: Calendar,
    meta: {
      requireAdmin: true
    }
  },
  {
    path: "/search/:search/",
    component: Search,
    meta: {
      requireAdmin: true
    }
  },
  {
    path: "/blog/admin",
    component: AdminPanelBlog,
    meta: {
      requireAdmin: true
    }
  },
  {
    path: "/projects",
    component: Projects,
    meta: {
      requireAdmin: true
    }
  },
  {
    path: "/projects/create",
    component: CreateProjects,
    meta: {
      requireAdmin: true
    }
  },
  {
    path: "/projects/:id/see",
    component: ProjectProfile,
    meta: {
      requireAdmin: true
    }
  },
  {
    path: "/projects/:id/settings",
    component: SettingsProject,
    meta: {
      requireAdmin: true
    }
  },
  {
    path: "/blog/creator",
    component: CreatorPanelBlog,
    meta: {
      requireAdmin: true
    }
  },
  {
    path: "/blog/:id/view",
    component: IndividualBlog,
    meta: {
      requireAdmin: true
    }
  },
  {
    path: "/tracking",
    component: Tracking,
    meta: {
      requireAdmin: true
    }
  },
  {
    path: "/tracking/admin",
    component: AdminPanelTracking,
    meta: {
      requireAdmin: true
    }
  },
  {
    path: "/tracking/manager",
    component: ManagerPanelTracking
  },
  {
    path: "/home",
    component: Home,
    meta: {
      requireAdmin: true
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
    path: "/user/:id/settings",
    component: SettingsComponent,
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
