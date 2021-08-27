<template>
  <form action="">
    <div class="row">
      <WorkOrderNumber /> <Date />
    </div>

  </form>
</template>

<script>
import WorkOrderNumber from "@/components/formParts/WorkOrderNumber";
import Date from "@/components/formParts/Date";

export default {
  name: "WorkOrderCreate",
  components: {
    WorkOrderNumber,
    Date
  }
}
</script>

<style scoped>

</style>