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
      <v-btn v-if="!isAuthenticated" flat @click="login()">
        <span>Sign In</span>
      </v-btn>
      <v-btn v-if="isAuthenticated" flat @click="logout()">
        <span>Sign Out</span>
      </v-btn>
    </v-toolbar>
    <v-content>
      <v-container align-center>
        <router-view></router-view>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
export default {
  name: "MetaHub",
  data() {
    return {
      drawer: undefined,
    };
  },
  computed: {
    isAuthenticated(){
      return this.$auth.isAuthenticated()
    }
  },
  methods: {
    login() {
      this.$router.push("/login");
    },
    logout() {
      this.$auth.logout();
      this.$router.push("/");
    }
  }
};
</script>
