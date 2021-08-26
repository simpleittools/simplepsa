<template>

  <div id="nav">
    <NavBar />
<!--    <router-link to="/">Home</router-link> |-->
<!--    <router-link to="/about">About</router-link>-->
  </div>
  <div class="container">
    <router-view/>
  </div>

</template>

<script>
import NavBar from "@/components/NavBar";

export default {
  components: {
    NavBar
  }
}

</script>

<style>
.input-group-text {
  width: 130px;
}

</style>
