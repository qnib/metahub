<template>
  <v-container pl-0 pr-0>
    <v-toolbar flat color="transparent" dense>
      <v-btn @click="showNewDialog()" color="primary">Add</v-btn>
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
        <template v-for="(fs, index) in machineTypes">
          <v-divider v-if="index>0" :key="'divider-'+fs.id"></v-divider>
          <v-list-tile :key="'list-tile-'+fs.id">
            <v-list-tile-content>
              <v-list-tile-title>{{ fs.name }}</v-list-tile-title>
              <v-list-tile-sub-title>
                <span v-for="(feature, index) in fs.features" :key="index">
                  <span v-if="index>0">,&nbsp;</span>
                  <span>{{formatFeature(feature)}}</span>
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
              <v-btn icon @click="deleteMachineType(fs.id)">
                <v-icon color="red">delete</v-icon>
              </v-btn>
            </v-list-tile-action>
          </v-list-tile>
        </template>
      </v-list>
    </v-container>
    <v-dialog :value="credentialsDialog" persistent width="500">
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
    <v-dialog :value="editDialog" persistent width="500">
      <v-card>
        <v-card-title primary-title class="headline">Edit Machine Type</v-card-title>
        <v-container grid-list-md>
          <v-flex xs12>
            <v-text-field label="Name" v-model="selection.name" required></v-text-field>
          </v-flex>
          <v-flex xs12>
            <v-combobox
              v-model="selection.features"
              :items="commonFeatures"
              chips
              label="Features"
              item-value="name"
              :return-object="false"
              multiple
              dense
            >
              <template v-slot:selection="data">
                <v-chip
                  :key="JSON.stringify(data.item)"
                  :selected="data.selected"
                  close
                  class="chip--select-multi"
                  @input="removeFeature(data.item)"
                >{{ formatFeature(data.item) }}</v-chip>
              </template>
              <template v-slot:item="data">
                <v-list-tile-content v-text="data.item.title"></v-list-tile-content>
              </template>
            </v-combobox>
          </v-flex>
        </v-container>
        <v-divider></v-divider>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="info" flat @click="closeEditDialog()">Cancel</v-btn>
          <v-btn color="primary" @click="confirmEditDialog()">Update</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-dialog :value="newDialog" persistent width="500">
      <v-card>
        <v-card-title primary-title class="headline">Add Machine Type</v-card-title>
        <v-card-text>
          <v-container grid-list-md>
            <v-flex xs12>
              <v-text-field label="Name" v-model="selection.name" required></v-text-field>
            </v-flex>
            <v-flex xs12>
              <v-combobox
                v-model="selection.features"
                :items="commonFeatures"
                chips
                label="Features"
                item-value="name"
                :return-object="false"
                multiple
                dense
              >
                <template v-slot:selection="data">
                  <v-chip
                    :key="JSON.stringify(data.item)"
                    :selected="data.selected"
                    close
                    class="chip--select-multi"
                    @input="removeFeature(data.item)"
                  >{{ formatFeature(data.item) }}</v-chip>
                </template>
                <template v-slot:item="data">
                  <v-list-tile-content v-text="data.item.title"></v-list-tile-content>
                </template>
              </v-combobox>
            </v-flex>
          </v-container>
        </v-card-text>
        <v-divider></v-divider>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="info" flat @click="cancelNewDialog()">Cancel</v-btn>
          <v-btn color="primary" @click="confirmNewDialog()">Add</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      machineTypes: [],
      loading: false,
      credentialsDialog: false,
      editDialog: false,
      newDialog: false,
      selection: {},
      newFeatureName: "",
      commonFeatures: [
        { header: "CPU" },
        { name: "cpu:skylake", title: "Skylake", group: "CPU" },
        { header: "GPU" },
        { name: "gpu:tegrak1", title: "Tegra K1", group: "GPU" }
      ]
    };
  },
  mounted() {
    this.loading++;
    this.axios.get("/machinetypes/list").then(this.featuresReceived);
  },
  methods: {
    featuresReceived(response) {
      this.loading--;
      this.machineTypes = response.data.machineTypes || [];
    },
    deleteMachineType(id) {
      this.loading++;
      this.axios
        .post("/machinetypes/delete", {
          id: id
        })
        .then(this.machineTypeRemoved);
      this.machineTypes = this.machineTypes.filter(fs => fs.id != id);
    },
    machineTypeRemoved() {
      this.loading--;
    },
    showEngineCredentials(machineType) {
      this.selection = machineType;
      this.credentialsDialog = true;
    },
    hideEngineCredentials() {
      this.credentialsDialog = false;
    },
    removeFeature(feature) {
      this.selection.features = this.selection.features.filter(function(f) {
        return f != feature;
      });
    },
    formatFeature(name) {
      for (var i in this.commonFeatures) {
        const f = this.commonFeatures[i];
        if (f.name == name) {
          return f.title;
        }
      }
      return name;
    },
    addFeature() {
      this.selection.features.push(this.newFeatureName);
      this.newFeatureName = "";
    },
    showEditDialog(machineType) {
      this.selection = JSON.parse(JSON.stringify(machineType));
      this.editDialog = true;
    },
    closeEditDialog() {
      this.editDialog = false;
    },
    confirmEditDialog() {
      this.editDialog = false;
      this.loading++;
      this.axios
        .post("/machinetypes/update", this.selection)
        .then(this.machineTypeUpdated);
      for (var i = 0; i < this.machineTypes.length; i++) {
        if (this.machineTypes[i].id == this.selection.id) {
          this.machineTypes[i] = this.selection;
        }
      }
    },
    machineTypeUpdated(response) {
      this.loading--;
    },
    showNewDialog() {
      this.selection = {
        name: "",
        features: []
      };
      this.newFeatureName = "";
      this.newDialog = true;
    },
    cancelNewDialog() {
      this.newDialog = false;
    },
    confirmNewDialog() {
      this.newDialog = false;
      this.loading++;
      this.axios
        .post("/machinetypes/add", this.selection)
        .then(this.machineTypeAdded);
    },
    machineTypeAdded(response) {
      this.loading--;
      this.machineTypes.push(response.data);
    }
  }
};
</script>

<style>
</style>
