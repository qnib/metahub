<template>
  <v-container>
    Machine Types
    <v-btn @click="add()">add</v-btn>
  </v-container>
</template>

<script>
export default {
  data: () => ({}),
  methods: {
    add() {
      this.axios.get("/v2/").then(response => {
        window.console.log(response);
      });
    }
  }
};
</script>

<style>
</style>
