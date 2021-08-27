<template>
  <form action="">
    <CompanyName />
    <DeviceName />
    <DeviceSerialNumber />
    <DeviceType />
    <DeviceStatus/>
    <button class="btn btn-success">Save</button>
  </form>
</template>

<script>
import CompanyName from "@/components/formParts/CompanyName";
import DeviceName from "@/components/formParts/DeviceName";
import DeviceSerialNumber from "@/components/formParts/DeviceSerialNumber";
import DeviceType from "@/components/formParts/DeviceType";
import DeviceStatus from "@/components/formParts/DeviceStatus";

export default {
  name: "DeviceAdd",
  components: {
    CompanyName,
    DeviceName,
    DeviceSerialNumber,
    DeviceType,
    DeviceStatus

  }
}
</script>

<style scoped>

</style>