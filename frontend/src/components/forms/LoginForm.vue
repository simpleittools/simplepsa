<template>
  <form action="">
    <EmailAddress />
    <Password />
    <button class="btn btn-success">Log In</button>
  </form>
</template>

<script>
  import EmailAddress from "@/components/formParts/EmailAddress";
  import Password from "@/components/formParts/Password";
  export default {
    name: "LoginForm",
    components: {
      EmailAddress,
      Password
    }
  }
</script>

<style scoped>

</style>