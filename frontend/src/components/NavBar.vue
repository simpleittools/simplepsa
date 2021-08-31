<template>
  <nav class="navbar navbar-expand-lg navbar-light bg-light">
    <div class="container-fluid">
      <a class="navbar-brand" href="#">Navbar</a>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNavDropdown" aria-controls="navbarNavDropdown" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarNavDropdown">
        <ul class="navbar-nav">
          <li class="nav-item">
            <router-link class="nav-link active" aria-current="page" :to="{name: 'Home'}">Home</router-link>
          </li>
          <li class="nav-item">
            <router-link class="nav-link" :to="{name: 'About'}">About</router-link>
          </li>
          <li class="nav-item">
            <router-link class="nav-link" :to="{name: 'Login'}">Login</router-link>
          </li>
          <li class="nav-item">
            <router-link class="nav-link" :to="{name: 'Admin'}">Admin</router-link>
          </li>
          <li class="nav-item dropdown">
            <a class="nav-link dropdown-toggle" href="#" id="workOrder" role="button" data-bs-toggle="dropdown" aria-expanded="false">
              WorkOrders
            </a>
            <ul class="dropdown-menu" aria-labelledby="navbarDropdownMenuLink">
              <li><router-link class="dropdown-item" :to="{name: 'WorkOrderStatus'}">Work Order Status</router-link></li>
              <li><router-link class="dropdown-item" :to="{name: 'WorkOrderCreate'}">Work Order Create</router-link></li>
              <li><a class="dropdown-item" href="#">Something else here</a></li>
            </ul>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  name: "NavBar"
}
</script>

<style scoped>

</style>