<template>
  <div class="col-md-6 col-lg-6 offset-md-3 ">
    <ClientAdd />
  </div>
</template>

<script>
import ClientAdd from "@/components/forms/ClientAdd";


export default {
  name: 'Home',
  components: {
    ClientAdd
  }
}
</script>
