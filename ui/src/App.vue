<template>
  <v-app>
    <v-navigation-drawer v-model="drawer" fixed app>
      <v-img :src="require('./assets/whales.jpg')"></v-img>
      <v-list>
        <v-subheader>MetaHub</v-subheader>
        <v-list-tile avatar ripple>
          <v-list-tile-avatar>
            <v-icon>dashboard</v-icon>
          </v-list-tile-avatar>
          <v-list-tile-content>
            <v-list-tile-title>Runtime Environments</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
        <v-divider inset></v-divider>
        <v-list-tile avatar ripple>
          <v-list-tile-avatar>
            <v-icon>dashboard</v-icon>
          </v-list-tile-avatar>
          <v-list-tile-content>
            <v-list-tile-title>Backend Registries</v-list-tile-title>
          </v-list-tile-content>
        </v-list-tile>
        <v-divider inset></v-divider>
        <v-list-tile avatar ripple>
          <v-list-tile-avatar>
            <v-icon>dashboard</v-icon>
          </v-list-tile-avatar>
          <v-list-tile-content>
            <v-list-tile-title>Rate Limiting</v-list-tile-title>
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
      <v-btn flat @click="authenticate('github')">
        <span>Sign In</span>
      </v-btn>
    </v-toolbar>
    <v-content>
      <v-container align-center>
        <Welcome/>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
import Welcome from "./components/Welcome";

export default {
  name: "App",
  components: {
    Welcome
  },
  data() {
    return {
      drawer: undefined
    };
  },
  methods: {
    authenticate(provider) {
      this.$auth.authenticate(provider).then(function() {
        // Execute application logic after successful social authentication
        window.console.log("yeah");
      });
    }
  }
};
</script>
