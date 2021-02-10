<template>
  <div class="mt-3 w-75 ml-auto mr-auto">
    <b-tabs content-class="mt-3" fill>
      <b-tab title="Working Day Table" class="" active>
        <div class="d-flex justify-content-between mb-3">
          <div></div>
          <b-button
            @click="saveTimeTable()"
            v-if="showSave && !isSavedResults"
            variant="outline-success"
            >Save</b-button
          >
          <b-button
              @click="updateTimeTable()"
              v-if="showSave && isSavedResults"
              variant="outline-success"
          >Update</b-button
          >
        </div>
        <div class=" d-flex justify-content-center">
          <h5>Choose a month:</h5>
        </div>
        <div class="mb-4 d-flex justify-content-center">
          <month-picker-input @change="showDate"></month-picker-input>
        </div>
        <template v-for="(iter, index) in dates">
          <div :key="index" class="row-date">
            <div class="d-flex justify-content-between">
              <h6>{{ new Date(iter.date).toDateString() }}</h6>
              <b-form-select
                class="w-25"
                @change="statusChanged(index)"
                v-model="iter.status"
                :options="options"
              ></b-form-select>
            </div>
            <div v-if="iter.status === '1'">
              <vue-timepicker
                v-model="iter.start_time"
                placeholder="Start Time"
              ></vue-timepicker>
              <vue-timepicker
                class="ml-3"
                v-model="iter.end_time"
                placeholder="End Time"
              ></vue-timepicker>
            </div>
            <hr />
          </div>
        </template>
      </b-tab>
      <b-tab title="Working Schedule">
       <div class="d-flex justify-content-between flex-wrap">
         <div class="row-user" v-for="iter in datesAll" :key="iter.id">
           <b-card img-alt="Image" img-top>
             <div
                 class="d-flex align-items-center"
             >
               <b-avatar
                   class="mr-1"
                   size="2rem"
                   :src="getImgUrl(iter.user.avatar)"
               ></b-avatar>
               <p class="mb-0 text-muted">
                 {{ iter.user.name + " " + iter.user.last_name }}
               </p>
             </div>
             <hr />
             <div class="d-flex" v-for="(liter, index) in options" :key="index">
               <template v-if="liter.value === iter.status">
                 <h5>{{liter.text}}</h5>
                 <h5 v-if="iter.status == 1">: {{iter.start_time}} - {{iter.end_time}}</h5>
               </template>
             </div>
           </b-card>
         </div>
       </div>
      </b-tab>
    </b-tabs>
  </div>
</template>

<script>
import VueTimepicker from "vue2-timepicker";
import "vue2-timepicker/dist/VueTimepicker.css";
import { MonthPickerInput } from "vue-month-picker";

export default {
  components: { VueTimepicker, MonthPickerInput },
  name: "TimeTables",
  data() {
    return {
      dates: [],
      showSave: true,
      isInit: false,
      isSavedResults: false,
      start_date: null,
      end_date: null,
      datesAll: [],
      date: {
        from: null,
        to: null,
        month: null,
        year: null
      },
      options: [
        { value: "", text: "Not working" },
        { value: "1", text: "Working on this day" },
        { value: "S", text: "Sick leave" },
        { value: "A", text: "Vacation" },
        { value: "BA", text: "Vacation without pay" },
        { value: "MA", text: "Study vacation" },
        { value: "DD", text: "Holiday for blood donors" }
      ]
    };
  },
  methods: {
    getImgUrl(image) {
      let images = process.env.VUE_APP_API + "/static" + image;
      return images;
    },
    statusChanged(index) {
      if (this.dates[index].status !== "1") {
        return
      }
      if (this.dates[index].start_time.HH == "") {
        this.dates[index].start_time.HH = "09"
      }
      if (this.dates[index].start_time.mm == "") {
        this.dates[index].start_time.mm = "00"
      }
      if (this.dates[index].end_time.HH == "") {
        this.dates[index].end_time.HH = "18"
      }
      if (this.dates[index].end_time.mm == "") {
        this.dates[index].end_time.mm = "00"
      }
    },
    showDate(date) {
      this.date = date;
      if (this.isInit) {
        let month = date.to.getMonth();
        if (month == 0) {
          month = 12;
        }
        this.getDaysInMonth(month - 1, date.to.getFullYear());
      } else {
        this.isInit = true;
        this.getDaysInMonth(new Date().getMonth(), new Date().getFullYear());
      }
      this.setIfToShowSaveButton();
    },
    setIfToShowSaveButton() {
      if (this.date.to.getFullYear() >= new Date().getFullYear()) {
        if (
          (this.date.to.getFullYear() === new Date().getFullYear() &&
            this.date.to.getMonth() >= new Date().getMonth()) ||
          this.date.to.getFullYear() !== new Date().getFullYear()
        ) {
          this.showSave = true;
          return;
        }
        this.showSave = false;
        return;
      }
      this.showSave = false;
    },
    getDaysInMonth(month, year) {
      this.start_date = null
      this.end_date = null
      let date = new Date(year, month, 1);
      let days = [];
      while (date.getMonth() === month) {
        if (!this.start_date) {
          this.start_date = new Date(date)
          console.log(this.start_date)
        }
        days.push({
          date: new Date(date),
          status: "",
          start_time: { HH: "09", mm: "00" },
          end_time: { HH: "18", mm: "00" }
        });
        this.end_date = new Date(date)
        date.setDate(date.getDate() + 1);
      }
      this.dates = days;
      this.getTimetable()
    },
    saveTimeTable() {
      window.axios.post(`/api/timetable/save/schedule`, {"dates":  this.dates}).then(() => {
        this.makeToast("Table saved", "success");
        this.getTimetableAll()
      });
    },
    updateTimeTable() {
      window.axios.post(`/api/timetable/update/schedule`, {"dates":  this.dates}).then(() => {
        this.makeToast("Table saved", "success");
        this.getTimetableAll()
      });
    },
    getTimetable() {
      window.axios.post(`/api/timetable/get/personal/schedule`, {"start_date":  this.start_date, "end_date": this.end_date}).then(res => {
        if(this.dates <= res.data.dates) {
          this.dates = res.data.dates
          this.isSavedResults = true
        } else {
          this.isSavedResults = false
        }
      });
    },
    getTimetableAll() {
      window.axios.post(`/api/timetable/get/everyone/schedule`, {"start_date": new Date(new Date().toDateString())}).then(res => {
          this.datesAll = res.data.dates
      });
    },
    makeToast(text, variant) {
      this.$bvToast.toast(text, {
        autoHideDelay: 5000,
        variant: variant,
        title: "Notification"
      });
    },
  },
  created() {
    const date = {
      from: new Date(),
      to: new Date(),
      month: new Date().getUTCMonth(),
      year: new Date().getUTCFullYear()
    };
    this.date = date;
    this.getTimetableAll()
  }
};
</script>
<style scoped>
.row-user{
  width: 49%;
  margin-bottom: 25px;
}
</style>
