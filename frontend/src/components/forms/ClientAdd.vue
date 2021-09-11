<template>
  <form @submit.prevent="handleSubmit">
    <CompanyName v-model="clientDetails.name" />
    <CompanyInitials v-model="clientDetails.initials" />
    <button class="btn btn-success">Add</button>
  </form>
</template>

<script>
import {ref, reactive} from "vue";
import CompanyName from "@/components/formParts/CompanyName";
import CompanyInitials from "@/components/formParts/CompanyInitials";
import axios from "axios";

import { useRouter } from 'vue-router'


export default {
  name: "ClientAdd",
  components: {
    CompanyName,
    CompanyInitials
  },
  setup() {
    let clientDetails = reactive({
      name: "",
      initials: ""
    })
    const clientInitialsInput = ref("")

    const isPending = ref(false)
    const router = useRouter()

    const reset = () => {
      clientDetails.name = ""
      clientDetails.initials = ""
    }

    const handleSubmit = async () => {
      await axios.post("client/create",{
        client_name: clientDetails.name,
        client_initials: clientDetails.initials
      })
      await router.push({name: "About"})
    }

    return {
      clientDetails,
      clientInitialsInput,
      handleSubmit
    }

  }
}
</script>

<style scoped>

</style>