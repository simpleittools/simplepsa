<template>
  <h1>Work Orders</h1>
  <button class="btn btn-primary" @click="manageWorkOrder">New Work Order</button>
  <div v-if="workOrderViews.showWorkOrderStatus">
    <WorkOrderStatus />
  </div>
  <div v-else-if="workOrderViews.showNewWorkOrder">
    <WorkOrderCreate />
  </div>
</template>

<script>
import WorkOrderStatus from "@/components/WorkOrderStatus";
import WorkOrderCreate from "@/components/forms/WorkOrderCreate";
import {reactive} from 'vue'
export default {
  name: "WorkOrders",
  components: {
    WorkOrderStatus,
    WorkOrderCreate
  },
  setup() {
    const workOrderViews = reactive({
      showWorkOrderStatus: true,
      showNewWorkOrder: false
    })

    const manageWorkOrder = () => {
      workOrderViews.showWorkOrderStatus = !workOrderViews.showWorkOrderStatus
      workOrderViews.showNewWorkOrder = !workOrderViews.showNewWorkOrder
    }

    return {
      workOrderViews,
      manageWorkOrder
    }
  }
}
</script>

<style scoped>

</style>