<template>
  <div data-app>
    <v-dialog :value="show" persistent width="500">
      <v-card>
        <v-card-title class="headline grey lighten-2" primary-title>Sign In</v-card-title>
        <v-card-text>Select an authentication provider to sign in:</v-card-text>
        <v-card-text>
          <v-btn @click="authenticate('github')">
            <span>Sign in with GitHub</span>
          </v-btn>
          <v-btn @click="authenticate('google')">
            <span>Sign in with Google</span>
          </v-btn>
        </v-card-text>
        <v-card-text v-if="hasError">{{ error }}</v-card-text>
        <v-divider></v-divider>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="info" flat @click="cancel()">Cancel</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
export default {
  data() {
    return {
      show: false,
      error: undefined,
      hasError: false,
      callbacks: []
    };
  },
  methods: {
    isLoggedIn() {
      return this.$auth.isAuthenticated() && this.getInfo();
    },
    getInfo() {
      return JSON.parse(localStorage.getItem("$account"));
    },
    login(cb) {
      if (cb == undefined) cb = function() {};
      if (this.isLoggedIn()) {
        this.$nextTick(function() {
          cb(this.getInfo());
        });
        return;
      }
      this.callbacks.push(cb);
      this.show = true;
    },
    invokeCallbacks(arg) {
      for (var i in this.callbacks) {
        this.callbacks[i](arg);
      }
      this.callbacks = [];
    },
    logout() {
      if (!this.isLoggedIn()) return;
      localStorage.removeItem("$account");
      this.$auth.logout();
      this.$emit("change", undefined);
    },
    authenticate(provider) {
      this.hasError = false;
      this.$auth
        .authenticate(provider)
        .then(this.success)
        .catch(this.failure);
    },
    success() {
      this.identify();
    },
    failure(error) {
      this.hasError = true;
      this.error = error;
    },
    cancel() {
      this.hasError = false;
      this.show = false;
      this.invokeCallbacks(undefined);
    },
    identify() {
      this.axios
        .get("/auth/identity")
        .then(this.identified)
        .catch(this.identifyError);
    },
    identifyError(error) {
      this.hasError = true;
      this.error = error;
    },
    identified(response) {
      this.hasError = false;
      this.show = false;
      const account=response.data;
      localStorage.setItem("$account", JSON.stringify(account));
      this.invokeCallbacks(account);
      this.$emit("change", account);
    }
  }
};
</script>

<style>
</style>
