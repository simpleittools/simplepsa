<template>
  <div class="col-md-6 col-lg-6 offset-md-3">
    <transition name="slide-fade">
      <div v-if="clientForm.displayClientForm">
        <ClientAdd />
      </div>
    </transition>

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

<style>
body {
  margin: 30px;
}

.slide-fade-enter-active {
  transition: all .3s ease-out;
}

.slide-fade-leave-active {
  transition: all .8s cubic-bezier(1.0, 0.5, 0.8, 1.0);
}

.slide-fade-enter-from,
.slide-fade-leave-to {
  transform: translateY(-20px);
  opacity: 0;
}

</style>
