<template>
  <div class="nav-dropdown">
    <div @click="toogle" class="nav-menu nav-row">
      <div class="head">
        <slot></slot>
        <a href="#">{{ title }}</a>
      </div>
      <svg
        xmlns="http://www.w3.org/2000/svg"
        v-if="!this.isOpened"
        class="icon icon-tabler icon-tabler-chevron-right"
        width="20"
        height="20"
        viewBox="0 0 24 24"
        stroke-width="2"
        stroke="#ffffff"
        fill="none"
        stroke-linecap="round"
        stroke-linejoin="round"
      >
        <path stroke="none" d="M0 0h24v24H0z" />
        <polyline points="9 6 15 12 9 18" />
      </svg>
      <svg
        xmlns="http://www.w3.org/2000/svg"
        v-if="this.isOpened"
        class="icon icon-tabler icon-tabler-chevron-down"
        width="20"
        height="20"
        viewBox="0 0 24 24"
        stroke-width="2"
        stroke="#ffffff"
        fill="none"
        stroke-linecap="round"
        stroke-linejoin="round"
      >
        <path stroke="none" d="M0 0h24v24H0z" />
        <polyline points="6 9 12 15 18 9" />
      </svg>
    </div>
    <div class="nav-menu-content" v-if="this.isOpened">
      <slot name="nav-menu-content"></slot>
    </div>
  </div>
</template>

<script>
export default {
  name: "NavMenuDropDown",
  props: ["title"],
  data() {
    return {
      isOpened: false
    };
  },
  methods: {
    toogle() {
      this.isOpened = !this.isOpened;
    }
  }
};
</script>

<style scoped>
.nav-menu-content {
  background-color: #607d8b;
}

.head {
  display: flex;
  align-items: center;
}
</style>
