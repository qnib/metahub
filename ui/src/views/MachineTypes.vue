<template>
  <v-container pl-0 pr-0>
    <v-toolbar flat color="transparent" dense>
      <v-btn :to="{ name: 'new-machine-type'}" color="primary">
        <v-icon left>add</v-icon>Machine Type
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
        <template v-for="(mt, index) in machineTypes">
          <v-divider v-if="index>0" :key="'divider-'+mt.id"></v-divider>
          <v-list-tile :key="'list-tile-'+mt.id">
            <v-list-tile-content>
              <v-list-tile-title>
                <router-link :to="{ name: 'edit-machine-type', params: { id: mt.id }}">{{ mt.name }}</router-link>
              </v-list-tile-title>
              <v-list-tile-sub-title v-if="mt.features">
                <span v-for="(feature, index) in mt.features" :key="index">
                  <span v-if="index>0">,&nbsp;</span>
                  <span>{{formatFeature(feature)}}</span>
                </span>
              </v-list-tile-sub-title>
              <v-list-tile-sub-title v-else>(no features specified)</v-list-tile-sub-title>
            </v-list-tile-content>
            <v-list-tile-action>
              <v-btn icon @click="deleteMachineType(mt.id)">
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
      machineTypes: [],
      loading: false
    };
  },
  mounted() {
    this.loading = true;
    this.axios
      .get("/machinetypes/list")
      .then(this.featuresReceived)
      .catch(this.listError);
  },
  methods: {
    listError(error) {
      this.loading = false;
      alert(error);
    },
    featuresReceived(response) {
      this.loading = false;
      this.machineTypes = response.data.machineTypes || [];
    },
    deleteMachineType(id) {
      this.loading = true;
      this.axios
        .post("/machinetypes/delete", {
          id: id
        })
        .then(this.machineTypeRemoved);
      this.machineTypes = this.machineTypes.filter(mt => mt.id != id);
    },
    machineTypeRemoved() {
      this.loading = false;
    }
  }
};
</script>

<style>
</style>
