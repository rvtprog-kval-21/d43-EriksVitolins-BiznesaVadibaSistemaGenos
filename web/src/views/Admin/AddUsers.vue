<template>
  <b-container>
    <div class="background">
      <div class="topBar">
        <h1>Add users</h1>
        <div class="input">
          <p class="mb-0">Skaits:</p>
          <b-form-input
            type="number"
            v-model="addFieldsCount"
            placeholder="Enter how many fields to add"
          ></b-form-input>
          <input
            class="btn-add-users btn"
            type="submit"
            @click="addField"
            value="Add"
          />
        </div>
      </div>
      <div class="input-fields">
        <div class="input-box" v-for="(field, index) in fields" :key="index">
          <div class="top">
            <h1>User</h1>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              @click="removeField(index)"
              class="icon icon-tabler icon-tabler-x"
              width="35"
              height="35"
              viewBox="0 0 24 24"
              stroke-width="2"
              stroke="#607D8B"
              fill="none"
              stroke-linecap="round"
              stroke-linejoin="round"
            >
              <path stroke="none" d="M0 0h24v24H0z" />
              <line x1="18" y1="6" x2="6" y2="18" />
              <line x1="6" y1="6" x2="18" y2="18" />
            </svg>
          </div>
          <hr />
          <div class="input-container">
            <div>
              <b-form-input
                class="mb-2"
                v-model="field.name"
                placeholder="Enter the users name"
              ></b-form-input>
              <b-form-input
                v-model="field.lastname"
                placeholder="Enter the users last name"
              ></b-form-input>
            </div>
            <div>
              <b-form-input
                class="mb-2"
                type="email"
                v-model="field.email"
                placeholder="Enter the users email"
              ></b-form-input>
              <b-form-select
                v-model="field.role"
                :options="options"
              ></b-form-select>
            </div>
          </div>
        </div>
      </div>
      <button @click="add">Submit</button>
    </div>
  </b-container>
</template>

<script>
import validate from "validate.js";

export default {
  name: "AddUsers",
  data() {
    return {
      fields: [{ name: "", lastname: "", email: "", role: "" }],
      addFieldsCount: 1,
      errorArray: [],
      alerts: "",
      options: [
        { value: "", text: "Please select users role" },
        { value: "regular", text: "Regular" },
        { value: "admin", text: "Admin" }
      ]
    };
  },
  methods: {
    addField: function() {
      for (let iter = 0; iter < this.addFieldsCount; iter++) {
        this.fields.push({
          name: "",
          lastname: "",
          email: "",
          role: "regular"
        });
      }
    },
    removeField: function(index) {
      this.fields.splice(index, 1);
    },
    add() {
      const constraints = this.getConstraints();
      for (let iter = 0; iter < this.fields.length; iter++) {
        const errors = validate(this.fields[iter], constraints);
        if (errors) {
          console.log(errors)
          this.makeToast(errors.email, "danger");
          this.makeToast(errors.role, "danger");
          this.makeToast(errors.lastname, "danger");
          this.makeToast(errors.name, "danger");
          return;
        }
        window.axios
          .post("/api/signup", this.fields[iter])
          .then(() => {
            this.fields.splice(iter, 1);
          })
          .catch(function(rej) {
            this.makeToast(rej.response.data.error, "danger");
          });
      }
    },
    makeToast(text, variant) {
      this.$bvToast.toast(text, {
        autoHideDelay: 5000,
        variant: variant,
        title: "Notification"
      });
    },
    getConstraints() {
      return {
        name: {
          presence: true,
          length: {
            minimum: 3,
            message: ": Need to have at least 3 letters"
          }
        },
        email: {
          presence: true,
          email: true
        },
        role: {
          presence: true,
          length: {
            minimum: 3
          }
        },
        lastname: {
          presence: true,
          length: {
            minimum: 3,
            message: ":  Need to have at least 3 letters"
          }
        }
      };
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
.container {
  padding: 100px 100px 100px 100px;
}

.background {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.topBar {
  width: 100%;
  height: 100px;
  display: flex;
  align-items: center;
  justify-content: space-around;
}

p {
  font-size: 18px;
}

.topBar .input {
  display: flex;
  align-items: center;
}

.topBar input {
  margin-left: 5px;
  margin-right: 15px;
}

.input-box {
  width: 900px;
  min-height: 200px;
  background-color: #f5f5f5;
  margin: auto auto 40px;
  border-radius: 25px;
}

.top {
  margin: auto;
  padding-top: 20px;
  width: 90%;
  display: flex;
  justify-content: space-between;
}

.top h1 {
  margin: 0;
}

.top svg:hover {
  stroke: red;
  cursor: pointer;
}

.btn {
  color: white;
  background-color: #76d275;
  padding: 8px 16px;
  border: none;
  border-radius: 20px;
}

.btn:hover {
  cursor: pointer;
  background-color: #43a047;
}

hr {
  width: 85%;
  border-top: 1px solid #cfd8dc;
}

.input-container {
  width: 80%;
  margin: auto;
  display: flex;
  justify-content: space-around;
}

.input-wrapper {
  display: flex;
  flex-direction: column;
  margin: 15px 0px 5px 0px;
  align-items: center;
}

button {
  width: 30%;
  color: #ffffff;
  background-color: #2196f3;
  border: none;
  padding: 12px 16px;
  vertical-align: middle;
  text-align: center;
  cursor: pointer;
  border-radius: 20px;
  font-size: 18px;
  margin-bottom: 40px;
}

button:hover {
  background-color: cornflowerblue;
}

.input-wrapper .input-field {
  border-radius: 25px;
  border: 1px solid grey;
  height: 22px;
  margin-top: 5px;
}

@media screen and (max-width: 1500px) {
  .input-box {
    width: 700px;

    .input-container {
      display: flex;
      flex-direction: column;
    }
  }
}

@media screen and (max-width: 1200px) {
  .input-box {
    width: 500px;
  }
}

@media screen and (max-width: 750px) {
  .input-box {
    width: 350px;
  }
}

@media screen and (max-width: 750px) {
  .wrapper {
    padding: 0;
    .container {
      padding: 0;
    }
    .topBar {
      flex-direction: column;
      margin-bottom: 20px;
    }
  }
}
</style>
