<template>
  <b-container class="mt-3">
    <template v-if="currentUser.role === 'admin' || isCreator">
      <div class="d-flex justify-content-end">
        <a href="/blog/admin">
          <b-button
            class="mr-3"
            v-if="currentUser.role === 'admin'"
            variant="outline-info"
            >Admin Panel</b-button
          >
        </a>
        <a href="/blog/creator"
          ><b-button v-if="isCreator" variant="outline-info"
            >Creator Panel</b-button
          ></a
        >
      </div>
    </template>
    <div class="mt-4">
      <template v-for="(iter, index) in blogs">
       <div :key="index" @click="$router.push('/blog/' + iter.id + '/view')" class="d-flex blog">
         <div class="image">
           <img :src="getImgUrl(iter.photo)" alt="">
         </div>
         <div class="content">
           <div class="top"><div class="info">
             <template v-if="iter.topic">
               <p>{{iter.topic}}</p>
             </template>
             <b-icon icon="dot"></b-icon>
             <template v-if="iter.publish_at">
               <p>{{getDate(iter.publish_at)}}</p>
             </template>
           </div>
             <template v-if="iter.title">
               <h3 class="mt-3 mb-3">
                 <VueShowdown
                         :markdown="iter.title"
                         flavor="github"
                         :options="{ emoji: true }"
                 ></VueShowdown>
               </h3>
             </template>

             <template v-if="iter.headtext">
               <p class="d-flex mb-3">
                 <VueShowdown
                         :markdown="iter.headtext"
                         flavor="github"
                         :options="{ emoji: true }"
                 ></VueShowdown>
               </p>
             </template>
           </div>
           <div class="bottom">
             <div class=" d-flex d-flex align-items-center">
               <b-avatar :src="getImgUrl(iter.user.avatar)"></b-avatar>
               <div class="user ml-3">
                 <p>{{iter.user.name + " " + iter.user.last_name}}</p>
                 <p class="grey">{{iter.user.title}}</p>
               </div>
             </div>
           </div>
         </div>
       </div>
      </template>
    </div>
  </b-container>
</template>

<script>
export default {
  name: "Blog",
  computed: {
    currentUser() {
      return this.$store.getters.currentUser;
    }
  },
  data() {
    return {
      errors: {},
      isCreator: true,
      blogs: []
    };
  },
  methods: {
    getBlogs() {
      window.axios
              .get("api/blog/all/")
              .then(res => {
                this.blogs = res.data.blogs;
              })
              .catch(rej => {
                console.log(rej.response.data.error)
              });
    },
    getImgUrl(image) {
      let images = process.env.VUE_APP_API + "/static" + image;
      return images;
    },
    getDate(date) {
      date = new Date(date);
      const day = date.getDate();
      const month = date.toLocaleString("default", { month: "long" });
      const year = date.getFullYear()
      return day + " " + month + " " + year;
    },
  },
  created() {
    this.getBlogs()
  }
};
</script>

<style scoped lang="scss">
  .blog{
    width: 100%;
    cursor: pointer;
    height: 400px;
    margin-bottom: 80px;
    .image{
      width: 60%;
      img{
        width: 100%;
        height: 100%;
      }
    }
    .content{
      display: flex;
      justify-content: space-between;
      flex-direction: column;
      padding-left: 10px;
      .info{
        display: flex;
        font-weight: bolder;
        align-items: center;
      }
      p{
        margin-bottom: 0;
      }
    }
  }
</style>
