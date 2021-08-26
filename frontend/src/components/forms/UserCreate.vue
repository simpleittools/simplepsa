<template>
  <form action="">
    <FirstName />
    <LastName />
    <EmailAddress />
    <Password />
    <Password2 />
  </form>
</template>

<script>
import FirstName from "@/components/formParts/FirstName";
import LastName from "@/components/formParts/LastName";
import EmailAddress from "@/components/formParts/EmailAddress";
import Password from "@/components/formParts/Password";
import Password2 from "@/components/formParts/Password2";
export default {
  name: "UserCreate",
  components: {
    FirstName,
    LastName,
    EmailAddress,
    Password,
    Password2,
  }
}
</script>

<style scoped>

</style>