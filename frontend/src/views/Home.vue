<template>
  <button @click="showClientForm" v-show="!formSelector.displayClientForm" class="btn btn-primary">New Client</button>
  <button @click="showNewDeviceForm" v-show="!formSelector.displayNewDeviceForm" class="btn btn-primary">New Device</button>
  <button @click="showWalkInForm" v-show="!formSelector.displayNewWalkInForm" class="btn btn-primary">New Walk In</button>
  <div class="col-md-6 col-lg-6 offset-md-3">

<!--    <transition name="slide">-->
      <div v-if="formSelector.displayNewWalkInForm">
        <WalkInAdd />
      </div>
<!--    </transition>-->
<!--    <transition name="slide">-->
      <div v-if="formSelector.displayNewDeviceForm">
        <DeviceAdd />
      </div>
<!--    </transition>-->
    <div v-if="formSelector.displayClientForm">
      <ClientAdd />
    </div>

  </div>

</template>

<script>
import ClientAdd from "@/components/forms/ClientAdd";
import { reactive } from "vue";
import DeviceAdd from "@/components/forms/DeviceAdd";
import WalkInAdd from "@/components/forms/WalkInAdd";


export default {
  name: 'Home',
  components: {
    DeviceAdd,
    ClientAdd,
    WalkInAdd
  },
  setup() {
    const formSelector = reactive({
      displayClientForm: true,
      displayNewDeviceForm: false,
      displayNewWalkInForm: false,

    })
    const showNewDeviceForm = () => {
      formSelector.displayClientForm = false
      formSelector.displayNewDeviceForm = true
      formSelector.displayNewWalkInForm = false
    }
    const showClientForm = () => {
      formSelector.displayClientForm = true
      formSelector.displayNewDeviceForm = false
      formSelector.displayNewWalkInForm = false
    }
    const showWalkInForm = () => {
      formSelector.displayClientForm = false
      formSelector.displayNewDeviceForm = false
      formSelector.displayNewWalkInForm = true
    }

    return {
      formSelector,
      showNewDeviceForm,
      showClientForm,
      showWalkInForm
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
