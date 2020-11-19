<template>
  <div class="background">
    <div class="root">
      <div class="login">
        <div v-if="biggerThanBreakPoint" class="content image">
          <div class="screen">
            <img src="../../assets/businessmeetingbw.svg" alt="" />
          </div>
        </div>
        <div class="content right-side">
          <div class="screen right">
            <h1>Login</h1>
            <b-alert v-if="authError" show variant="danger"
              ><p>{{ authError }}</p></b-alert
            >
            <form @submit.prevent="authenticate">
              <div class="login-email">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="icon icon-tabler icon-tabler-user"
                  width="28"
                  height="28"
                  viewBox="0 0 24 24"
                  stroke-width="1"
                  stroke="#212121"
                  fill="none"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                >
                  <path stroke="none" d="M0 0h24v24H0z" />
                  <circle cx="12" cy="7" r="4" />
                  <path d="M6 21v-2a4 4 0 0 1 4 -4h4a4 4 0 0 1 4 4v2" />
                </svg>
                <input
                  type="email"
                  name="email"
                  v-model="form.email"
                  placeholder="Email"
                />
              </div>
              <div class="login-password">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="icon icon-tabler icon-tabler-lock"
                  width="28"
                  height="28"
                  viewBox="0 0 24 24"
                  stroke-width="1"
                  stroke="#212121"
                  fill="none"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                >
                  <path stroke="none" d="M0 0h24v24H0z" />
                  <rect x="5" y="11" width="14" height="10" rx="2" />
                  <circle cx="12" cy="16" r="1" />
                  <path d="M8 11v-4a4 4 0 0 1 8 0v4" />
                </svg>
                <input
                  type="password"
                  name="password"
                  v-model="form.password"
                  placeholder="Password"
                />
              </div>
              <div class="login-btn">
                <input type="submit" value="Login" />
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { login } from "../../helpers/auth";

export default {
  name: "Login",
  data() {
    return {
      form: {
        email: "",
        password: ""
      },
      biggerThanBreakPoint: false,
      error: null
    };
  },
  created() {
    window.addEventListener("resize", this.handleResize);
    this.handleResize();
  },
  methods: {
    authenticate() {
      this.$store.dispatch("login");
      if (!this.$data.form.email.length || !this.$data.form.password.length) {
        this.$store.commit("loginFailed", {
          error: "Email or password haven't been filled"
        });
      } else {
        login(this.$data.form)
          .then(res => {
            this.$store.commit("loginSuccess", res);
            this.$router.push({ path: "/home" });
          })
          .catch(error => {
            this.$store.commit("loginFailed", {
              error: error.response.data.error
            });
          });
      }
    },
    handleResize() {
      if (window.innerWidth < 1128) {
        this.biggerThanBreakPoint = false;
      } else {
        this.biggerThanBreakPoint = true;
      }
    }
  },
  computed: {
    authError() {
      return this.$store.getters.authError;
    }
  }
};
</script>

<style scoped lang="scss">
.background {
  width: 100%;
  height: 100%;
  background-color: #f5f5f5;
}
.root {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
  width: 100%;
}

.login {
  width: 70%;
  height: 700px;
  background-color: white;
  border-radius: 25px;
  box-shadow: 0 8px 16px 0 rgba(0, 0, 0, 0.2);
  display: flex;
}

.content {
  width: 50%;

  &.right-side {
    padding: 100px 0 100px 0;
  }
}

.screen {
  height: 500px;
  margin: auto;
}

.right.screen {
  width: 80%;
}

h1 {
  margin-bottom: 50px;
}

form {
  display: flex;
  flex-direction: column;
}

.login-email {
  margin-bottom: 30px;
}

.login-btn {
  margin-top: 40px;
  display: flex;
  justify-content: center;
  margin-left: 40px;
}

.login-btn input {
  background-color: #007c91;
  border-radius: 12px;
  border: 1px solid #007c91;
  display: inline-block;
  cursor: pointer;
  color: #ffffff;
  font-family: Arial;
  font-size: 16px;
  padding: 18px 47px;
  text-decoration: none;
  text-shadow: 0px 1px 0px #2f6627;
}

.login-btn input:hover {
  background-color: #00acc1;
}

.login-btn input:active {
  position: relative;
  top: 1px;
}

.login-email input,
.login-password input {
  width: 100%;
  border: none;
  border-bottom: 1px solid #bdbdbd;
  font-size: 16px;
  padding-bottom: 7px;
  padding-top: 7px;
  text-align: center;
}

.login-email,
.login-password {
  display: flex;
}

svg {
  margin-right: 10px;
}

@media screen and (max-width: 1127px) {
  .content {
    width: 100%;
  }
}
</style>
