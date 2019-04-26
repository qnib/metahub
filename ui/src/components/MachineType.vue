<template>
  <v-container pl-0 pr-0>
    <v-toolbar flat color="transparent" dense>
      <v-btn color="primary" :disabled="!formValid || loading" @click="save()">Save</v-btn>
      <v-btn :to="{ name: 'machine-types'}">Cancel</v-btn>
      <v-spacer></v-spacer>
      <v-progress-circular v-if="loading" :indeterminate="true" style="float: right;"></v-progress-circular>
    </v-toolbar>
    <v-container
      mt-3
      style="border-top-width: thin;
    border-top-color: #E0E0E0;
    border-top-style: solid;"
    >
      <v-form v-model="formValid" v-if="machineType">
        <v-flex xs12>
          <v-text-field
            label="Name"
            v-model="machineType.name"
            :rules="[rules.required]"
            :disabled="loading"
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
            :disabled="loading"
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
      loading: false,
      formValid: false,
      newFeatureName: "",
      commonFeatures: this.getCommonFeatures(),
      rules: {
        required: value => !!value || "Required."
      },
      machineType: undefined
    };
  },
  mounted() {
    this.loading = true;
    this.axios
      .get("/machinetypes/get", {
        params: { 
          id: this.$route.params.id
        }
      })
      .then(this.loaded)
      .catch(this.loadError);
  },
  methods: {
    loadError(error) {
      this.loading = false;
      alert(error);
      this.$router.push({ name: "machine-types" });
    },
    loaded(response) {
      this.loading = false;
      this.machineType = response.data;
    },
    addFeature() {
      this.machineType.features.push(this.newFeatureName);
      this.newFeatureName = "";
    },
    save() {
      this.loading = true;
      this.axios
        .post("/machinetypes/update", this.machineType)
        .then(this.saved)
        .catch(this.saveError);
    },
    saveError(error) {
      this.loading = false;
      alert(error);
    }, 
    saved() {
      this.loading = false;
      this.$router.push({ name: "machine-types" });
    }
  }
};
</script>

<style>
</style>
