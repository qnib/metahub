
window.app = new Vue({
  el: '#app',
  template: `
  <v-app>

    <v-navigation-drawer v-model="drawer" fixed app>
    <v-img src="./icons/icon-512x512.png"></v-img>
    <v-list>
        <v-subheader>MetaHub</v-subheader>
        <v-list-tile avatar ripple>
          <v-list-tile-avatar><v-icon>dashboard</v-icon></v-list-tile-avatar>
          <v-list-tile-content><v-list-tile-title>Runtime Environments</v-list-tile-title></v-list-tile-content>
        </v-list-tile>
        <v-divider inset></v-divider>
        <v-list-tile avatar ripple>
          <v-list-tile-avatar><v-icon>dashboard</v-icon></v-list-tile-avatar>
          <v-list-tile-content><v-list-tile-title>Backend Registries</v-list-tile-title></v-list-tile-content>
        </v-list-tile>
        <v-divider inset></v-divider>
        <v-list-tile avatar ripple>
          <v-list-tile-avatar><v-icon>dashboard</v-icon></v-list-tile-avatar>
          <v-list-tile-content><v-list-tile-title>Rate Limiting</v-list-tile-title></v-list-tile-content>
        </v-list-tile>
        <v-divider inset></v-divider>
      </v-list>
    </v-navigation-drawer>

    <v-toolbar app fixed>
      <v-toolbar-side-icon large @click.stop="drawer = !drawer"></v-toolbar-side-icon>
      <v-toolbar-title>MetaHub</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-progress-circular v-if="loading" indeterminate color="primary"></v-progress-circular>
    </v-toolbar>

    <v-content>
      <v-container align-center ma-0 pa-0 style="height: calc(100% - 64px);">
      </v-container>
    </v-content>
  </v-app>
  `,
  data: {
    drawer: undefined,
    loading: false,
  },
});

