<template>
  <v-app style="background-color: white;">
    <v-navigation-drawer v-model="drawer" app>
      <v-toolbar>
        <v-toolbar-title class="headline">
          <span>Meta</span>
          <span class="font-weight-light">Hub</span>
        </v-toolbar-title>
      </v-toolbar>
      <v-list>
        <v-list-tile avatar ripple to="/">
          <v-list-tile-avatar>
            <v-icon>dashboard</v-icon>
          </v-list-tile-avatar>
          <v-list-tile-content>
            <v-list-tile-title>Welcome</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
        <v-divider inset></v-divider>
        <v-list-tile avatar ripple to="/machinetypes">
          <v-list-tile-avatar>
            <v-icon>dashboard</v-icon>
          </v-list-tile-avatar>
          <v-list-tile-content>
            <v-list-tile-title>Machine Types</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
        <v-divider inset></v-divider>
      </v-list>
    </v-navigation-drawer>
    <v-toolbar app>
      <v-toolbar-side-icon large @click.stop="drawer = !drawer"></v-toolbar-side-icon>
      <v-toolbar-title class="headline">{{$route.meta.title}}</v-toolbar-title>
      <router-view name="tools"></router-view>
      <v-spacer></v-spacer>
      <v-toolbar-items v-if="!noAuth">
        <v-btn v-if="!account" flat @click="loginClicked()">
          <span>Sign In</span>
        </v-btn>
        <v-menu v-if="account" offset-y>
          <template v-slot:activator="{ on }">
            <v-btn flat v-on="on">{{ account.displayName }}</v-btn>
          </template>
          <v-list>
            <v-list-tile @click="logoutClicked()">
              <v-list-tile-title>Sign Out</v-list-tile-title>
            </v-list-tile>
          </v-list>
        </v-menu>
      </v-toolbar-items>
    </v-toolbar>
    <v-content>
      <router-view></router-view>
    </v-content>
    <LoginDialog @close="loginDialogClosed()" ref="login-dialog"/>
  </v-app>
</template>

<script>
import LoginDialog from "./components/LoginDialog";
export default {
  name: "MetaHub",
  components: {
    LoginDialog
  },
  data() {
    return {
      drawer: undefined,
      account: this.$account.getInfo(),
      noAuth: this.$account == undefined
    };
  },
  mounted() {
    if (this.$account) {
      this.$account.$on("change", this.accountChanged);
    }
  },
  beforeDestroy() {
    if (this.$account) {
      this.$account.$off("change");
    }
  },
  methods: {
    loginClicked() {
      this.$account.login();
    },
    accountChanged(account) {
      this.account = account;
      if (!account && this.$route.meta.requiresAuth) {
        this.$router.push({ name: "welcome" });
      }
    },
    logoutClicked() {
      this.$account.logout();
    }
  }
};
</script>
