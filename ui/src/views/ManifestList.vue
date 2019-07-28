<template>
  <v-container pl-0 pr-0>
    <v-toolbar flat color="transparent" dense>
      <v-btn icon :to="{ name: 'manifest-lists'}">
        <v-icon>arrow_back</v-icon>
      </v-btn>
      <v-btn color="primary" v-if="!isNew" :disabled="editing || inProgress" @click="edit()">Edit</v-btn>
      <v-btn
        v-if="editing || isNew"
        color="primary"
        :disabled="!formValid || inProgress"
        @click="save()"
      >Save</v-btn>
      <v-btn v-if="editing || isNew" :disabled="inProgress" @click="cancel()">Cancel</v-btn>
      <v-spacer></v-spacer>
      <v-progress-circular v-if="inProgress" :indeterminate="true" style="float: right;"></v-progress-circular>
    </v-toolbar>
    <v-container
      mt-3
      style="border-top-width: thin;
    border-top-color: #E0E0E0;
    border-top-style: solid;"
    >
      <v-form v-model="formValid" v-if="showForm">
        <v-layout row wrap>
          <v-flex xs12>
            <v-text-field
              ref="nameTextField"
              label="Name"
              v-model="manifestList.name"
              :rules="[rules.required]"
              :disabled="formDisabled"
              :readonly="formReadOnly"
              :class="formReadOnly ? 'nounderline' : undefined"
              @input="updateLogin"
            ></v-text-field>
          </v-flex>
        </v-layout>
      </v-form>
    </v-container>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      editing: false,
      isNew: false,
      loading: false,
      inProgress: false,
      formValid: false,
      newFeatureName: "",
      commonFeatures: this.getCommonFeatures(),
      rules: {
        required: value => !!value || "Required."
      },
      manifestListBeforeEdit: undefined,
      manifestList: {
        name: "",
        reponame: "",
        tagname: "",
      },
      showCredentialsPassword: false,
      loginPrefix: "",
      applyNameChangedToLogin: false
    };
  },
  computed: {
    showForm() {
      return !this.loading;
    },
    formDisabled() {
      return this.inProgress;
    },
    formReadOnly() {
      return !this.isNew && !this.editing;
    },
    showStats() {
      return !this.loading && !this.isNew && !this.editing;
    }
  },
  mounted() {
    this.isNew = this.$route.params.id ? false : true;
    if (this.$route.params.id) {
      this.load(this.$route.params.id);
    } else {
      this.new();
    }
  },
  methods: {
    new() {
      this.applyNameChangedToLogin = true;
      const displayName = this.$account.getInfo().displayName;
      this.$refs.nameTextField.focus();
    },
    load(id) {
      this.manifestList.id = id;
      this.loading = this.inProgress = true;
      this.axios
        .get("/manifestlists/get", {
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
      this.$router.push({ name: "manifest-lists" });
    },
    loaded(response) {
      this.loading = this.inProgress = false;
      this.manifestList = response.data;
      this.extractAccountNameFromLogin();
    },
    edit() {
      this.editing = true;
      this.manifestListBeforeEdit = JSON.parse(JSON.stringify(this.manifestList));
      this.$refs.nameTextField.focus();
    },
    cancel() {
      if (this.editing) {
        this.editing = false;
        this.manifestList = JSON.parse(
          JSON.stringify(this.manifestListBeforeEdit)
        );
      } else {
        this.$router.push({ name: "manifest-lists" });
      }
    },
    save() {
      this.inProgress = true;

      var toSend = JSON.parse(JSON.stringify(this.manifestList));
      toSend.login = this.loginPrefix + toSend.login;

      var req;
      if (this.manifestList.id) {
        req = this.axios.post("/manifestlists/update", toSend);
      } else {
        req = this.axios.post("/manifestlists/add", toSend);
      }
      req.then(this.saved).catch(this.saveError);
    },
    saveError(error) {
      this.inProgress = false;
      alert(error);
    },
    saved(response) {
      this.inProgress = false;
      if (this.isNew) {
        this.$router.push({ name: "manifest-lists" });
      } else {
        this.editing = false;
        this.manifestList = response.data;
      }
    }
  }
};
</script>

<style>
h1 {
  font-size: 1em;
}
.v-card__title {
  font-size: larger;
  padding-bottom: 0;
}
.v-card__text {
  padding-bottom: 1px;
}

.nounderline.v-text-field > .v-input__control > .v-input__slot:before {
  border-style: none;
}
.nounderline.v-text-field > .v-input__control > .v-input__slot:after {
  border-style: none;
}
</style>
