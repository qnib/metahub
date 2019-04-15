<template>
  <v-container>
    <v-list two-line subheader>
      <v-subheader>Feature Sets</v-subheader>
      <v-list-tile v-for="fs in featureSets" :key="fs.name">
        <v-list-tile-content>
          <v-list-tile-title>{{ fs.name }}</v-list-tile-title>
          <v-list-tile-sub-title>
            <span v-for="(feature, index) in fs.features" :key="index">
              <span v-if="index>0">, </span>
              <span>{{feature}}</span>
            </span>
          </v-list-tile-sub-title>
        </v-list-tile-content>
        <v-list-tile-action>
          <v-btn icon @click="editFeatureSet(fs.name)">
            <v-icon color="blue">edit</v-icon>
          </v-btn>
          <v-btn icon @click="deleteFeatureSet(fs.name)">
            <v-icon color="red">delete</v-icon>
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
    },
    // eslint-disable-next-line
    editFeatureSet(name) {},
    // eslint-disable-next-line
    deleteFeatureSet(name) {}
  }
};
</script>

<style>
</style>
