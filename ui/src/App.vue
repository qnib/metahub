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
      <v-toolbar-items>
        <v-btn v-if="!isAuthenticated" flat @click="loginClicked()">
          <span>Sign In</span>
        </v-btn>
        <v-btn v-if="isAuthenticated" flat @click="logoutClicked()">
          <span>Sign Out</span>
        </v-btn>
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
      isAuthenticated: this.isLoggedIn()
    };
  },
  methods: {
    loginClicked() {
      this.login();
    },
    loginDialogClosed() {
      this.isAuthenticated = this.isLoggedIn();
    },
    logoutClicked() {
      this.logout();
      this.isAuthenticated = this.isLoggedIn();
      if (this.$route.meta.requiresAuth) {
        this.$router.push({ name: "welcome" });
      }
    }
  }
};
</script>
