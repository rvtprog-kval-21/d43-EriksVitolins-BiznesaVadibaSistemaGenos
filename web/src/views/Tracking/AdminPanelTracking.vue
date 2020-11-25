<template>
    <b-container class="mt-3 body">
        <div class="d-flex justify-content-end">
            <a href="/tracking"><b-button variant="outline-info">Back</b-button></a>
        </div>
        <hr />
        <div class="mt-3 mb-3 d-flex justify-content-between">
            <h5>Add a blogger:</h5>
            <vSelect
                    class="w-75"
                    v-model="selected"
                    label="email"
                    :options="options"
            ></vSelect>
            <b-button variant="success" @click="addUser()">Add</b-button>
        </div>
        <div class="d-flex mt-4 mb-2 justify-content-between">
            <h3>Bloggers:</h3>
            <b-button variant="danger" @click="deleteUser()">Kick Checked</b-button>
        </div>
        <vue-good-table
                :select-options="{ enabled: true }"
                :search-options="{ enabled: true }"
                @on-selected-rows-change="selectionChanged"
                :columns="columns"
                :rows="users"
        />
    </b-container>
</template>

<script>
    import vSelect from "vue-select";
    import "vue-select/dist/vue-select.css";

    export default {
        name: "AdminPanelTracking",
        components: { vSelect },
        data() {
            return {
                options: [],
                selected: "",
                users: [],
                columns: [
                    {
                        label: "Id",
                        field: "id",
                        type: "number"
                    },
                    {
                        label: "Email",
                        field: "email"
                    },
                    {
                        label: "Name",
                        field: "name"
                    },
                    {
                        label: "Title",
                        field: "title"
                    }
                ],
                checkedItems: []
            };
        },
        methods: {
            selectionChanged(select) {
                this.checkedItems = select.selectedRows;
            },
            getUsers() {
                window.axios.post("api/admin/users").then(res => {
                    this.options = res.data.users;
                });
            },
            addUser() {
                window.axios
                    .get("api/admin/manager/" + this.selected.id + "/add")
                    .then(res => {
                        this.makeToast(res.data.message, "success");
                        this.selected = "";
                        this.getBloggers();
                    })
                    .catch(rej => {
                        this.makeToast(rej.response.data.error, "danger");
                    });
            },
            getManagers() {
                window.axios
                    .get("api/admin/users/managers")
                    .then(res => {
                        this.users = res.data.users;
                    })
                    .catch(rej => {
                        this.makeToast(rej.response.data.error, "danger");
                    });
            },
            deleteUser() {
                for (let iter = 0; iter < this.checkedItems.length; iter++) {
                    const email = this.checkedItems[iter].email;
                    window.axios
                        .get("api/admin/manager/" + this.checkedItems[iter].id + "/delete")
                        .then(res => {
                            this.makeToast(email + " " + res.data.message, "success");
                            this.getBloggers();
                        })
                        .catch(rej => {
                            this.makeToast(email + " " + rej.response.data.error, "danger");
                        });
                }
            },
            makeToast(text, variant) {
                this.$bvToast.toast(text, {
                    autoHideDelay: 5000,
                    variant: variant,
                    title: "Notification"
                });
            }
        },
        async created() {
            await this.getManagers();
            this.getUsers();
        }
    }
</script>

<style scoped>
    .body {
        min-height: 800px;
    }
</style>
