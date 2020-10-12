<template>
    <div class="overflow-auto">
        <div class="">
            <b-alert v-if="this.errors.error" variant="danger" show="">{{this.errors.error}}</b-alert>
        </div>
        <b-row>
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
                            <b-button :disabled="!filter" @click="filter = ''">Clear</b-button>
                        </b-input-group-append>
                    </b-input-group>
                </b-form-group>
            </b-col>
        </b-row>
        <b-table :filter="filter"
                 :filter-included-fields="filterOn"
                 primary-key="id" :fields="fields"
                 :tbody-transition-props="transProps"
                 head-variant="light" :busy="isBusy" class="overflow-auto"
                  striped hover :items="users">
            <template v-slot:table-busy>
                <div class="text-center text-danger my-2">
                    <b-spinner class="align-middle"></b-spinner>
                    <strong>Loading...</strong>

                </div>
            </template>
            <template v-slot:cell(name)="data">
                {{ data.item.name }} {{ data.item.lastname }}
            </template>
        </b-table>
    </div>
</template>

<script>
export default {
  name: "UserListTest",
    data() {
        return {
            transProps: {
                // Transition name
                name: 'flip-list'
            },
            users: {},
            isBusy: false,
            errors: {},
            filter: null,
            filterOn: [],
            fields: [
                { key: 'id', sortable: true },
                { key: 'name', sortable: true },
                { key: 'email', sortable: true },
                { key: 'role', sortable: true },
                { key: 'created_at', sortable: true }
            ]
        };
    },
    mounted() {
        this.loadUsers();
    },
    methods: {
        loadUsers() {
            this.toggleBusy()
            this.errors = [];
            const vue = this;
            window.axios
                .post("/api/users", { sort_by: this.sort_by })
                .then(response => {
                    this.users = response.data.data;
                    this.page = this.users.current_page;
                    this.toggleBusy()
                })
                .catch(function(errors) {
                    vue.errors = { error: errors.response.data.message };
                });
        },
        toProfile(id) {
            this.$router.push("/user/" + id);
        },
        toggleBusy() {
            this.isBusy = !this.isBusy
        }
    }
};
</script>

<style scoped>
    .form-row{
        margin-right: 0;
    }
    .row{
        margin-right: 0;
    }
</style>
