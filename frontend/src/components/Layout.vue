<template>
    <div>
        <Navbar :class="navbar"/>
        <Home :class="home"/>
    </div>
</template>
<script>
import { Navbar } from "@/components/items/Navbar.vue";
import { Home } from "@/components/Home";
export default {
  name: "Layout",
  components: {
    Navbar,
    Home
  }
};
</script>

