<template>
    <b-container class="mt-3 container">
        <template v-if="currentUser.role === 'admin' || isManager">
            <div class="d-flex justify-content-end">
                <a href="/tracking/admin">
                    <b-button
                            class="mr-3"
                            v-if="currentUser.role === 'admin'"
                            variant="outline-info"
                    >Admin Panel</b-button
                    >
                </a>
                <a href="/tracking/manager"
                ><b-button v-if="isManager" variant="outline-info"
                >Creator Panel</b-button
                ></a
                >
            </div>
        </template>
       <div class="mt-4 form">
           <h3>Submissions:</h3>
           <div class="mt-4">
               <h5 class="mb-2">Subject:</h5>
               <b-form-input class="w-75" v-model="form.subject" placeholder="Enter Subject"></b-form-input>
           </div>

           <div class="mt-4">
               <h5 class="mb-2">Description:</h5>
               <b-form-textarea
                       v-model="form.description"
                       placeholder="Enter Description"
                       rows="3"
                       max-rows="6"
               ></b-form-textarea>
           </div>
           <div class="mt-4">
               <h5 class="mb-2">Attachments:</h5>
               <b-form-file
                       class="w-75"
                       multiple
                       v-model="form.files"
                       :state="Boolean(form.files)"
                       placeholder="Choose a file or drop it here..."
                       drop-placeholder="Drop file here..."
               ></b-form-file>
           </div>
           <div class="d-flex justify-content-between mt-5">
               <div></div>
               <b-button @click="goSubmit" variant="success">Submit</b-button>
           </div>
       </div>
    </b-container>
</template>

<script>
    export default {
        name: "Tracking",
        computed: {
            currentUser() {
                return this.$store.getters.currentUser;
            }
        },
        data() {
            return {
                errors: {},
                isManager: false,
                blogs: [],
                formID: 0,
                form: {
                    subject: "",
                    description: "",
                    files: null
                }
            };
        },
        methods: {
            getIsManager() {
                window.axios
                    .get("api/manager/check/ismanager/")
                    .then(res => {
                        this.isManager = res.data.isManager;
                    })
                    .catch(rej => {
                        console.log(rej.response.data.error)
                    });
            },
            goSubmit(){
                if (this.form.subject != '' && this.form.description != '') {
                    let formData = new FormData();
                    formData.append("subject", `${this.form.subject}`);
                    formData.append("description", `${this.form.description}`);
                    window.axios
                        .post("api/tracking/add/",formData, {
                            headers: {
                                "Content-Type": "multipart/form-data"
                            }
                        })
                        .then(res => {
                            this.formID = res.data.id
                            this.makeToast("Submitted successfully", "success")
                            this.form.subject = ""
                            this.form.description = ""
                        })
                        .catch(rej => {
                            this.makeToast(rej.response.data.error, "danger")
                        });

                    for (let iter = 0; iter < this.form.files.length; iter += 1) {
                        let formData = new FormData();
                        console.log(this.form.files[iter])
                        formData.append("file", `${this.form.files[iter]}`);
                        formData.append("id", `${this.formID}`);
                        window.axios
                            .post("api/tracking/add/attachment/",formData, {
                                headers: {
                                    "Content-Type": "multipart/form-data"
                                }
                            })
                            .then(() => {
                                console.log("saved")
                            })
                            .catch(rej => {
                                this.makeToast(rej.response.data.error, "danger")
                            });
                    }
                    this.formID = 0
                    this.form.files = null
                } else {
                    this.makeToast("Subject or description wasn't filled", "danger")
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
        created() {
            this.getIsManager()
        }
    }
</script>

<style scoped lang="scss">
    .form{
        h5{
            margin: 0;
        }
    }
    .container{
        height: 100%;
    }
</style>