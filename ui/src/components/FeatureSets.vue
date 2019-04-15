<template>
  <v-container pa-0 ma-0>
    <v-list two-line>
      <v-subheader>
        Feature Sets&nbsp;&nbsp;&nbsp;&nbsp;
        <v-btn @click="addNewFeature()" color="primary" fab small>
          <v-icon>add</v-icon>
        </v-btn>
      </v-subheader>
      <v-progress-linear :active="loading>0" :indeterminate="true"></v-progress-linear>
      <template v-for="(fs, index) in featureSets">
        <v-divider v-if="index>0" :key="'divider-'+fs.id"></v-divider>
        <v-list-tile :key="'list-tile-'+fs.id">
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
              Client Credentials&nbsp;
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
      </template>
    </v-list>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      featureSets: [],
      loading: false
    };
  },
  mounted() {
    this.loading++;
    this.axios.get("/featuresets/list").then(this.featuresReceived);
  },
  methods: {
    featuresReceived(response) {
      this.loading--;
      this.featureSets = response.data.featureSets || [];
    },
    addNewFeature() {
      this.loading++;
      this.axios
        .post("/featuresets/add", {
          name: "test123",
          features: ["a", "b"]
        })
        .then(this.featureAdded);
    },
    featureAdded(response) {
      this.loading--;
      this.featureSets.push(response.data);
    },
    // eslint-disable-next-line
    editFeatureSet(name) {},
    deleteFeatureSet(id) {
      this.loading++;
      this.axios
        .post("/featuresets/delete", {
          id: id
        })
        .then(this.featureRemoved);
      this.featureSets = this.featureSets.filter(fs => fs.id != id);
    },
    featureRemoved() {
      this.loading--;
    },
    showEngineCredentials(name) {
      alert(name);
    }
  }
};
</script>

<style>
</style>
