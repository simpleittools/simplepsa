<template>
  <h1>Create New Work Order</h1>
  <form @submit.prevent="sendForApproval">
    <WorkOrderNumber />
    <Date />
    <ClientAFE />
    <CompanyName />
    <ClientRequestor />
    <ProjectName />
    <PaymentInFullDue />


    <ol>
      <li>Reason <br>
        <textarea name="" id="" cols="30" rows="10"></textarea>
      </li>
      <li>Scope of work <br>
        <textarea name="" id="" cols="30" rows="10"></textarea>
      </li>
      <li>Line Item Estimates <br>
        <table class="table">
          <thead>
          <th>Item #</th>
          <th>QTY</th>
          <th>Group</th>
          <th>Description</th>
          <th>Price Each</th>
          <th>Ext. Price</th>
          </thead>
          <tbody>
          <tr v-for="(workOrderLine, k) in workOrderLines" :key="k">
            <!--              todo: get line number, currently getting count of lines-->
            <td>{{countLines}}</td>
            <td><input type="number" v-model="workOrderLine.quantity"></td>
            <td><input type="text"></td>
            <td><input type="text"></td>
            <td><input type="text"></td>
            <td><input type="text"></td>
            <button type="button" class="btn btn-info" @click="addRow">+</button>
          </tr>
          </tbody>
        </table>
      </li>
      <li>Deliverables <br>
        <textarea name="" id="" cols="30" rows="10"></textarea>
      </li>
    </ol>

    <button class="btn btn-primary">Save</button>
    <!--    TODO: launch modal asking to Save As New - Editing, or Send for Internal Approval -->
  </form>

  <Modal v-model="saveWorkOrder.isOpen" max-width="500px">
    <div class="input-group">
      <button class="btn btn-success" @click="saveForEdit">Save for Edit</button>
      <button class="btn btn-primary" @click="submitForReview">Submit for review</button>
    </div>
  </Modal>

  <Modal v-model="sendReviewTo.isOpen" max-width="500px">
    <div class="input-group">
      <input type="radio" name="sendReviewBy"> Position
      <input type="radio" name="sendReviewBy"> Person
    </div>
  </Modal>
</template>

<script>
import {reactive, computed, ref} from 'vue';
import {Modal} from 'vue-neat-modal';
import {useRouter} from 'vue-router'
import WorkOrderNumber from "@/components/formParts/WorkOrderNumber";
import Date from "@/components/formParts/Date";
import ClientAFE from "@/components/formParts/ClientAFE";
import CompanyName from "@/components/formParts/CompanyName";
import ClientRequestor from "@/components/formParts/ClientRequestor";
import ProjectName from "@/components/formParts/ProjectName";
import PaymentInFullDue from "@/components/formParts/PaymentInFullDue";

export default {
  name: "WorkOrderCreate",
  components: {
    WorkOrderNumber,
    Date,
    ClientAFE,
    CompanyName,
    ClientRequestor,
    ProjectName,
    PaymentInFullDue,
    Modal
  },
  setup() {
    const router = useRouter()
    const workOrderLines = reactive( [{
      lineCount: '',
      quantity: '',
      group: '',
      description: '',
      priceEach: '',
      extPrice: '',
    }])

    const sendForApproval = () => {
      saveWorkOrder.isOpen = true
    }

    const sendReviewTo = reactive({
      isOpen: false
    })

    let countLines = computed(() => workOrderLines.length)

    const saveWorkOrder = reactive({
      isOpen: false
    })

    const addRow = () => {
      workOrderLines.push({
        lineCount: '',
        quantity: '',
        group: '',
        description: '',
        priceEach: '',
        extPrice: '',
      })
      console.log(countLines.value)
    }

    const saveForEdit = () => {
      saveWorkOrder.isOpen = false
      router.push({name: "WorkOrderStatus"})

    }
    const submitForReview = () => {
      console.log("Submit for Review")
      saveWorkOrder.isOpen = false
      sendReviewTo.isOpen = true
    }
    return {
      sendForApproval,
      addRow,
      saveForEdit,
      submitForReview,
      workOrderLines,
      countLines,
      saveWorkOrder,
      sendReviewTo,
    }
  }
}
</script>

<style scoped>

</style>