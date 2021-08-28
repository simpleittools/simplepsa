<template>
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
</template>

<script>
import {reactive, computed} from 'vue'
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
    PaymentInFullDue
  },
  setup() {
    const workOrderLines = reactive( [{
      lineCount: '',
      quantity: '',
      group: '',
      description: '',
      priceEach: '',
      extPrice: '',
    }])

    const sendForApproval = () => {
      console.log("Saved")
    }
    let countLines = computed(() => workOrderLines.length)


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
    return {
      sendForApproval,
      addRow,
      workOrderLines,
      countLines
    }
  }
}
</script>

<style scoped>

</style>