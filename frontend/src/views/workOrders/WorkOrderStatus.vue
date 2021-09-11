<template>
  <h1>Work Order Status</h1>
  <table class="table">
    <thead>
    <th>WO#</th>
    <th>Status</th>
    <th>Client</th>
    <th>Project Name</th>
    <th>POC</th>
    <th>Created By</th>
    <th>Date Created</th>
    <th>Date Sent</th>
    </thead>
    <tbody>
    <tr v-for="workOrder in workOrders">
<!--      {{workOrder}}-->
      <td>{{workOrder.work_order_name}}</td>
      <td>Status</td>
      <td>{{workOrder.client.client_name}}</td>
      <td>{{workOrder.project_name}}</td>
      <td>{{workOrder.requestor}}</td>
      <td>User</td>
      <td>{{workOrder.work_order_date}}</td>
      <td>Date Sent</td>
    </tr>


    </tbody>
  </table>
</template>

<script>
import getWorkOrders from "@/composables/getWorkOrders";
export default {
  name: "WorkOrderStatus",
  setup() {
    const {workOrders, error, load} = getWorkOrders()
    load()
    return {workOrders, error }
  }
}
</script>

<style scoped>

</style>