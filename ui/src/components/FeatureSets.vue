<template>
  <v-container>
    <v-list two-line subheader>
      <v-subheader>Feature Sets</v-subheader>
      <v-list-tile v-for="fs in featureSets" :key="fs.name">
        <v-list-tile-content>
          <v-list-tile-title>{{ fs.name }}</v-list-tile-title>
          <v-list-tile-sub-title>{{ fs.features }}</v-list-tile-sub-title>
        </v-list-tile-content>
        <v-list-tile-action>
          <v-btn icon ripple>
            <v-icon color="grey lighten-1">info</v-icon>
          </v-btn>
        </v-list-tile-action>
      </v-list-tile>
    </v-list>
    <v-btn @click="addNewFeature()">Add</v-btn>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      featureSets: []
    };
  },
  mounted() {
    this.axios.get("/featuresets/list").then(this.featuresReceived);
  },
  methods: {
    featuresReceived(response) {
      this.featureSets = response.data.featureSets || [];
    },
    addNewFeature() {
      this.axios
        .post("/featuresets/add", {
          name: "test123"
        })
        .then(this.featureAdded);
    },
    featureAdded(response) {
      window.console.log(response);
      this.featureSets.push(response.data);
    }
  }
};
</script>

<style>
</style>
