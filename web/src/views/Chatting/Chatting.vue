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
           <div class="group-row d-flex align-items-center justify-content-between"  :class="{ 'opened' : selectedID == iter.id}" @click="getRoom(iter.id)" :key="index">
              <div class="d-flex align-items-center">
                <div class="mr-3">
                  <b-avatar :src="getImgUrl(iter.avatar)"></b-avatar>
                </div>
                <div >
                  <h5 class="mb-0">{{iter.name}}</h5>
                  <p v-if="iter.not_seen_messages.length > 0" class="mb-0 red">{{iter.not_seen_messages.length}} New messages</p>
                  <div v-else>
                    <p class="grey mb-0" v-if="iter.messages.length > 0 && iter.messages[iter.messages.length - 1]">{{iter.messages[iter.messages.length - 1].message}}</p>
                  </div>
                </div>
              </div>
             <div>
               <div>
                 <p class="grey mb-0" v-if="iter.messages.length > 0 && iter.messages[iter.messages.length - 1]">{{ new Date(iter.messages[iter.messages.length - 1].sent).getHours() + ":"+ new Date(iter.messages[iter.messages.length - 1].sent).getUTCMinutes() + " " + new Date(iter.messages[iter.messages.length - 1].sent).toDateString()}}</p>
               </div>
             </div>
           </div>
         </template>
        </div>
      </div>
      <div class="right" v-if="selectedID !==0 && room">
        <div class="header d-flex justify-content-between">
          <b-button variant="white" class="d-flex align-items-center" @click="[profileOpened = !profileOpened, settingsOpened = false]">
            <b-avatar v-if="room.avatar !== ''" :src="getImgUrl(room.avatar)"></b-avatar>
            <h5 class="hove-text">{{room.name}}</h5>
          </b-button>
          <div class="d-flex">
          </div>
          <div>
            <b-button v-if="!settingsOpened && !profileOpened" @click="settingsOpened = true" variant="white">
              <b-icon icon="gear" ></b-icon>
            </b-button>
            <b-button v-if="settingsOpened && !profileOpened" @click="settingsOpened = false" variant="white">
              <b-icon icon="arrow-left" ></b-icon>
            </b-button>
            <b-button v-if="profileOpened && !settingsOpened" @click="profileOpened = false" variant="white">
              <b-icon icon="arrow-left" ></b-icon>
            </b-button>
          </div>
        </div>
        <div v-if="!settingsOpened && !profileOpened" class="chat-field">
          <ul v-chat-scroll="{always: true, smooth: true, scrollonremoved:true, smoothonremoved: false}" class="chat flex-column">
            <template v-for="(iter,index) in this.room.messages">
              <li :key="index" class="mb-3">
                <p></p>
                <div :class="[(currentUser.id == iter.user_id) ? 'personal' : null, (currentUser.id != iter.user_id) ? 'other' : null]">
                  <div class="message-field">
                    <div class="message-author">{{iter.user.name + " " + iter.user.last_name}}</div>
                    <div class="message-content">{{iter.message}}</div>
                    <div class="message-date">{{ new Date(iter.sent).getHours() + ":"+ new Date(iter.sent).getUTCMinutes() + " " + new Date(iter.sent).toDateString()}}</div>
                  </div>
                </div>
              </li>
            </template>
          </ul>
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
        <div v-if="settingsOpened" class="settings p-4">
          <div
              class="borders mt-5"
          >
            <h4>Group name:</h4>
            <b-form-input
                class="w-75"
                v-model="room.name"
                placeholder="Enter project name"
            ></b-form-input>
            <b-button @click="saveName()" variant="success" class="mt-2" v-if="room.name != ''"
            >Save</b-button
            >
          </div>
          <div class="borders mt-5">
            <div class=" editor">
              <h5>Project About:</h5>
              <b-form-textarea
                  id="textarea"
                  v-model="room.about"
                  placeholder="Enter something..."
                  rows="3"
                  max-rows="6"
              ></b-form-textarea>
              <b-button variant="success" class="mt-2" @click="saveAbout()" v-if="room.about != ''"
              >Save</b-button
              >
            </div>
          </div>
          <div class="borders mt-5">
            <h4>Group avatar:</h4>
            <b-form-file
                v-model="avatar"
                :state="Boolean(avatar)"
                placeholder="Choose a photo or drop it here..."
                drop-placeholder="Drop photo here..."
                accept="image/*"
            ></b-form-file>
            <b-button variant="success" class="mt-2" @click="saveAvatar()" v-if="avatar != null"
            >Save</b-button
            >
          </div>
          <div class="borders mt-5">
            <h5>Invite to Group:</h5>
            <vSelect
                class="w-75"
                multiple
                v-model="invitees"
                label="email"
                :options="nonMembers"
            />
            <b-button
                class="mt-2"
                variant="success"
                @click="inviteUsers()"
                v-if="invitees.length > 0"
            >Add</b-button
            >
          </div>
        </div>
        <div class="profile p-4" v-if="profileOpened">
          <div class="d-flex mb-3">
            <b-avatar v-if="room.avatar !== ''" :src="getImgUrl(room.avatar)"></b-avatar>
            <h3 class="ml-3">{{room.name}}</h3>
          </div>
          <div class="about">
            <p class="p-3">{{room.about}}</p>
          </div>
          <div class="members">

          </div>
          <b-button class="w-100" variant="danger">Leave</b-button>
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
      avatar: null,
      profileOpened:false,
      settingsOpened: false,
      selectedID: 0,
      CreatingNewGroup: false,
      connection: null,
      nonMembers: [],
      newGroup: {
        name: "",
        about: "",
        participants: [],
      },
      rooms: [],
      roomsSearched: [],
      invitees: [],
      options: [],
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
    saveName(){
      window.axios.post("api/chatting/settings/add/name",this.room).then(() => {
        this.makeToast("name update successfully", "success")
        this.getRooms()
        this.getRoom(this.selectedID)
      })
          .catch(() => {
            this.makeToast("There was a problem", "danger")
          });
    },
    saveAbout(){
      window.axios.post("api/chatting/settings/add/about",this.room).then(() => {
        this.makeToast("about update successfully", "success")
        this.getRoom(this.selectedID)
      })
          .catch(() => {
            this.makeToast("There was a problem", "danger")
          });
    },
    saveAvatar() {
      let formData = new FormData();
      formData.append("avatar", this.avatar);
      formData.append("id", this.selectedID)
      window.axios
          .post(
              `api/chatting/settings/add/avatar`,
              formData,
              {
                headers: {
                  "Content-Type": "multipart/form-data"
                }
              }
          )
          .then(() => {
            this.makeToast("Group avatar changed successfully", "success");
            this.getRooms()
            this.getRoom(this.selectedID)
          })
          .catch(rej => {
            this.makeToast(rej.response.data.error, "danger");
          });
    },
    getNonMembers() {
      window.axios.post("api/chatting/settings/get/non/members",{"id": this.selectedID}).then(res => {
        let members =  res.data.users
        if (members && members.length > 0) {
          this.nonMembers = res.data.users;
        } else {
          this.nonMembers = []
        }
      })
    },
    inviteUsers() {
      window.axios.post("api/chatting/settings/add/members", {"id": this.selectedID, "participants": this.invitees}).then(() => {
        this.makeToast("members added successfully", "success")
        this.getRooms()
        this.getRoom(this.selectedID)
        this.invitees = []
      })
          .catch(() => {
            this.makeToast("There was a problem", "danger")
          });
    },
    SendMessage(){
      window.axios.post("api/chatting/send/regular/message",{"message": this.message, "rooms_id": this.selectedID}).then(() => {
        this.message = "";
        this.connection.send("room " + this.selectedID)
      })
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
    updateRoom(id) {
      if (id == this.selectedID && id == this.room.id) {
       this.getUnreadMessages()
        this.getRooms()
      } else {
        this.getRooms()
      }
    },
    bubble(arr) {
      let len = this.roomsSearched.length;
      console.log(len)
      for (let i = 0; i < len ; i++) {
        for(let j = 0 ; j < len - i - 1; j++){
          let a = new Date(this.roomsSearched[j].updated_at)
          let b = new Date(this.roomsSearched[j + 1].updated_at)
          if (a.getTime() < b.getTime()) {
            let temp = this.roomsSearched[j];
            this.roomsSearched[j] = this.roomsSearched[j+1];
            this.roomsSearched[j + 1] = temp;
          }
        }
      }
      return arr;
    },
    async getUnreadMessages() {
      await window.axios.post("api/chatting/get/rooms/unseen/chats",{"rooms_id": this.selectedID}).then(res => {
        let arr = res.data.messages;
        if (arr) {
          this.room.messages = this.room.messages.concat(arr)
        } else {
          return []
        }
      }).catch(() => {
        return []
      })

    },
    searchRoom() {
      this.roomsSearched = [];
      for (let i = 0; this.rooms.length > i; i++) {
        if (this.searchTerm === ""|| this.rooms[i].name.includes(this.searchTerm)) {
          this.roomsSearched.push(this.rooms[i])
        }
      }
      this.bubble()
    },
    getRoom(id) {
      this.selectedID = id
      window.axios.post("api/chatting/get/room", {"id": this.selectedID}).then(res => {
        this.room = res.data.room
        this.getNonMembers()
        this.getRooms()
        this.room.newMessage = false
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

     let name = process.env.VUE_APP_API.replace("http://", "")

     this.connection = new WebSocket("ws://"+ name+ "/get/websocket")
     this.connection.onmessage = function(event) {
       console.log(event);
     }

     this.connection.onopen = function(event) {
       console.log(event)
       console.log("Successfully connected to the echo websocket server...")
     }

     let vue = this
     this.connection.onmessage = function (evt) {
       let messages = evt.data;
       messages = messages.split(" ")
       vue.updateRoom(messages[1])
     };
   },
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
      .red{
        color: red;
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
        list-style-type:none;
        li{
          width: 100%;
        }
        .personal{
          width: 50%;
          margin-left: 50%;
          display: flex;
          justify-content: flex-end;
          .message-field{
            background: #c8e6c9;
          }
        }
        .other{
          width: 50%;
          margin-right: 50%;
          display: flex;
          .message-field{
            background: white;
          }
        }

        .message-field{
          border-radius: 12px;
          padding: 10px;
          .message-date{
            font-size: 12px;
            color: #9e9e9e;
          }
          .message-author{
            color: orange;
            font-size: 13px;
            padding-bottom: 5px;
          }
        }

      }
      .bottom{
        height: 100px;
        display: flex;
        justify-content: space-between;
        align-items: center;
      }
    }
    .hove-text{
     &:hover{
       color: red;
     }
    }
  }
  .grey{
    color: #9e9e9e;
    font-size: 12px;
  }
  .header{
    height: 60px;
    width: 100%;
    background: white;
    align-items: center;
  }
  .w-60{}
  .about{
    border: 1px #9e9e9e solid;
    margin-bottom: 20px;
  }
}
</style>