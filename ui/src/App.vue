<template>
  <v-app>
    <v-navigation-drawer v-model="drawer" fixed app>
      <v-img :src="require('./assets/whales.jpg')"></v-img>
      <v-list>
        <v-subheader>MetaHub</v-subheader>
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
    <v-toolbar app fixed>
      <v-toolbar-side-icon large @click.stop="drawer = !drawer"></v-toolbar-side-icon>
      <v-toolbar-title class="headline">
        <span>Meta</span>
        <span class="font-weight-light">Hub</span>
      </v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn v-if="!isAuthenticated" flat @click="loginClicked()">
        <span>Sign In</span>
      </v-btn>
      <v-btn v-if="isAuthenticated" flat @click="logoutClicked()">
        <span>Sign Out</span>
      </v-btn>
    </v-toolbar>
    <v-content>
      <v-container align-center>
        <router-view></router-view>
      </v-container>
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
      isAuthenticated: this.$auth.isAuthenticated()
    };
  },
  methods: {
    loginClicked() {
      this.login();
    },
    loginDialogClosed() {
      this.isAuthenticated = this.$auth.isAuthenticated();
    },
    logoutClicked() {
      this.logout();
    }
  }
};
</script>
