<template>
    <div>
        <div class="products-title">
            <h4>Products</h4>
        </div>
        <div class="products-view-wrapper">
            <ProductView :class="product-view" v-for="i in 10" :key="i"/>
        </div>
    </div>
</template>
<script>
import ProductView from "./items/ProductView";
export default {
  name: 'Home',
  components: {
    ProductView
  }
};
</script>
<style>
.products-view-wrapper {
    display: inline-flex;
}
</style>

