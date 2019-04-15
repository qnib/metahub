<template>
  <v-container>
    <v-list two-line subheader>
      <v-subheader>Feature Sets</v-subheader>
      <v-list-tile v-for="fs in featureSets" :key="fs.id">
        <v-list-tile-content>
          <v-list-tile-title>{{ fs.name }}</v-list-tile-title>
          <v-list-tile-sub-title>
            <span v-for="(feature, index) in fs.features" :key="index">
              <span v-if="index>0">,&nbsp;</span>
              <span>{{feature}}</span>
            </span>
          </v-list-tile-sub-title>
        </v-list-tile-content>
        <v-list-tile-content>
          <v-btn @click="showEngineCredentials(fs.id)" flat small color="primary">
            Engine Credentials&nbsp;
            <v-icon>account_circle</v-icon>
          </v-btn>
        </v-list-tile-content>
        <v-list-tile-action>
          <v-btn icon @click="editFeatureSet(fs.id)">
            <v-icon color="blue">edit</v-icon>
          </v-btn>
        </v-list-tile-action>
        <v-list-tile-action>
          <v-btn icon @click="deleteFeatureSet(fs.id)">
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
          name: "test123",
          features: ["a", "b"]
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
    deleteFeatureSet(id) {
      this.axios.post("/featuresets/delete", {
        id: id
      });
      this.featureSets = this.featureSets.filter(fs => fs.id != id);
    },
    showEngineCredentials(name) {
      alert(name);
    }
  }
};
</script>

<style>
</style>
