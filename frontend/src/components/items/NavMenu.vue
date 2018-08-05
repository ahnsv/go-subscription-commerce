<template>
    <div>
        <a src="#"> <div class="navmenu--trending">Trending</div> </a>
        <a src="#"> <div class="navmenu--products">Products</div> </a>
        <a src="#"> <div class="navmenu--home">Go commerce</div> </a>
        <a src="#"> <div class="navmenu--explore">Explore</div> </a>
        <a src="#"> <div class="navmenu--blog">Blog</div> </a>
    </div>
</template>
<script>
export default {
    name: 'NavMenu',
}

</script>
<style scoped>
.navmenu {
    display: inline-flex;
    padding: 1rem;
}
.navmenu a {
    margin-right: 1rem;
}
</style>
