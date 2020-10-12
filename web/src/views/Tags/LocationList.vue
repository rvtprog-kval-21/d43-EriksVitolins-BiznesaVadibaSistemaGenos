<template>
  <div class="init-screen">
    <alertComponent v-if="this.errors.error" class="alert-danger">
      <p v-for="(iter, index) in errors" :key="index">{{ iter.name }}</p>
    </alertComponent>
    <alertComponent v-if="this.alerts.message" class="alert-success">
      <p>{{ this.alerts.message }}</p>
    </alertComponent>
    <div class="container">
      <div class="btns">
        <div class="top-btns">
          <div
            v-if="!isPartOfLocation"
            class="add-btn btn"
            @click="joinLocation"
          >
            <p>Join</p>
          </div>
          <div v-else class="add-btn btn" @click="leave">
            <p>Leave</p>
          </div>
          <div class="add-btn btn" v-if="currentUser.role === 'admin'">
            <p>Add</p>
          </div>
          <div class="kick-btn btn" v-if="currentUser.role === 'admin'">
            <p>Kick</p>
          </div>
        </div>
      </div>
      <div class="table-head grid">
        <div class="">
          <input
            @click="selectAllUsers"
            v-model="allUsersSelected"
            type="checkbox"
          />
        </div>

        <div class=""><p>Nr</p></div>
        <div>Name</div>
      </div>
      <div class="hr"></div>
      <div class="table">
        <div class="row grid" v-for="(iter, index) in users" :key="index">
          <div><input type="checkbox" v-model="iter.isChecked" /></div>
          <div>
            <p>{{ index }}</p>
          </div>
          <div>
            <p>{{ iter.name + " " + iter.lastname }}</p>
          </div>
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
    return {
      errors: [],
      alerts: [],
      location: {},
      users: {},
      allUsersSelected: false,
      isPartOfLocation: false
    };
  },
  mounted() {
    this.loadUsers();
  },
  methods: {
    joinLocation: function() {
      this.errors = {};
      this.alerts = {};
      const vue = this;
      window.axios
        .post("/api/location/join", {
          location: this.$route.params.id,
          user: this.currentUser.id
        })
        .then(res => {
          this.location = res.data.data;
          this.alerts = { message: "Location joined" };
          this.loadUsers();
          this.isPartOfLocation = true;
        })
        .catch(function(rej) {
          vue.errors = { error: rej.response.data.message };
        });
    },
    selectAllUsers: function() {
      this.users.forEach(iter => (iter.isChecked = !this.allUsersSelected));
    },
    loadUsers: function() {
      this.errors = {};
      const vue = this;
      window.axios
        .get("/api/location/" + this.$route.params.id + "/users")
        .then(res => {
          this.users = res.data.data;
          this.isPartOf();
        })
        .catch(function(rej) {
          vue.errors = { error: rej.response.data.message };
        });
    },
    isPartOf: function() {
      for (let iter = 0; iter < this.users.length; iter++) {
        console.log(this.users[iter].id);
        if (this.users[iter].id === this.currentUser.id) {
          this.isPartOfLocation = true;
          return;
        }
      }
    },
    leave: function() {
      this.errors = {};
      this.alerts = {};
      const vue = this;
      window.axios
        .post("/api/location/leave", {
          location: this.$route.params.id,
          user: this.currentUser.id
        })
        .then(() => {
          this.alerts = { message: "Location left" };
          this.isPartOfLocation = false;
          this.loadUsers();
        })
        .catch(function(rej) {
          vue.errors = { error: rej.response.data.message };
        });
    }
  },
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

  .top-btns {
    display: flex;
    justify-content: space-between;
  }

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

  .hr {
    border-bottom: 1px solid grey;
    margin-top: 7px;
    margin-bottom: 7px;
  }

  .grid {
    display: grid;
    grid-template-columns: 60px 60px auto;
    margin-top: 15px;
  }
}
</style>
