<template>
  <div id="app">
    <keep-alive>
      <router-view></router-view>
    </keep-alive>

  </div>
</template>

<script>
export default {
  name: 'App',
  methods: {
    // click() {
    //   this.$router.push("/login")
    // }
  }
}
</script>


<style>

#app{
  display: flex;
  height: 100%;
  width: 100%;
}
</style>
