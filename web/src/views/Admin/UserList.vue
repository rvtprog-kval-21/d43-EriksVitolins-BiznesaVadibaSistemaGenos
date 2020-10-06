<template>
  <div class="content">
    <div class="table">
      <alertComponent v-if="this.errors.error" class="alert-danger">
        <p v-if="this.errors.error">{{ this.errors.error }}</p>
      </alertComponent>
      <div class="top">
        <div class="profile">User List</div>
        <div class="header">
          <div class="left">
            <div class="total" v-if="this.users.total">
              users: {{ this.users.total }}
            </div>
            <div class="sort">
              sort by:
              <select v-model="sort_by" @change="loadUsers">
                <option value="email DESC">Email descending</option>
                <option value="email ASC">Email ascending</option>
                <option value="role DESC">Role descending</option>
                <option value="role ASC">Role ascending</option>
                <option selected value="created_at DESC"
                  >Date descending</option
                >
                <option value="created_at ASC">Date ascending</option>
              </select>
            </div>
          </div>
          <div class="right">
            <div class="input">
              <p>search:</p>
              <input
                v-model="search_term"
                placeholder="email"
                @keyup.enter="search"
                type="text"
              />
            </div>
            <button><a href="/admin/userAdd">Add Users</a></button>
          </div>
        </div>
      </div>
      <div class="body">
        <div class="table-header grid">
          <p class="id">Avatar</p>
          <p>Email</p>
          <p>Role</p>
          <p>Created At</p>
        </div>
        <div
          class="user grid"
          v-for="(user, index) in this.users.data"
          :key="index"
        >
          <div class="id">
            <img :src="getImgUrl(user.id)" alt="" />
          </div>
          <div>
            <p class="click-me" @click="toProfile(user.id)">{{ user.email }}</p>
          </div>
          <div>
            <p>{{ user.role }}</p>
          </div>
          <div>
            <p>{{ user.created_at }}</p>
          </div>
          <div class="more-options">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              @click="toProfile(user.id)"
              class="icon icon-tabler icon-tabler-dots-vertical"
              width="20"
              height="20"
              viewBox="0 0 24 24"
              stroke-width="2"
              stroke="#607D8B"
              fill="none"
              stroke-linecap="round"
              stroke-linejoin="round"
            >
              <path stroke="none" d="M0 0h24v24H0z" />
              <circle cx="12" cy="12" r="1" />
              <circle cx="12" cy="19" r="1" />
              <circle cx="12" cy="5" r="1" />
            </svg>
          </div>
        </div>
      </div>
      <div class="footer">
        <pagination
          :data="this.users"
          @pagination-change-page="getResults"
        ></pagination>
      </div>
    </div>
  </div>
</template>

<script>
import alertComponent from "../../components/alerts/alert-component";
export default {
  name: "UserList",
  components: { alertComponent },
  data() {
    return {
      users: {},
      errors: {},
      page: 1,
      search_term: "",
      sort_by: "created_at DESC"
    };
  },
  mounted() {
    this.loadUsers();
  },
  methods: {
    loadUsers() {
      this.errors = [];
      const vue = this;
      window.axios
        .post("/api/users?page=" + this.page, { sort_by: this.sort_by })
        .then(response => {
          this.users = response.data.data;
          this.page = this.users.current_page;
        })
        .catch(function(errors) {
          vue.errors = { error: errors.response.data.message };
        });
    },
    getImgUrl(id) {
      let images =
        process.env.VUE_APP_API + "/storage/avatars/" + id + "/avatar.png";
      return images;
    },
    getResults(page = 1) {
      this.errors = [];
      const vue = this;
      window.axios
        .post("/api/users?page=" + page, { sort_by: this.sort_by })
        .then(response => {
          this.users = response.data.data;
          this.page = this.users.current_page;
        })
        .catch(function(errors) {
          vue.errors = { error: errors.response.data.message };
        });
    },
    search() {
      this.errors = [];
      const vue = this;
      if (this.search_term.length > 0) {
        window.axios
          .post("/api/users/search", { search: this.search_term })
          .then(response => {
            if (response.data.data) {
              this.users = response.data.data;
            } else {
              this.errors = { error: response.data.errors };
            }
          })
          .catch(function(errors) {
            vue.errors = { error: errors.response.data.message };
          });
      } else {
        this.loadUsers();
      }
    },
    toProfile(id) {
      this.$router.push("/user/" + id);
    }
  }
};
</script>

<style lang="scss" scoped>
.content {
  overflow: auto;
  width: 80%;
  margin-right: auto;
  margin-left: auto;
  padding: 0 0 0 0;

  .table {
    display: flex;
    justify-content: space-between;
    flex-direction: column;

    p {
      word-break: break-all;
      font-size: 18px;
    }

    .profile {
      font-size: 32px;
      margin-right: 10px;
      color: #ffc107;
    }

    .header {
      display: flex;
      justify-content: space-between;
      margin-bottom: 25px;

      .left,
      .right {
        display: flex;
        align-items: center;

        select {
          border: none;
          color: #636b6f;
        }

        div {
          color: #636b6f;
          margin-right: 20px;
          font-size: 18px;
        }
      }

      .right {
        .input {
          margin-right: 25px;
          display: flex;

          p {
            margin-right: 5px;
          }

          input {
            border: 1px solid #636b6f;
            border-radius: 5px;
          }
        }

        button {
          border: none;
          background: none;

          a {
            padding: 6px 6px;
            text-align: center;
            cursor: pointer;
            border: none;
            background-color: #62757f;
            color: white;

            &:hover {
              background-color: #4a5568;
            }
          }
        }
      }
    }

    .table-header {
      margin-bottom: 20px;
    }

    .grid {
      display: grid;
      grid-template-columns: 60px auto 100px 150px 20px;
      grid-template-rows: 50px;

      div {
        overflow: hidden;
        display: flex;
        align-items: center;

        p {
          text-overflow: ellipsis;
          overflow: hidden;
          font-size: 16px;
        }
      }
    }

    .footer {
      display: flex;
      justify-content: center;
      align-items: center;

      .pagination {
        display: flex;
      }
    }

    .id {
      display: flex;
      justify-content: center;

      img {
        width: 35px;
      }
    }

    .grid.table-header {
      grid-template-rows: 20px;
    }

    .body {
      .user {
        background-color: #eceff1;
        margin-bottom: 15px;
        border-radius: 10px;

        .click-me {
          cursor: pointer;
          color: #1976d2;

          &:hover {
            color: #004ba0;
            text-decoration: none;
          }
        }

        .more-options {
          display: flex;
          align-items: center;

          svg {
            cursor: pointer;

            &:hover {
              stroke: red;
            }
          }
        }
      }
    }

    @media screen and (max-width: 750px) {
      .header {
        flex-direction: column;
        align-items: center;
        .left {
          margin-bottom: 50px;
        }
      }

      .grid {
        grid-template-rows: 200px;
      }
    }
  }
}
</style>

<style lang="scss">
.pagination {
  .page-item {
    margin: 25px 10px;
    list-style: none;

    .page-link {
      padding: 12px 12px;
      text-align: center;
      cursor: pointer;
      border: none;
      background-color: #62757f;
      color: white;

      &:hover {
        background-color: #4a5568;
      }
    }

    &.active .page-link {
      background-color: #90a4ae;

      span {
        display: none;
      }
    }
  }
}
</style>
