<template>
  <v-container pl-0 pr-0>
    <v-toolbar flat color="transparent" dense>
      <v-btn icon :to="{ name: 'machine-types'}">
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
              v-model="machineType.name"
              :rules="[rules.required]"
              :disabled="formDisabled"
              :readonly="formReadOnly"
              :class="formReadOnly ? 'nounderline' : undefined"
              @input="updateLogin"
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
              :disabled="formDisabled"
              :readonly="formReadOnly"
              :class="formReadOnly ? 'nounderline' : undefined"
            >
              <template v-slot:selection="data">
                <v-chip
                  :key="JSON.stringify(data.item)"
                  :selected="data.selected"
                  :close="!formReadOnly"
                  class="chip--select-multi"
                  @input="removeFeature(data.item)"
                  :disabled="formDisabled"
                >{{ formatFeature(data.item) }}</v-chip>
              </template>
              <template v-slot:item="data">
                <v-list-tile-content v-text="data.item.title"></v-list-tile-content>
              </template>
            </v-combobox>
          </v-flex>
          <v-flex xs12>
            <v-card>
              <v-card-title>Client Login Credentials</v-card-title>
              <v-card-text>
                <v-text-field
                  v-if="formReadOnly"
                  label="Username"
                  :value="loginPrefix+machineType.login"
                  readonly
                  class="nounderline"
                ></v-text-field>
                <v-text-field
                  v-else
                  label="Username"
                  :prefix="loginPrefix"
                  placeholder="name"
                  v-model="machineType.login"
                  :rules="[rules.required]"
                  :disabled="formDisabled"
                  :class="formReadOnly ? 'nounderline' : undefined"
                  @change="applyNameChangedToLogin=false"
                ></v-text-field>
                <v-text-field
                  label="Password"
                  v-model="machineType.password"
                  readonly
                  class="nounderline"
                  :append-icon="showCredentialsPassword ? 'visibility' : 'visibility_off'"
                  @click:append="showCredentialsPassword = !showCredentialsPassword"
                  :type="showCredentialsPassword ? 'text' : 'password'"
                ></v-text-field>
              </v-card-text>
            </v-card>
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
      machineTypeBeforeEdit: undefined,
      machineType: {
        name: "",
        features: [],
        login: "",
        password: ""
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
    updateLogin() {
      if (this.machineType.login == "") {
        this.applyNameChangedToLogin = true;
      }
      if (!this.applyNameChangedToLogin) return;
      this.machineType.login = this.machineType.name
        .replace(/\W/g, "")
        .toLowerCase();
    },
    new() {
      this.applyNameChangedToLogin = true;
      const displayName = this.$account.getInfo().displayName;
      this.loginPrefix = displayName.replace(/\W/g, "").toLowerCase() + "-";
      var Password = {
        _pattern: /[a-zA-Z0-9_\-+.]/,
        _getRandomByte: function() {
          var result;
          // http://caniuse.com/#feat=getrandomvalues
          if (window.crypto && window.crypto.getRandomValues) {
            result = new Uint8Array(1);
            window.crypto.getRandomValues(result);
            return result[0];
          } else if (window.msCrypto && window.msCrypto.getRandomValues) {
            result = new Uint8Array(1);
            window.msCrypto.getRandomValues(result);
            return result[0];
          } else {
            return Math.floor(Math.random() * 256);
          }
        },
        generate: function(length) {
          return Array.apply(null, { length: length })
            .map(function() {
              var result;
              // eslint-disable-next-line
              while (true) {
                result = String.fromCharCode(this._getRandomByte());
                if (this._pattern.test(result)) {
                  return result;
                }
              }
            }, this)
            .join("");
        }
      };
      this.machineType.password = Password.generate(24);
      this.$refs.nameTextField.focus();
    },
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
      this.extractAccountNameFromLogin();
    },
    extractAccountNameFromLogin() {
      var myRegexp = /^(\w+-)(.+)$/g;
      var match = myRegexp.exec(this.machineType.login);
      if (match == null) {
        this.inProgress = false;
        alert("Invalid format for login '" + this.machineType.login + "'");
        this.$router.push({ name: "machine-types" });
        return;
      }
      this.loginPrefix = match[1];
      this.machineType.login = match[2];
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
    edit() {
      this.editing = true;
      this.machineTypeBeforeEdit = JSON.parse(JSON.stringify(this.machineType));
      this.$refs.nameTextField.focus();
    },
    cancel() {
      if (this.editing) {
        this.editing = false;
        this.machineType = JSON.parse(
          JSON.stringify(this.machineTypeBeforeEdit)
        );
      } else {
        this.$router.push({ name: "machine-types" });
      }
    },
    save() {
      this.inProgress = true;

      var toSend = JSON.parse(JSON.stringify(this.machineType));
      toSend.login = this.loginPrefix + toSend.login;

      var req;
      if (this.machineType.id) {
        req = this.axios.post("/machinetypes/update", toSend);
      } else {
        req = this.axios.post("/machinetypes/add", toSend);
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
        this.$router.push({ name: "machine-types" });
      } else {
        this.editing = false;
        this.machineType = response.data;
        this.extractAccountNameFromLogin();
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
