<template>
  <div class="col-md-6 col-lg-6 offset-md-3" v-if="clientForm.displayClientForm">
    <ClientAdd />
  </div>
  <button @click="toggleClientForm">Toggle</button>
</template>

<script>
import ClientAdd from "@/components/forms/ClientAdd";
import { reactive } from "vue";


export default {
  name: 'Home',
  components: {
    ClientAdd
  },
  setup() {
    const clientForm = reactive({
      displayClientForm: true
    })
    const toggleClientForm = () => {
      clientForm.displayClientForm = !clientForm.displayClientForm
    }

    return {
      clientForm,
      toggleClientForm
    }
  }
}
</script>
