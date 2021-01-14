<template>
  <div class="board">
    <div class="calendar">
      <div class="calendar-controls"></div>
      <calendar-view
        :show-date="showDate"
        :items="events"
        :show-times="showTimes"
        :period-changed-callback="periodChanged"
        :starting-day-of-week="startingDayOfWeek"
        @click-item="onClickItem"
        class="theme-default holiday-us-traditional holiday-us-official"
      >
        <calendar-view-header
          slot="header"
          slot-scope="t"
          :periodStart="periodStart"
          :header-props="t.headerProps"
          @input="setShowDate"
        />
      </calendar-view>
    </div>
    <div class="info">
      <div class="left mb-5">
        <h5 class="mb-3">Add Event:</h5>
        <div class="form">
          <b-form-input
            v-model="form.title"
            placeholder="Enter your title"
          ></b-form-input>
        </div>
        <div class="form">
          <b-form-textarea
            id="textarea"
            v-model="form.description"
            placeholder="Enter event description"
            rows="3"
            max-rows="6"
          ></b-form-textarea>
        </div>
        <div class="form">
          <b-form-checkbox id="public" v-model="form.public" name="public">
            Event is public
          </b-form-checkbox>
        </div>
        <div class="form">
          <p>When will this event happen?</p>
          <date-picker
            class="w-100"
            v-model="form.time"
            type="datetime"
            range
          ></date-picker>
        </div>
        <div class="form">
          <p>Who else will participate?</p>
          <vSelect
            class="w-100"
            v-model="form.users"
            v-if="!form.public"
            label="email"
            multiple
            :options="users"
          ></vSelect>
        </div>
        <div class="form">
          <p>Which tags to add to event?</p>
          <vSelect
            class="w-100"
            v-model="form.tags"
            v-if="!form.public"
            label="name"
            multiple
            :options="tags"
          ></vSelect>
        </div>
        <b-button @click="createEvent()" variant="success"
          >Create Event</b-button
        >
      </div>
      <div class="right">
        <div v-if="item != null && !isEditing">
          <b-card :title="item.title" img-alt="Image" img-top>
            <template v-for="(iter, index) in item.originalItem.members">
              <div
                class="d-flex align-items-center"
                :key="index"
                v-if="iter.isOwner == true"
              >
                <template v-if="iter.user.id == currentUser.id">
                  {{ setItemOwner(iter.isOwner) }}
                </template>
                <b-avatar
                  class="mr-1"
                  size="2rem"
                  :src="getImgUrl(iter.user.avatar)"
                ></b-avatar>
                <p class="mb-0 text-muted">
                  {{ iter.user.name + " " + iter.user.last_name }}
                </p>
              </div>
            </template>
            <hr />
            <b-card-text>
              {{ item.originalItem.description }}
            </b-card-text>
            <template v-if="item.originalItem.members.length > 1">
              <p>Participants:</p>
              <div class="d-flex">
                <template v-for="(iter, index) in item.originalItem.members">
                  <div
                    class="d-flex align-items-center"
                    :key="index"
                    v-if="iter.isOwner == false"
                  >
                    <b-avatar
                      class="mr-1"
                      size="2rem"
                      :src="getImgUrl(iter.user.avatar)"
                    ></b-avatar>
                  </div>
                </template>
              </div>
            </template>
            <template #footer>
              <small class="text-muted"
                >{{ item.startDate + " ~ " + item.endDate }}}</small
              >
            </template>
          </b-card>
          <div class="mt-3">
            <b-button @click="editEvent()" variant="warning" class="mr-3" v-if="isItemOwner"
              >Edit</b-button
            >
            <b-button
              variant="danger"
              @click="leaveEvent(item.originalItem.id)"
              v-if="!isItemOwner && item.originalItem.public != true"
              >Leave</b-button
            >
            <b-button
              variant="danger"
              @click="deleteEvent(item.originalItem.id)"
              v-if="isItemOwner"
              >Delete</b-button
            >
          </div>
        </div>
        <div v-if="isEditing" class="info">
          <div class="left mb-5">
            <h5 class="mb-3">Edit Event:</h5>
            <div class="form">
              <b-form-input
                v-model="edit.title"
                placeholder="Enter your title"
              ></b-form-input>
            </div>
            <div class="form">
              <b-form-textarea
                id="textarea"
                v-model="edit.description"
                placeholder="Enter event description"
                rows="3"
                max-rows="6"
              ></b-form-textarea>
            </div>
            <div class="form">
              <b-form-checkbox id="public" v-model="edit.public" name="public">
                Event is public
              </b-form-checkbox>
            </div>
            <div class="form">
              <p>When will this event happen?</p>
              <date-picker
                class="w-100"
                v-model="edit.time"
                type="datetime"
                range
              ></date-picker>
            </div>
            <div class="form">
              <p>Who else will participate?</p>
              <vSelect
                class="w-100"
                v-model="edit.users"
                v-if="!edit.public"
                label="email"
                multiple
                :options="users"
              ></vSelect>
            </div>
            <div>
              <b-button @click="updateEvent(item.originalItem.id)" class="mr-3" variant="success">Save</b-button>
              <b-button @click="isEditing = false" variant="primary"
                >Back</b-button
              >
            </div>
          </div>
          <div class="right">
            <div v-if="item != null">
              <b-card :title="item.title" img-alt="Image" img-top>
                <template v-for="(iter, index) in item.originalItem.members">
                  <div
                    class="d-flex align-items-center"
                    :key="index"
                    v-if="iter.isOwner == true"
                  >
                    <template v-if="iter.user.id == currentUser.id">
                      {{ setItemOwner(iter.isOwner) }}
                    </template>
                    <b-avatar
                      class="mr-1"
                      size="2rem"
                      :src="getImgUrl(iter.user.avatar)"
                    ></b-avatar>
                    <p class="mb-0 text-muted">
                      {{ iter.user.name + " " + iter.user.last_name }}
                    </p>
                  </div>
                </template>
                <hr />
                <b-card-text>
                  {{ item.originalItem.description }}
                </b-card-text>
                <template v-if="item.originalItem.members.length > 1">
                  <p>Participants:</p>
                  <div class="d-flex">
                    <template
                      v-for="(iter, index) in item.originalItem.members"
                    >
                      <div
                        class="d-flex align-items-center"
                        :key="index"
                        v-if="iter.isOwner == false"
                      >
                        <b-avatar
                          class="mr-1"
                          size="2rem"
                          :src="getImgUrl(iter.user.avatar)"
                        ></b-avatar>
                      </div>
                    </template>
                  </div>
                </template>
                <template #footer>
                  <small class="text-muted"
                    >{{ item.startDate + " ~ " + item.endDate }}}</small
                  >
                </template>
              </b-card>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { CalendarView, CalendarViewHeader } from "vue-simple-calendar";
