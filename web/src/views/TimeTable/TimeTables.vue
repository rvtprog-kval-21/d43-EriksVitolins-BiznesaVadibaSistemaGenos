<template>
  <div class="mt-3 w-75 ml-auto mr-auto">
    <b-tabs content-class="mt-3" fill>
      <b-tab title="Working Day Table" class="" active>
        <div class="d-flex justify-content-between mb-3">
          <div></div>
          <b-button @click="saveTimeTable()" v-if="showSave" variant="outline-success">Save</b-button>
        </div>
        <div class=" d-flex justify-content-center">
          <h5>Choose a month:</h5>
        </div>
        <div class="mb-4 d-flex justify-content-center">
          <month-picker-input @change="showDate"></month-picker-input>
        </div>
        <template v-for="(iter,index) in dates">
          <div :key="index" class="row-date">
           <div class="d-flex justify-content-between">
             <h6>{{iter.date.toDateString()}}</h6>
             <b-form-select class="w-25" v-model="iter.status" :options="options"></b-form-select>
           </div>
            <div v-if="iter.status === '1'">
              <vue-timepicker v-model="iter.start_time" placeholder="Start Time"></vue-timepicker>
              <vue-timepicker class="ml-3" v-model="iter.end_time" placeholder="End Time"></vue-timepicker>
            </div>
            <hr>
          </div>
        </template>
      </b-tab>
      <b-tab title="Working Hours"><p>I'm the second tab</p></b-tab>
    </b-tabs>
  </div>
</template>

<script>
import VueTimepicker from 'vue2-timepicker'
import 'vue2-timepicker/dist/VueTimepicker.css'
import { MonthPickerInput } from 'vue-month-picker'

export default {
  components: { VueTimepicker, MonthPickerInput },
  name: "TimeTables",
  data() {
    return {
     dates: [],
      showSave: true,
      isInit: false,
      date: {
        from: null,
        to: null,
        month: null,
        year: null
      },
      options: [
        {value: null, text: 'Not working'},
        {value: "1", text: 'Working on this day'},
        {value: "S", text: 'Sick leave'},
        {value: "A", text: 'Vacation'},
        {value: "BA", text: 'Vacation without pay'},
        {value: "MA", text: 'Study vacation'},
        {value: "DD", text: 'Holiday for blood donors'},
      ]
    };
  },
  methods: {
    showDate (date) {
      this.date = date
      if (this.isInit) {
        let month = date.to.getMonth()
        if (month == 0) {
          month = 12
        }
        this.getDaysInMonth(month - 1, date.to.getFullYear())
      } else {
        this.isInit = true
        this.getDaysInMonth(new Date().getMonth() ,new Date().getFullYear())
      }
      this.setIfToShowSaveButton()
    },
    setIfToShowSaveButton() {
      if (this.date.to.getFullYear() >= new Date().getFullYear()){
        if ((this.date.to.getFullYear() === new Date().getFullYear() && this.date.to.getMonth() >= new Date().getMonth()) || this.date.to.getFullYear() !== new Date().getFullYear()) {
          this.showSave = true
          return
        }
        this.showSave = false
        return
      }
      this.showSave = false
    },
    saveTimeTable(){

    },
    getDaysInMonth(month, year) {
    let date = new Date(year, month, 1);
    let days = [];
    while (date.getMonth() === month) {
      days.push({"date": new Date(date), "status": null, "start_time": {"HH": "09", "mm": "00"}, "end_time": {"HH": "18", "mm": "00"}});
      date.setDate(date.getDate() + 1);
    }
    this.dates = days
  },
  },
  created() {
    const date = {"from": new Date, "to": new Date,"month": new Date().getUTCMonth() , "year": new Date().getUTCFullYear()}
    this.date = date
  }
};
</script>

<style scoped></style>
