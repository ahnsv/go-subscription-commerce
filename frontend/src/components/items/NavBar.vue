<template>
    <div>
    <div class="navbar--btns">
        <a src="#"><i class="fas fa-bars"></i></a>
        |
        <a src="#"><i class="fas fa-search"></i></a>
        <form>
            <input type="text">
            <button type="submit">Search</button>
        </form>
    </div>
    <Sidebar class="menu" />
    <NavMenu class="navmenu" />
    <div class="navbar--user">
        <div class="user-cart">
            <i class="fas fa-shopping-cart"></i>
        </div>
        <div class="user-avatar"></div>
    </div>
    </div>
</template>
<script>
import Sidebar from "./Sidebar";
import NavMenu from "./NavMenu";
export default {
  name: "NavBar",
  components: {
    Sidebar,
    NavMenu
  }
};
</script>
<style>
.menu {
  display: hide;
}
.navbar--user,
.navbar--btns {
  display: -webkit-inline-box;
  padding: 1rem;
}
.navbar--btns a {
  margin-right: 0.5rem;
}
.navbar--btns form {
    display: none;
}
</style>
