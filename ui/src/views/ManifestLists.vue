<template>
  <v-container pl-0 pr-0>
    <v-toolbar flat color="transparent" dense>
      <v-btn :to="{ name: 'new-manifest-list'}" color="primary">
        <v-icon left>add</v-icon>Manifest List
      </v-btn>
      <v-spacer></v-spacer>
      <v-progress-circular v-if="loading" :indeterminate="true" style="float: right;"></v-progress-circular>
    </v-toolbar>
    <v-container
      pa-2
      mt-3
      style="border-top-width: thin;
    border-top-color: #E0E0E0;
    border-top-style: solid;"
    >
      <v-list two-line>
        <template v-for="(ml, index) in manifestList">
          <v-divider v-if="index>0" :key="'divider-'+ml.id"></v-divider>
          <v-list-tile :key="'list-tile-'+ml.id">
            <v-list-tile-content>
              <v-list-tile-title>
                <router-link :to="{ name: 'edit-manifest-list', params: { id: ml.id }}">{{ ml.name }}</router-link>
              </v-list-tile-title>
            </v-list-tile-content>
            <v-list-tile-action>
              <v-btn icon @click="deleteManifestList(ml.id)">
                <v-icon color="red">delete</v-icon>
              </v-btn>
            </v-list-tile-action>
          </v-list-tile>
        </template>
      </v-list>
    </v-container>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      manifestList: [],
      loading: false
    };
  },
  mounted() {
    this.loading = true;
    this.axios
      .get("/manifestlists/list")
      .then(this.featuresReceived)
      .catch(this.listError);
  },
  methods: {
    listError(error) {
      this.loading = false;
      alert(error);
    },
    deleteManifestList(id) {
      this.loading = true;
      this.axios
        .post("/manifestlists/delete", {
          id: id
        })
        .then(this.manifestListRemoved);
      this.manifestList = this.manifestList.filter(ml => mt.id != id);
    },
    manifestListRemoved() {
      this.loading = false;
    }
  }
};
</script>

<style>
</style>
