<template>
    <div class="calendar">
        <calendar-view
                :show-date="showDate"
                class="theme-default holiday-us-traditional holiday-us-official">
            <calendar-view-header
                    slot="header"
                    slot-scope="t"
                    :header-props="t.headerProps"
                    @input="setShowDate" />
        </calendar-view>
    </div>
</template>

<script>
    import { CalendarView, CalendarViewHeader } from "vue-simple-calendar"
    // The next two lines are processed by webpack. If you're using the component without webpack compilation,
    // you should just create <link> elements for these. Both are optional, you can create your own theme if you prefer.
    require("vue-simple-calendar/static/css/default.css")
    require("vue-simple-calendar/static/css/holidays-us.css")
    export default {
        name: "Calendar",
        data: function() {
            return { showDate: new Date() }
        },
        components: {
            CalendarView,
            CalendarViewHeader,
        },
        methods: {
            setShowDate(d) {
                this.showDate = d;
            },
        }
    }
</script>

<style scoped>
.calendar{
    display: flex;
    flex-direction: column;
    flex-grow: 1;
    width: 80%;
    margin: 20px auto;
}
</style>