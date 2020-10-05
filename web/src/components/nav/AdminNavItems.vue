<template>
  <div>
    <NavMenuDropDown title="Users">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="icon icon-tabler icon-tabler-users"
        width="18"
        height="18"
        viewBox="0 0 24 24"
        stroke-width="2"
        stroke="#ffffff"
        fill="none"
        stroke-linecap="round"
        stroke-linejoin="round"
      >
        <path stroke="none" d="M0 0h24v24H0z" />
        <circle cx="9" cy="7" r="4" />
        <path d="M3 21v-2a4 4 0 0 1 4 -4h4a4 4 0 0 1 4 4v2" />
        <path d="M16 3.13a4 4 0 0 1 0 7.75" />
        <path d="M21 21v-2a4 4 0 0 0 -3 -3.85" />
      </svg>
      <div slot="nav-menu-content">
        <router-link
          class="nav-row nav-links"
          tag="li"
          active-class="active"
          to="/admin/users"
          exact
        >
          <a>List</a>
        </router-link>
        <router-link
          class="nav-row nav-links"
          tag="li"
          active-class="active"
          to="/admin/userAdd"
          exact
        >
          <a>Add</a>
        </router-link>
      </div>
    </NavMenuDropDown>
    <div class="nav-item nav-row">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="icon icon-tabler icon-tabler-bug"
        width="18"
        height="18"
        viewBox="0 0 24 24"
        stroke-width="2"
        stroke="#ffffff"
        fill="none"
        stroke-linecap="round"
        stroke-linejoin="round"
      >
        <path stroke="none" d="M0 0h24v24H0z" />
        <path d="M9 9v-1a3 3 0 0 1 6 0v1" />
        <path d="M8 9h8a6 6 0 0 1 1 3v3a5 5 0 0 1 -10 0v-3a6 6 0 0 1 1 -3" />
        <line x1="3" y1="13" x2="7" y2="13" />
        <line x1="17" y1="13" x2="21" y2="13" />
        <line x1="12" y1="20" x2="12" y2="14" />
        <line x1="4" y1="19" x2="7.35" y2="17" />
        <line x1="20" y1="19" x2="16.65" y2="17" />
        <line x1="4" y1="7" x2="7.75" y2="9.4" />
        <line x1="20" y1="7" x2="16.25" y2="9.4" />
      </svg>
      <div class="name">
        <a href="">Problēmu Bedre</a>
      </div>
    </div>
  </div>
</template>

<script>
import NavMenuDropDown from "./NavMenuDropDown";

export default {
  name: "AdminNavItems",
  components: {
    NavMenuDropDown
  }
};
</script>

<style scoped></style>
