<template>
    <div>
        <Navbar class="navbar"/>
        <Home class="home"/>
    </div>
</template>
<script>
import css from "../assets/normalize.css";
import Navbar from "./items/Navbar.vue";
import Home from "./Home";
export default {
  name: "Layout",
  components: {
    Navbar,
    Home
  }
};
</script>
<style scoped>
.navbar {
  padding: 0 1rem;
  display: -webkit-box;
  text-align: center;
  background-color: whitesmoke;
}
.home {
  padding: 1rem;
}
</style>


