<template>
  <div class="input-group">
    <span class="input-group-text">Company Initials:</span>
    <input type="text" class="form-control" @input="$emit('update:modelValue', $event.target.value)">
  </div>
</template>

<script>
import {ref} from "vue";

export default {
  name: "CompanyName",
  setup(){
    let companyInitialInput = ref()

    return {
      companyInitialInput
    }
  }
}
</script>

<style scoped>

</style>