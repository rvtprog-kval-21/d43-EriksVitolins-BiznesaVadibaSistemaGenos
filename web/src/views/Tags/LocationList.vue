<template>
  <div class="init-screen">
    <alertComponent v-if="this.errors.error" class="alert-danger">
      <p v-for="(iter, index) in errors" :key="index">{{ iter.name }}</p>
    </alertComponent>
    <alertComponent v-if="this.alerts.message" class="alert-success">
      <p>{{ this.alerts.message }}</p>
    </alertComponent>
    <div class="container">
      <div class="top-btns" v-if="currentUser.role === 'admin'">
        <div class="add-btn btn">
          <p>Add</p>
        </div>
        <div class="kick-btn btn">
          <p>Kick</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import alertComponent from "../../components/alerts/alert-component";
export default {
  name: "LocationList",
  components: {
    alertComponent
  },
  data() {
    return {};
  },
  mounted() {
    this.loadLocations();
  },
  methods: {},
  computed: {
    currentUser() {
      return this.$store.getters.currentUser;
    }
  }
};
</script>

<style scoped lang="scss">
.init-screen {
  padding: 40px 50px;
  width: 100%;

  .btn {
    display: flex;
    justify-content: center;
    width: 100px;
    padding-top: 10px;
    padding-bottom: 10px;
    background-color: #e57373;
    border: none;
    color: white;
    cursor: pointer;

    p {
      color: white;
    }
    &:hover {
      background-color: #af4448;
    }
  }
}
</style>
