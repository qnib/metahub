<template>
  <v-container pl-0 pr-0>
    <v-toolbar flat color="transparent" dense>
      <v-btn @click="addNewFeature()" color="primary">Add</v-btn>
      <v-spacer></v-spacer>
      <v-progress-circular v-if="loading>0" :indeterminate="true" style="float: right;"></v-progress-circular>
    </v-toolbar>
    <v-container
      pa-2
      mt-3
      style="border-top-width: thin;
    border-top-color: #E0E0E0;
    border-top-style: solid;"
    >
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
              <v-btn @click="showEngineCredentials(fs)" flat small color="primary">
                Client Credentials&nbsp;
                <v-icon>account_circle</v-icon>
              </v-btn>
            </v-list-tile-content>
            <v-list-tile-action>
              <v-btn icon @click="showEditDialog(fs)">
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
    <v-dialog :value="showCredentials" persistent width="500">
      <v-card>
        <v-card-title primary-title>Registry Client Credentials</v-card-title>
        <v-card-text>.............</v-card-text>
        <v-divider></v-divider>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="info" flat @click="hideEngineCredentials()">Close</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-dialog :value="editFeatureSet" persistent width="500">
      <v-card>
        <v-card-title primary-title>Edit Feature Set</v-card-title>
        <v-card-text>.............</v-card-text>
        <v-divider></v-divider>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="info" flat @click="closeEditDialog()">Close</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      featureSets: [],
      loading: false,
      showCredentials: false,
      editFeatureSet: false
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
    showEngineCredentials(featureSet) {
      this.showCredentialsForFeatureSet = featureSet;
      this.showCredentials = true;
    },
    hideEngineCredentials() {
      this.showCredentials = false;
      this.showCredentialsForFeatureSet = undefined;
    },
    // eslint-disable-next-line
    showEditDialog(featureSet) {
      this.editFeatureSet = true;
    },
    closeEditDialog() {
      this.editFeatureSet = false;
    }
    //editFeatureSetDialog
  }
};
</script>

<style>
</style>
