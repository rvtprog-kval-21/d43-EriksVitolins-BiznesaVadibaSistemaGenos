<template>
  <div class="init-screen">
    <alertComponent v-if="this.errors.error" class="alert-danger">
      <p v-for="(iter, index) in errors" :key="index">{{ iter.name }}</p>
    </alertComponent>
    <alertComponent v-if="this.alerts.message" class="alert-success">
      <p>{{ this.alerts.message }}</p>
    </alertComponent>
    <div
      v-if="addIsOpened === false && editIsOpened === false"
      class="container"
    >
      <div class="top-btns" v-if="currentUser.role === 'admin'">
        <div @click="addIsOpened = !addIsOpened" class="add-btn btn">
          <p>Add</p>
        </div>
        <div @click="deleteLocations" class="delete-btn btn">
          <p>Delete</p>
        </div>
        <div @click="editIsOpened = !editIsOpened" class="edit-btn btn">
          <p>Edit</p>
        </div>
      </div>

      <div class="top-info">
        <div class="selectAll" v-if="currentUser.role === 'admin'">
          <input
            @click="selecteAllLocations"
            v-model="allLocationsSelected"
            type="checkbox"
          />
        </div>
        <div class="nr">
          <p>Nr</p>
        </div>
        <div class="name">
          <p>Location</p>
        </div>
        <div class="User-count">
          <p>User Count</p>
        </div>
      </div>
      <div class="hr"></div>

      <div
        class="table-row grid-table"
        v-for="(iter, index) in locations"
        :key="index"
      >
        <input
          type="checkbox"
          v-model="iter.isChecked"
          v-if="currentUser.role === 'admin'"
        />
        <p>{{ iter.id }}</p>
        <p class="link" @click="goToLocation(iter.id)">{{ iter.Name }}</p>
        <p>0</p>
      </div>
    </div>
    <div v-if="addIsOpened" class="Add-container">
      <div class="header-module">
        <div class=""></div>
        <svg
          @click="addIsOpened = !addIsOpened"
          xmlns="http://www.w3.org/2000/svg"
          class="icon icon-tabler icon-tabler-arrow-left"
          width="36"
          height="36"
          viewBox="0 0 24 24"
          stroke-width="2"
          stroke="#3F51B5"
          fill="none"
          stroke-linecap="round"
          stroke-linejoin="round"
        >
          <path stroke="none" d="M0 0h24v24H0z" fill="none" />
          <line x1="5" y1="12" x2="19" y2="12" />
          <line x1="5" y1="12" x2="11" y2="18" />
          <line x1="5" y1="12" x2="11" y2="6" />
        </svg>
      </div>
      <div class="container">
        <div class="box">
          <p>Location:</p>
          <input v-model="newLocation" type="text" />
          <div @click="addNewLocation" class="btn btn-submit">
            <p>Submit</p>
          </div>
        </div>
      </div>
    </div>

    <div v-if="editIsOpened" class="Edit-container">
      <div class="header-module">
        <div class=""></div>
        <svg
          @click="editIsOpened = !editIsOpened"
          xmlns="http://www.w3.org/2000/svg"
          class="icon icon-tabler icon-tabler-arrow-left"
          width="36"
          height="36"
          viewBox="0 0 24 24"
          stroke-width="2"
          stroke="#3F51B5"
          fill="none"
          stroke-linecap="round"
          stroke-linejoin="round"
        >
          <path stroke="none" d="M0 0h24v24H0z" fill="none" />
          <line x1="5" y1="12" x2="19" y2="12" />
          <line x1="5" y1="12" x2="11" y2="18" />
          <line x1="5" y1="12" x2="11" y2="6" />
        </svg>
      </div>

      <div class="edit-box" v-for="(iter, index) in locations" :key="index">
        <div v-if="iter.isChecked" class="holder">
          <p>{{ index }}:</p>
          <input v-model="iter.Name" :placeholder="iter.Name" type="text" />
        </div>
      </div>
      <div class="btn btn-save" @click="saveEditedLocations">
        <p>Save</p>
      </div>
    </div>
  </div>
