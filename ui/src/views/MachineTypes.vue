<template>
  <v-container pl-0 pr-0>
    <v-toolbar flat color="transparent" dense>
      <v-btn :to="{ name: 'new-machine-type'}" color="primary">
        <v-icon left>add</v-icon>
        Machine Type
      </v-btn>
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
              <v-btn @click="showEngineCredentials(mt)" flat small color="primary">
                <span class="hidden-md-and-down">Client&nbsp;</span>Credentials&nbsp;
                <v-icon>account_circle</v-icon>
              </v-btn>
            </v-list-tile-action>
            <v-list-tile-action>
              <v-btn icon @click="deleteMachineType(mt.id)">
                <v-icon color="red">delete</v-icon>
              </v-btn>
            </v-list-tile-action>
          </v-list-tile>
        </template>
      </v-list>
    </v-container>
    <v-dialog :value="credentialsDialog" persistent width="650">
      <v-card>
        <v-card-title primary-title class="headline">Client Credentials</v-card-title>
        <v-card-text>
          <v-container grid-list-md>
            <v-flex xs12>
              <v-text-field label="Login" v-model="selection.login" readonly></v-text-field>
            </v-flex>
            <v-flex xs12>
              <v-text-field
                label="Password"
                v-model="selection.password"
                readonly
                :append-icon="showCredentialsPassword ? 'visibility' : 'visibility_off'"
                @click:append="showCredentialsPassword = !showCredentialsPassword"
                :type="showCredentialsPassword ? 'text' : 'password'"
              ></v-text-field>
            </v-flex>
          </v-container>
        </v-card-text>
        <v-divider></v-divider>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="info" flat @click="hideEngineCredentials()">Close</v-btn>
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
      showCredentialsPassword: false,
      newDialog: false,
      newDialogValid: false,
      selection: {},
      newFeatureName: "",
      commonFeatures: this.getCommonFeatures(),
      rules: {
        required: value => !!value || "Required."
      }
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
      this.machineTypes = this.machineTypes.filter(mt => mt.id != id);
    },
    machineTypeRemoved() {
      this.loading--;
    },
    showEngineCredentials(machineType) {
      this.showCredentialsPassword = false;
      this.selection = machineType;
      this.credentialsDialog = true;
    },
    hideEngineCredentials() {
      this.credentialsDialog = false;
    }
  }
};
</script>

<style>
</style>
