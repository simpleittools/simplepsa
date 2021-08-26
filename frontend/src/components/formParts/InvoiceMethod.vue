<template>
  <label for="invoice_delivery_select" class="radio_select_label">Invoice Method:</label>
  <div class="input-group" id="invoice_delivery_select">
    <div class="input-group-text">
      <input type="radio" class="form-check-input" aria-label="Email" id="invoice_email" name="invoice_delivery" v-model="invoiceMethodEmail" value="true">
      <label for="invoice_email" class="form-check-label">&nbsp;Email</label>
    </div>
    <div class="input-group-text">
      <input type="radio" class="form-check-input" aria-label="Email" id="invoice_snail_mail" name="invoice_delivery" v-model="invoiceMethodEmail" value="false">
      <label for="invoice_email" class="form-check-label">&nbsp;Mail</label>
    </div>
  </div>
</template>

<script>
  import {ref} from "vue";

  export default {
    name: "InvoiceMethod",
    setup() {
      const invoiceMethodEmail = ref()
      console.log(invoiceMethodEmail)

      return {
        invoiceMethodEmail
      }
    }
  }
</script>

<style scoped>
.radio_select_label {
  background-color: #E9ECEF;
  display: flex;
  align-items: center;
  padding: .375rem .75rem;
  font-size: 1rem;
  font-weight: 400;
  line-height: 1.5;
  width: 130px;
  color: #212529;
  text-align: center;
  white-space: nowrap;
  border: 1px solid #ced4da;
  border-radius: .25rem;
}
</style>