<template>
  <v-dialog :value="dialog" width="500">
    <v-card>
      <v-card-title class="headline grey lighten-2" primary-title>Sign In to MetaHub</v-card-title>
      <v-card-text>
        <v-btn @click="authenticate('github')">
          <span>Sign in with GitHub</span>
        </v-btn>
      </v-card-text>
      <v-card-text v-if="hasError">{{ error }}</v-card-text>
      <v-divider></v-divider>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="info" flat @click="hide()">Cancel</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  data() {
    return {
      dialog: false,
      hasError: false
    };
  },
  methods: {
    show() {
      this.hasError = false;
      this.dialog = true;
    },
    hide() {
      this.dialog = false;
      this.$emit("close", true);
    },
    authenticate(provider) {
      this.hasError = false;
      this.$auth
        .authenticate(provider)
        .then(this.success)
        .catch(this.failure);
    },
    success() {
      window.console.log("logged in");
      this.hide();
    },
    failure(error) {
      //window.console.log(error);
      this.hasError = true;
      this.error = error;
    }
  }
};
</script>

<style>
</style>
