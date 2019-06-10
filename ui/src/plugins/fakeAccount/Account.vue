<template>
  <div></div>
</template>

<script>
export default {
  methods: {
    isLoggedIn() {
      return true;
    },
    getInfo() {
      return {
        displayName: "dummy"
      };
    },
    login(cb) {
      if (cb == undefined) cb = function() {};
      this.$nextTick(function() {
        cb(this.getInfo());
      });
    },
    logout() {}
  }
};
</script>

<style>
</style>
