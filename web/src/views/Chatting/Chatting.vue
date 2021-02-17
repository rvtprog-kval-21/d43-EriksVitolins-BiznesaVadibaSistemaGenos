<template>
  <div class="h-100">
    <div class="chat" v-if="!CreatingNewGroup">
      <div class="left">
        <div class="header d-flex justify-content-between">
          <b-input-group size="sm" class="w-75 h-75 ml-3">
            <b-input-group-prepend is-text>
              <b-icon icon="search"></b-icon>
            </b-input-group-prepend>
            <b-form-input v-model="searchTerm" type="search" class="h-100" placeholder="Search Group"></b-form-input>
          </b-input-group>
          <b-button variant="white" @click="CreatingNewGroup = !CreatingNewGroup"><b-icon icon="plus-circle-fill" variant="primary"></b-icon></b-button>
        </div>
        <div class="groups">
         <template v-for="(iter,index) in roomsSearched">
           <div class="group-row"  :class="{ 'opened' : selectedID == iter.id}" @click="getRoom(iter.id)" :key="index">
             <p>{{iter.name}}</p>
           </div>
         </template>
        </div>
      </div>
      <div class="right" v-if="selectedID !==0 && room">
        <div class="header d-flex justify-content-between">
          <div class="d-flex">
            <b-avatar v-if="room.avatar !== ''" :src="getImgUrl(room.avatar)"></b-avatar>
           <h5>{{room.name}}</h5>
          </div>
          <div>
            <b-button variant="white">
              <b-icon icon="gear" ></b-icon>
            </b-button>
          </div>
        </div>
        <div class="chat-field">
          <div class="chat">Hey</div>
          <div class="bottom">
            <div class="pl-5">
              <b-icon class="mr-2 hover" icon="mic" variant="primary" font-scale="1.5"></b-icon>
            </div>
            <div class="w-75">
              <b-form-textarea
                  v-model="message"
                  placeholder="Type your message"
                  rows="1"
                  max-rows="4"
              ></b-form-textarea>
            </div>
            <div class="pr-5">
              <b-icon class="mr-2 hover" icon="paperclip" variant="primary" font-scale="1.5"></b-icon>
              <b-icon v-if="message === ''" icon="forward"  v-b-popover.hover.top="'Can not send message when it is empty'" font-scale="1.5"></b-icon>
              <b-icon class="hover" v-else @click="SendMessage()" icon="forward" variant="primary" font-scale="1.5"></b-icon>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div v-if="CreatingNewGroup">
      <b-container class="mt-4">
        <b-button class="mb-5" variant="outline-primary" @click="CreatingNewGroup = !CreatingNewGroup">Back</b-button>
        <div class="mb-3">
          <p>Name:</p>
          <b-form-input v-model="newGroup.name" class="h-100 mb-3" placeholder="Group Name"></b-form-input>
          <p>About:</p>
          <b-form-textarea v-model="newGroup.about" class="h-100"  rows="3"
                           max-rows="6" placeholder="Group About"></b-form-textarea>
        </div>
        <p>Participants:</p>
        <vSelect
            class="w-100"
            v-model="newGroup.participants"
            label="email"
            multiple
            :options="users"
        ></vSelect>
        <div class="mt-5 d-flex justify-content-between">
          <div></div>
          <b-button variant="outline-success" @click="createNewGroup()">Save</b-button>
        </div>
      </b-container>
    </div>
  </div>
</template>

<script>
import vSelect from "vue-select";
import "vue-select/dist/vue-select.css";

export default {
  name: "Chatting",
  components: {
    vSelect
  },
  data() {
    return {
      selectedID: 0,
      CreatingNewGroup: false,
      newGroup: {
        name: "",
        about: "",
        participants: [],
      },
      rooms: [],
      roomsSearched: [],
      searchTerm: "",
      message: "",
      room: {},
      users: [],
    }
  },
  computed: {
    currentUser() {
      return this.$store.getters.currentUser;
    }
  },
  methods: {
    getImgUrl(image) {
      let images = process.env.VUE_APP_API + "/static" + image;
      return images;
    },
    getUsers() {
      window.axios.post("api/users/not/included").then(res => {
        this.users = res.data.users;
      });
    },
    SendMessage(){

    },
    createNewGroup() {
      if (this.newGroup.name === "") {
        this.makeToast("Name needs to be filled", "danger")
        return
      }

      window.axios.post("api/chatting/new/group",this.newGroup).then(() => {
        this.makeToast("group made successfully", "success")
        this.CreatingNewGroup = false;
        this.newGroup = {
          name:"",
          about: "",
          participants: [],
        }
        this.getRooms()
      })
      .catch(() => {
        this.makeToast("There was a problem", "danger")
      });
    },
    getRooms() {
      window.axios.get("api/chatting/get/rooms").then(res => {
        this.rooms = res.data.rooms
        this.searchRoom()
      })
    },
    searchRoom() {
      this.roomsSearched = [];
      for (let i = 0; this.rooms.length > i; i++) {
        if (this.searchTerm === ""|| this.rooms[i].name.includes(this.searchTerm)) {
          console.log(this.roomsSearched)
          this.roomsSearched.push(this.rooms[i])
        }
      }
    },
    getRoom(id) {
      this.selectedID = id
      window.axios.post("api/chatting/get/room", {"id": this.selectedID}).then(res => {
        this.room = res.data.room
      })
    },
    makeToast(text, variant) {
      this.$bvToast.toast(text, {
        autoHideDelay: 5000,
        variant: variant,
        title: "Notification"
      });
    },
  },
  watch: {
    // whenever question changes, this function will run
    searchTerm: function () {
      this.searchRoom()
    },
  },
   created() {
    this.getUsers()
     this.getRooms()
   }
}
</script>

<style scoped lang="scss">

.hover{
  cursor: pointer;
}
.chat{
  height: 100%;
  display: flex;
  .left{
    width: 30%;
    background: white;
    height: 100%;
    overflow: auto;
    .groups{
      padding: 10px;
      overflow: auto;
      .group-row{
        border-radius: 8px;
        margin-bottom: 5px;
        padding: 0 15px;
        min-height: 75px;
        &:hover{
          cursor: pointer;
          background: #fafafa;
        }
      }
      .opened{
        background: #e1f5fe;
      }
    }
  }
  .right{
    width: 70%;
    height: 100%;
    .header{
      border-left: 1px solid darkgrey;
      border-bottom: 1px solid darkgrey;
      padding: 0 15px;
      h5{
        margin-bottom: 0;
        margin-left: 10px;
        align-items: center;
        display: flex;
      }
    }
    .chat-field{
      height: 94%;
      display: flex;
      flex-direction: column;
      .chat{
        overflow: auto;
      }
      .bottom{
        height: 100px;
        display: flex;
        justify-content: space-between;
        align-items: center;
      }
    }
  }
  .header{
    height: 60px;
    width: 100%;
    background: white;
    align-items: center;
  }
  .w-60{}
}
</style>