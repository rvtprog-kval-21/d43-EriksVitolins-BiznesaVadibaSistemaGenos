<template>
    <div>
       <div class="row-blog">
           <h5>Publish date:</h5>
           <b-form-datepicker class="w-75" v-model="form.publish_at"></b-form-datepicker>
       </div>
        <div class="">
           <div class="row-blog">
               <h5>Article title:</h5>
               <b-form-input class="w-75" v-model="form.title" placeholder="Enter your title"></b-form-input>
           </div>
            <div class="mt-1 mb-4">
                <VueShowdown
                        :markdown="form.title"
                        flavor="github"
                        :options="{ emoji: true }"></VueShowdown>
            </div>
        </div>
       <div>
           <div class="row-blog">
               <h5>Article head text:</h5>
               <b-form-input class="w-75" v-model="form.headtext" placeholder="Enter your headtext"></b-form-input>
           </div>
           <div class="mt-1 mb-4">
               <VueShowdown
                       :markdown="form.headtext"
                       flavor="github"
                       :options="{ emoji: true }"></VueShowdown>
           </div>
       </div>
       <div>
           <div class="row-blog">
               <h5>Article topic:</h5>
               <b-form-input class="w-75" v-model="form.topic" placeholder="Enter your topic"></b-form-input>
           </div>
           <div class="mt-1 mb-4">
               <VueShowdown
                       :markdown="form.topic"
                       flavor="github"
                       :options="{ emoji: true }"></VueShowdown>
           </div>
       </div>
       <div>
           <div class="row-blog">
               <h5>Article photo:</h5>
               <b-form-file
                       v-model="form.photo"
                       :state="Boolean(form.photo)"
                       class="w-75"
                       placeholder="Choose a photo or drop it here..."
                       drop-placeholder="Drop photo here..."
                       accept="image/*"
               ></b-form-file>
           </div>
       </div>
       <div>
           <div class="row-blog">
               <h5>Article content:</h5>
               <b-form-textarea
                       v-model="form.content"
                       placeholder="Article content"
                       rows="9"
                       class="w-75"
                       no-resize
               ></b-form-textarea>
           </div>
           <div class="mt-1 mb-4">
               <VueShowdown
                       :markdown="form.content"
                       flavor="github"
                       :options="{ emoji: true }"></VueShowdown>
           </div>
       </div>
        <div class="d-flex justify-content-between">
            <b-button @click="clearForm" variant="warning">Reset</b-button>
            <b-button variant="success">Save</b-button>
        </div>
    </div>
</template>

<script>
    export default {
        name: "CreateBlog",
        props: ["isEdit","editForm"],
        data () {
            return {
                form: {
                    title: "",
                    publish_at: "",
                    headtext: "",
                    photo: null,
                    topic: "",
                    content: "",
                }
            }
        },
        methods: {
            clearForm() {
                this.form.title = ""
                this.form.publish_at = null
                this.form.headtext = ""
                this.form.photo = null
                this.form.topic = ""
                this.form.content = ""
            },
            init() {
               if (this.isEdit === true) {
                   this.form = this.editForm
               }
            },
            addUser() {
                window.axios
                    .post("api/blog/add", this.form)
                    .then(res => {
                        this.makeToast(res.data.message, "success");
                    })
                    .catch(rej => {
                        this.makeToast(rej.response.data.error, "danger");
                    });
            },
            getBloggers() {
                window.axios
                    .post("api/blog/update/" + this.selected.id, this.form)
                    .then(res => {
                        this.makeToast(res.data.message, "success");
                    })
                    .catch(rej => {
                        this.makeToast(rej.response.data.error, "danger");
                    });
            },
        },
        created() {
            this.init()
        }
    }
</script>

<style scoped lang="scss">
    .row-blog{
        display: flex;
        margin-top: 20px;
        align-items: center;
        justify-content: space-between;
        h5{
            margin-right: 25px;
            margin-bottom: 0;
        }
    }
</style>