<template>
  <div class="board">
    <div class="calendar">
      <div class="calendar-controls"></div>
      <calendar-view
              :show-date="showDate"
              :show-times="showTimes"
              :starting-day-of-week="startingDayOfWeek"
              class="theme-default holiday-us-traditional holiday-us-official"
      >
        <calendar-view-header
                slot="header"
                slot-scope="t"
                :header-props="t.headerProps"
                @input="setShowDate"
        />
      </calendar-view>
    </div>
    <div class="info">
      <div class="left mb-5">
        <h5 class="mb-3">Add Event:</h5>
        <div class="form">
          <b-form-input v-model="form.title" placeholder="Enter your title"></b-form-input>
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
          <b-form-checkbox
                  id="public"
                  v-model="form.public"
                  name="public"
          >
            Event is public
          </b-form-checkbox>
        </div>
        <div class="form">
          <p>When will this event happen?</p>
          <date-picker class="w-100" v-model="form.time" type="datetime" range></date-picker>
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
        <b-button @click="createEvent()" variant="success">Create Event</b-button>
      </div>
      <div class="right">Rights</div>
    </div>
  </div>
</template>

<script>
import { CalendarView, CalendarViewHeader } from "vue-simple-calendar";
import DatePicker from 'vue2-datepicker';
import 'vue2-datepicker/index.css';
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
      startingDayOfWeek: 1,
      users: [],
      tags: [],
      form: {
        title: "",
        description: "",
        public: false,
        time:"",
        users: [],
        tags: [],
      }
    };
  },
  components: {
    CalendarView,
    CalendarViewHeader,
    DatePicker,
    vSelect
  },
  methods: {
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
    createEvent() {
      window.axios.post("api/calendar/create/new/event/", this.form).then(res => {
        this.tags = res.data.tags;
      });
    },
  },
  created() {
    this.getUsers()
    this.getTags()
  }
};
</script>

<style scoped lang="scss">
.board{
  width: 80%;
  margin: 20px auto;
  .calendar {
    display: flex;
    flex-direction: column;
    flex-grow: 1;
  }
  .info{
    margin-top: 30px;
    display: flex;
    .left{
      width: 40%;
      .form{
        width: 80%;
        margin-bottom: 20px;
      }
    }
    .right{
      width: 60%;
    }
  }
}
</style>
