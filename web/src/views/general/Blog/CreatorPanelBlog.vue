<template>
  <b-container class="mt-3 body">
    <div class="d-flex justify-content-end">
      <a href="/blog"><b-button variant="outline-info">Back</b-button></a>
    </div>
    <hr />
    <div class="mt-3 mb-3">
      <div class="d-flex justify-content-end">
        <b-button v-if="!isCreating || isEditing" @click="isCreating = true" variant="outline-primary">Create a Post</b-button>
        <b-button v-else @click="[isCreating = false, isEditing = false]" variant="outline-primary">Home</b-button>
      </div>
      <template v-if="isCreating">
        <CreateBlog></CreateBlog>
      </template>
      <template v-if="isEditing">
        <CreateBlog :isEdit="true"></CreateBlog>
      </template>
      <template v-if="!isEditing && !isCreating">
        Wow
      </template>
    </div>
  </b-container>
</template>

<script>
  import CreateBlog from "../../../components/blog/CreateBlog";
export default {
  name: "CreatorPanelBlog",
  components: {CreateBlog},
  data() {
    return {
      isCreating: false,
      isEditing: false
    };
  },
};
</script>

<style>
  .container{
    min-height: 800px;
  }
</style>
