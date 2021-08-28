<template>
  <label for="invoice_delivery_select" class="radio_select_label">Invoice Method:</label>
  <div class="input-group" id="invoice_delivery_select">
    <div class="input-group-text">
      <input type="radio" class="form-check-input" aria-label="Net 30" id="net_thirty" name="payment_due" value="true">
      <label for="net_thirty" class="form-check-label">&nbsp;Net 30</label>
    </div>
    <div class="input-group-text">
      <input type="radio" class="form-check-input" aria-label="Immediate Payment" id="payment_upfront" name="payment_due" value="false">
      <label for="payment_upfront" class="form-check-label">&nbsp;Immediate</label>
    </div>
  </div>
</template>

<script>
export default {
  name: "PaymentInFullDue"
}
</script>

<style scoped>

</style>