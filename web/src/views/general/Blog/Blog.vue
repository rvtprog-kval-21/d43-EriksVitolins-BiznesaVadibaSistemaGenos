<template>
  <b-container class="mt-3">
    <template v-if="currentUser.role === 'admin' || isCreator">
      <div class="d-flex justify-content-end">
        <a href="/blog/admin">
          <b-button
            class="mr-3"
            v-if="currentUser.role === 'admin'"
            variant="outline-info"
            >Admin Panel</b-button
          >
        </a>
        <a href="/blog/creator"
          ><b-button v-if="isCreator" variant="outline-info"
            >Creator Panel</b-button
          ></a
        >
      </div>
    </template>
    <template>
      <h1>Hello</h1>
    </template>
  </b-container>
</template>

<script>
export default {
  name: "Blog",
  computed: {
    currentUser() {
      return this.$store.getters.currentUser;
    }
  },
  data() {
    return {
      errors: {},
      isCreator: true
    };
  }
};
</script>

<style scoped></style>
