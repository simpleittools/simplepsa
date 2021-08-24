<template>
  <button @click="showClientForm" v-show="!clientForm.displayClientForm">New Client</button>
  <button @click="showNewDeviceForm" v-show="!clientForm.displayNewDeviceForm">New Device</button>
  <div class="col-md-6 col-lg-6 offset-md-3">

    <transition name="slide">
      <div v-if="clientForm.displayClientForm">
        <ClientAdd />
      </div>
    </transition>
    <transition name="slide">
      <div v-if="clientForm.displayNewDeviceForm">
        <DeviceAdd />
      </div>
    </transition>

  </div>

</template>

<script>
import ClientAdd from "@/components/forms/ClientAdd";
import { reactive } from "vue";
import DeviceAdd from "@/components/forms/DeviceAdd";


export default {
  name: 'Home',
  components: {
    DeviceAdd,
    ClientAdd
  },
  setup() {
    const clientForm = reactive({
      displayClientForm: true,
      displayNewDeviceForm: false
    })
    const showNewDeviceForm = () => {
      clientForm.displayClientForm = false
      clientForm.displayNewDeviceForm = true
    }
    const showClientForm = () => {
      clientForm.displayClientForm = true
      clientForm.displayNewDeviceForm = false
    }

    return {
      clientForm,
      showNewDeviceForm,
      showClientForm
    }
  }
}
</script>

<style>
body {
  margin: 30px;
}

.slide-enter-active {
  transition: all 2s ease-out;
}

.slide-leave-active {
  transition: all 1s cubic-bezier(1.0, 0.5, 0.8, 1.0);
}

.slide-enter-from,
.slide-leave-to {
  transform: translateY(-50%);
  opacity: 0;
}

</style>
