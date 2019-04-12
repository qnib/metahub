<template>
  <v-container>
    <v-btn @click="authenticate('github')">
      <span>Sign in with GitHub</span>
    </v-btn>
  </v-container>
</template>

<script>
export default {
  data() {
    return {};
  },
  methods: {
    authenticate(provider) {
      var self = this;
      this.$auth
        .authenticate(provider)
        .then(function() {
          window.console.log("logged in");
          const target = self.$route.query.redirect || "";
          self.$router.push(target)
        })
        .catch(function(error) {
          window.console.log(error);
          alert(error);
        });
    }
  }
};
</script>

<style>
</style>
