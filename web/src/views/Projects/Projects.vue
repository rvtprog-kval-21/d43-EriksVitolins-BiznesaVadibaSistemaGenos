<template>
  <b-container class="mt-4">
    <div
      v-if="currentUser.role === 'admin'"
      class="d-flex justify-content-between"
    >
      <div></div>
      <b-button @click="$router.push('/projects/create')" variant="outline-success">Create a new project</b-button>
    </div>
  </b-container>
</template>

<script>
export default {
  name: "Projects",
  computed: {
    currentUser() {
      return this.$store.getters.currentUser;
    }
  },
  methods: {
  }
};
</script>

<style scoped></style>