import DatePicker from "vue2-datepicker";
import "vue2-datepicker/index.css";
import vSelect from "vue-select";
import "vue-select/dist/vue-select.css";
// The next two lines are processed by webpack. If you're using the component without webpack compilation,
// you should just create <link> elements for these. Both are optional, you can create your own theme if you prefer.
require("vue-simple-calendar/static/css/default.css");
require("vue-simple-calendar/static/css/holidays-us.css");
export default {
  name: "Calendar",
  data: function() {
    return {
      showDate: new Date(),
      showTimes: true,
      periodStart: "",
      item: null,
      periodEnd: "",
      startingDayOfWeek: 1,
      isEditing: false,
      users: [],
      isItemOwner: false,
      events: [],
      edit: {
        title: "",
        description: "",
        public: false,
        time: "",
        users: []
      },
      tags: [],
      calendarForm: {
        public: "true",
        startDate: "",
        endDate: ""
      },
      form: {
        title: "",
        description: "",
        public: false,
        time: "",
        users: [],
        tags: []
      }
    };
  },
  components: {
    CalendarView,
    CalendarViewHeader,
    DatePicker,
    vSelect
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
    deleteEvent(id) {
      window.axios.get(`/api/calendar/delete/${id}/event`).then(() => {
        this.makeToast("Event deleted", "success");
        this.item = null;
        this.getEvents();
      });
    },
    leaveEvent(id) {
      window.axios.get(`/api/calendar/leave/${id}/event`).then(() => {
        this.makeToast("Event left", "success");
        this.item = null;
        this.getEvents();
      });
    },
    makeToast(text, variant) {
      this.$bvToast.toast(text, {
        autoHideDelay: 5000,
        variant: variant,
        title: "Notification"
      });
    },
    periodChanged(range) {
      this.calendarForm.startDate = range.displayFirstDate;
      this.calendarForm.endDate = range.displayLastDate;
      this.getEvents();
    },
    setShowDate(d) {
      this.showDate = d;
    },
    getUsers() {
      window.axios.post("api/users").then(res => {
        this.users = res.data.users;
      });
    },
    getTags() {
      window.axios.get("api/projects/gather/members/tags/list").then(res => {
        this.tags = res.data.tags;
      });
    },
    setItemOwner(iAmOwner) {
      this.isItemOwner = iAmOwner;
    },
    getEvents() {
      window.axios
        .post("api/calendar/get/current/events/", this.calendarForm)
        .then(res => {
          this.events = res.data.events;
        });
    },
    onClickItem(e) {
      this.item = e;
      this.isItemOwner = false;
    },
    editEvent() {
      console.log("test")
      this.isEditing = true;
      this.edit.title = this.item.originalItem.title;
      this.edit.description = this.item.originalItem.description;
      this.edit.public = this.item.originalItem.public;
      this.edit.time = [
        new Date(this.item.originalItem.startDate),
        new Date(this.item.originalItem.endDate)
      ];
      const localUsers = [];
      for (let iter = 0; iter < this.item.originalItem.members.length; iter++) {
        localUsers.push(this.item.originalItem.members[iter].user);
      }
      this.edit.users = localUsers;
    },
    updateEvent(id) {
      window.axios
        .post(`api/calendar/update/${id}/event`, this.edit)
        .then(() => {
          this.isEditing = false;
          this.makeToast("Event updated", "success");
          this.item = null
          this.getEvents();
        });
    },
    createEvent() {
      window.axios
        .post("api/calendar/create/new/event/", this.form)
        .then(res => {
          this.tags = res.data.tags;
          this.form = {
            title: "",
            description: "",
            public: false,
            time: "",
            users: [],
            tags: []
          };
          this.makeToast("Event created", "success");
          this.getEvents();
        });
    }
  },
  created() {
    this.getUsers();
    this.getTags();
  }
};
</script>

<style scoped lang="scss">
.board {
  width: 80%;
  margin: 20px auto;
  .calendar {
    display: flex;
    flex-direction: column;
    flex-grow: 1;
  }
  .info {
    margin-top: 30px;
    display: flex;
    .left {
      width: 40%;
      .form {
        width: 80%;
        margin-bottom: 20px;
      }
    }
    .right {
      width: 60%;
    }
  }

  .calendar {
    min-height: 600px;
  }
}
</style>
