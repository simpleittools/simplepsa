<template>
  <form action="">
    <DeviceName />
    <DeviceSerialNumber />
    <button class="btn btn-success">Save</button>
  </form>
</template>

<script>
import DeviceName from "@/components/formParts/DeviceName";
import DeviceSerialNumber from "@/components/formParts/DeviceSerialNumber";
export default {
  name: "DeviceAdd",
  components: {
    DeviceName,
    DeviceSerialNumber
  }
}
</script>

<style scoped>

</style>