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
        <div v-if="item != null">
          <b-card :title="item.title" img-alt="Image" img-top>
            <template v-for="(iter,index) in item.originalItem.members">
              <div class="d-flex align-items-center" :key="index" v-if="iter.isOwner == true">
                <template v-if="iter.user.id == currentUser.id">
                  {{setItemOwner(iter.isOwner)}}
                </template>
                <b-avatar class="mr-1" size="2rem" :src="getImgUrl(iter.user.avatar)"></b-avatar>
                <p class="mb-0 text-muted">{{iter.user.name + " " + iter.user.last_name}}</p>
              </div>
            </template>
            <hr>
            <b-card-text>
              {{item.originalItem.description}}
            </b-card-text>
            <template v-if="item.originalItem.members.length > 1">
              <p>Participants:</p>
              <div class="d-flex">
                <template v-for="(iter,index) in item.originalItem.members">
                  <div class="d-flex align-items-center" :key="index" v-if="iter.isOwner == false">
                    <b-avatar class="mr-1" size="2rem" :src="getImgUrl(iter.user.avatar)"></b-avatar>
                  </div>
                </template>
              </div>
            </template>
            <template #footer>
              <small class="text-muted">{{item.startDate + " ~ " + item.endDate}}}</small>
            </template>
          </b-card>
          <div class="mt-3">
            <b-button variant="warning" class="mr-3" v-if="isItemOwner">Edit</b-button>
            <b-button variant="danger" v-if="!isItemOwner && item.originalItem.public != true">Leave</b-button>
            <b-button variant="danger" v-if="isItemOwner">Delete</b-button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {
  CalendarView,
  CalendarViewHeader
} from "vue-simple-calendar";
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
      users: [],
      isItemOwner: false,
      events: [],
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
    periodChanged(range) {
      this.calendarForm.startDate = range.displayFirstDate
      this.calendarForm.endDate = range.displayLastDate
      this.getEvents()
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
      this.isItemOwner = iAmOwner
    },
    getEvents() {
      window.axios
        .post("api/calendar/get/current/events/", this.calendarForm)
        .then(res => {
          this.events = res.data.events;
        });
    },
    onClickItem(e) {
      this.item = e
      this.isItemOwner = false
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
          this.getEvents()
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

  .calendar{
    min-height: 600px;
  }
}
</style>