</template>

<script>
import alertComponent from "../../components/alerts/alert-component";
export default {
  name: "Location",
  components: {
    alertComponent
  },
  data() {
    return {
      addIsOpened: false,
      editIsOpened: false,
      newLocation: "",
      errors: {},
      alerts: {},
      locations: [],
      allLocationsSelected: false,
      locationsSelected: []
    };
  },
  mounted() {
    this.loadLocations();
  },
  methods: {
    goToLocation: function(id) {
      this.$router.push("/location/" + id);
    },
    selecteAllLocations: function() {
      this.locations.forEach(
        iter => (iter.isChecked = !this.allLocationsSelected)
      );
    },
    addNewLocation: function() {
      this.errors = {};
      this.alerts = {};
      const vue = this;
      window.axios
        .post("/api/locationAdd", { name: this.newLocation })
        .then(res => {
          this.alerts = { message: res.data.message };
          this.newLocation = "";
          this.addIsOpened = false;
          this.loadLocations()
        })
        .catch(function(rej) {
          if (rej.response.data.errors) {
            vue.errors = { error: rej.response.data.errors };
          } else {
            vue.errors = { error: rej.response.data.message };
          }
        });
    },
    loadLocations() {
      this.errors = [];
      const vue = this;
      window.axios
        .get("/api/locations")
        .then(response => {
          this.locations = response.data.data;
        })
        .catch(function(errors) {
          vue.errors = { error: errors.response.data.message };
        });
    },
    saveEditedLocations() {
      this.alerts = [];
      this.errors = [];
      const vue = this;
      for (let i = 0; i < this.locations.length; i++) {
        if (this.locations[i].isChecked) {
          window.axios
            .post("/api/locations/edit", {
              name: this.locations[i].Name,
              id: this.locations[i].id
            })
            .then(() => {
              this.alerts = { message: "Changes made successfully" };
              this.locations[i].isChecked = false;
              this.editIsOpened = false;
            })
            .catch(function(errors) {
              vue.errors = { error: errors.response.data.errors };
            });
        }
      }
      this.loadLocations();
    },
    deleteLocations() {
      this.alerts = [];
      this.errors = [];
      const vue = this;
      for (let i = 0; i < this.locations.length; i++) {
        if (this.locations[i].isChecked) {
          window.axios
            .post("/api/locations/delete", { id: this.locations[i].id })
            .then(() => {
              this.alerts = { message: "Changes made successfully" };
            })
            .catch(function(errors) {
              vue.errors = { error: errors.response.data.errors };
            });
        }
      }
      this.loadLocations();
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

  .hr {
    border-bottom: 1px solid grey;
    margin-top: 5px;
    opacity: 55%;
  }
  .top-btns {
    display: flex;
    justify-content: space-between;
    margin-bottom: 30px;
  }

  .link{
    cursor: pointer;
    color: #1976d2;

    &:hover {
      color: #004ba0;
      text-decoration: none;
    }
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

  .top-info,
  .grid-table {
    display: grid;
    grid-template-columns: 50px 50px auto 100px;
  }

  .table-row {
    margin-top: 5px;
  }

  .header-module {
    display: flex;
    justify-content: space-between;
    svg:hover {
      stroke: red;
      cursor: pointer;
    }
  }

  .Add-container {
    .box {
      display: flex;
      flex-direction: column;
      align-items: center;
      height: 120px;
      justify-content: space-between;
    }
  }

  .Edit-container {
    .edit-box {
      margin-bottom: 20px;
      .holder {
        display: flex;
        p {
          margin-right: 5px;
        }
      }
    }
  }
}
@media only screen and (max-width: 600px) {
  .init-screen {
    padding: 0;
  }
}
</style>
