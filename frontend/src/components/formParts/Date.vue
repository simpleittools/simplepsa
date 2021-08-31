<template>
  <div class="input-group">
    <span class="input-group-text">Date</span>
    <input type="date" class="form-control" v-model="selectedDate">
  </div>
</template>

<script>
export default {
  name: "Date",
  props: {
    selectedDate: Date
  }
}
</script>

<style scoped>

</style>