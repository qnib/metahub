<template>
  <v-container pa-0 ma-0>
    <v-container
      ma-0
      style="border-bottom-width: thin;
    border-bottom-color: #E0E0E0;
    border-bottom-style: solid;"
    >
      <v-btn @click="addNewFeature()" color="primary">
        Add
      </v-btn>
      <v-progress-circular v-if="loading>0" :indeterminate="true" style="float: right;"></v-progress-circular>
    </v-container>
    <v-container pt-0 pb-0 ma-0>
      <v-list two-line>
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
