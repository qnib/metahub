<template>
  <v-container pl-0 pr-0>
    <v-toolbar flat color="transparent" dense>
      <v-btn color="primary" :disabled="!formValid || inProgress" @click="save()">
        <span v-if="machineType.id">Update</span>
        <span v-if="!machineType.id">Add</span>
      </v-btn>
      <v-btn :to="{ name: 'machine-types'}">Cancel</v-btn>
      <v-spacer></v-spacer>
      <v-progress-circular v-if="inProgress" :indeterminate="true" style="float: right;"></v-progress-circular>
    </v-toolbar>
    <v-container
      mt-3
      style="border-top-width: thin;
    border-top-color: #E0E0E0;
    border-top-style: solid;"
    >
      <v-form v-model="formValid" v-if="machineType && !loading">
        <v-flex xs12>
          <v-text-field
            label="Name"
            v-model="machineType.name"
            :rules="[rules.required]"
            :disabled="inProgress"
          ></v-text-field>
        </v-flex>
        <v-flex xs12>
          <v-combobox
            v-model="machineType.features"
            :items="commonFeatures"
            chips
            label="Features"
            item-value="name"
            :return-object="false"
            multiple
            dense
            hide-selected
            :disabled="inProgress"
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
      </v-form>
    </v-container>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      new: false,
      loading: false,
      inProgress: false,
      formValid: false,
      newFeatureName: "",
      commonFeatures: this.getCommonFeatures(),
      rules: {
        required: value => !!value || "Required."
      },
      machineType: {
        name: "",
        features: []
      }
    };
  },
  mounted() {
    this.new = this.$route.params.id ? false : true;
    if (this.$route.params.id) {
      this.load(this.$route.params.id);
    }
  },
  methods: {
    load(id) {
      this.machineType.id = id;
      this.loading = this.inProgress = true;
      this.axios
        .get("/machinetypes/get", {
          params: {
            id: id
          }
        })
        .then(this.loaded)
        .catch(this.loadError);
    },
    loadError(error) {
      this.inProgress = false;
      alert(error);
      this.$router.push({ name: "machine-types" });
    },
    loaded(response) {
      this.loading = this.inProgress = false;
      this.machineType = response.data;
    },
    removeFeature(feature) {
      this.machineType.features = this.machineType.features.filter(function(f) {
        return f != feature;
      });
    },
    addFeature() {
      this.machineType.features.push(this.newFeatureName);
      this.newFeatureName = "";
    },
    save() {
      this.inProgress = true;
      var req;
      if (this.machineType.id) {
        req = this.axios.post("/machinetypes/update", this.machineType);
      } else {
        req = this.axios.post("/machinetypes/add", this.machineType);
      }
      req.then(this.saved).catch(this.saveError);
    },
    saveError(error) {
      this.inProgress = false;
      alert(error);
    },
    saved() {
      this.inProgress = false;
      this.$router.push({ name: "machine-types" });
    }
  }
};
</script>

<style>
</style>
