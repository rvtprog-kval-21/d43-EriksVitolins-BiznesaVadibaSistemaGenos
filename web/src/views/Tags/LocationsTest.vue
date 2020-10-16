<template>
  <b-container>
    <div class="" v-if="addIsOpened === false && editIsOpened === false">
      <b-row>
        <b-col class="alarms">
          <div class="danger" v-if="this.errors.error">
            <b-alert
              show
              variant="danger"
              v-for="(iter, index) in errors"
              :key="index"
            >
              <p>{{ iter.name }}</p>
            </b-alert>
          </div>
          <div class="alert">
            <b-alert show variant="success" v-if="this.alerts.message">
              <p>{{ this.alerts.message }}</p>
            </b-alert>
          </div>
        </b-col>
      </b-row>
      <b-container class="bv-example-row mb-4">
        <b-row>
          <b-col
            ><b-button @click="addIsOpened = !addIsOpened">Add</b-button></b-col
          >
          <b-col><b-button @click="deleteLocations">Delete</b-button></b-col>
          <b-col
            ><b-button @click="editIsOpened = !editIsOpened"
              >Edit</b-button
            ></b-col
          >
        </b-row>
      </b-container>
      <b-row class="mb-4">
        <b-col>
          <b-form-group
            label="Filter"
            label-cols-sm="3"
            label-align-sm="right"
            label-size="sm"
            label-for="filterInput"
            class="mb-0"
          >
            <b-input-group size="sm">
              <b-form-input
                v-model="filter"
                type="search"
                id="filterInput"
                placeholder="Type to Search"
              ></b-form-input>
              <b-input-group-append>
                <b-button :disabled="!filter" @click="filter = ''"
                  >Clear</b-button
                >
              </b-input-group-append>
            </b-input-group>
          </b-form-group>
        </b-col>
      </b-row>
      <b-row>
        <b-col class="tables ml-2 mr-2">
          <b-table
            striped
            responsive="lg"
            hover
            :items="locations"
            :fields="fields"
            :sort-by.sync="sortBy"
            :sort-desc.sync="sortDesc"
            :filter="filter"
            :filter-included-fields="filterOn"
            primary-key="id"
            head-variant="light"
          >
            <template v-slot:head(isChecked)="">
              <input
                type="checkbox"
                @click="selecteAllLocations"
                v-model="allLocationsSelected"
              />
            </template>
            <template v-slot:cell(isChecked)="data">
              <input type="checkbox" v-model="data.item.isChecked" />
            </template>
            <template v-slot:cell(name)="data">
              <p @click="goToLocation(data.item.id)" class="link-profile">
                {{ data.item.Name }}
              </p>
            </template>
          </b-table>
        </b-col>
      </b-row>
    </div>
    <b-row v-if="addIsOpened" class="justify-content-end">
      <b-icon
        @click="addIsOpened = !addIsOpened"
        class="button-back mb-5 h1"
        icon="arrow-left"
      ></b-icon>
      <div class="vw-100"></div>
      <b-col sm="4 d-flex justify-content-end align-items-center">
        <label>Name:</label>
      </b-col>
      <b-col sm="4">
        <b-form-input v-model="newLocation" :type="text"></b-form-input>
      </b-col>
      <b-col sm="4">
        <b-button @click="addNewLocation" variant="outline-primary"
          >Submit</b-button
        >
      </b-col>
    </b-row>
    <b-row v-if="editIsOpened">
      <b-icon
        @click="editIsOpened = !editIsOpened"
        class="button-back mb-5 h1"
        icon="arrow-left"
      ></b-icon>
      <div class="vw-100"></div>
      <b-col
        cols="6"
        class="edit-box"
        v-for="(iter, index) in locations"
        :key="index"
      >
        <div v-if="iter.isChecked" class="holder">
          <p>{{ index }}:</p>
          <b-form-input v-model="iter.Name" :type="text"></b-form-input>
        </div>
      </b-col>
      <b-button @click="saveEditedLocations" variant="outline-primary"
        >Save</b-button
      >
    </b-row>
  </b-container>
</template>

<script>
export default {
  name: "LocationsTest.vue",
  data() {
    return {
      fields: [
        "isChecked",
        { key: "id", sortable: true },
        { key: "Name", sortable: true },
        { key: "count", label: "User Count", sortable: true }
      ],
      addIsOpened: false,
      editIsOpened: false,
      newLocation: "",
      sortBy: "age",
      sortDesc: false,
      errors: {},
      alerts: {},
      locations: [],
      allLocationsSelected: false,
      locationsSelected: [],
      filter: null,
      filterOn: []
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
          this.loadLocations();
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
.row {
  margin-right: 0;
}
.link-profile {
  color: #00acc1;
  cursor: pointer;
  &:hover {
    color: #af4448;
  }
}
.button-back {
  cursor: pointer;
  color: #004ba0;
  &:hover {
    color: #af4448;
  }
}
</style>
