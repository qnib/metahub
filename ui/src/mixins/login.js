import Vue from 'vue'

Vue.mixin({
  methods: {
    login: function (cb) {
      if (this.isLoggedIn()) {
        Vue.nextTick(function () {
          cb(true);
        })
        return;
      }
      var self = this;
      const refs = this.$root.$children[0].$refs;
      const loginDialog = refs["login-dialog"];
      const closed = function () {
        loginDialog.$off("close", closed);
        if (cb) {
          cb(self.isLoggedIn());
        }
      };
      loginDialog.$on("close", closed);
      loginDialog.show();
    },
    logout: function () {
      this.$auth.logout();
      //this.$router.push("/");
    },
    isLoggedIn: function () {
      return this.$auth.isAuthenticated();
    }
  }
});
