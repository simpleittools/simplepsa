<template>
  <div class="input-group">
    <span class="input-group-text">Company:</span>
    <input type="text" class="form-control" @input="$emit('update:modelValue', $event.target.value)">
  </div>
</template>

<script>
import {ref} from 'vue';
export default {
  name: "CompanyName",
  setup(){
    let companyNameInput = ref()

    return {
      companyNameInput
    }
  }
}
</script>

<style scoped>

</style>