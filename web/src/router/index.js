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
const ManagerPanelTracking = () =>
  import("../views/Tracking/ManagerPanelTracking");

const Projects = () => import("../views/Projects/Projects");
const CreateProjects = () => import("../views/Projects/CreateProjects");
const ProjectProfile = () => import("../views/Projects/ProjectProfile");
const SettingsProject = () => import("../views/Projects/SettingsProject");

const Submissions = () => import("../views/general/Submissions");

const ProfileComponent = () => import("../views/Profile/ProfileComponent");
const AddUsers = () => import("../views/Admin/AddUsers");
const UserList = () => import("../views/Admin/UserList");
const Home = () => import("../views/Home");
const Timetable = () => import("../views/TimeTable/TimeTables");
const Notifications = () => import("../views/Notifications/Notifications");

const Calendar = () => import("../views/Calendar/Calendar")
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
    path: "/notifications",
    component: Notifications
  },
  {
    path: "/calendar",
    component: Calendar
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
    path: "/projects/:id/see",
    component: ProjectProfile
  },
  {
    path: "/projects/:id/settings",
    component: SettingsProject
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
