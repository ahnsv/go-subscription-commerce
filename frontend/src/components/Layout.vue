<template>
    <div>
        <Navbar class="navbar"/>
        <Home class="home"/>
    </div>
</template>
<script>
import Navbar  from "./items/Navbar.vue";
import Home from "./Home";
export default {
  name: "Layout",
  components: {
    Navbar,
    Home
  }
};
</script>

