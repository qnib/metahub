import Vue from 'vue'

Vue.mixin({
  methods: {
    getCommonFeatures() {
      return [
        { header: "CPU" },
        { name: "cpu:skylake", title: "Skylake", group: "CPU" },
        { name: "cpu:skylake512", title: "Skylake512", group: "CPU" },
        { name: "cpu:broadwell", title: "Broadwell", group: "CPU" },
        { name: "cpu:znver1", title: "Ryzen", group: "CPU" },
        { header: "GPU" },
        { name: "nvcap:5.2", title: "Tesla M60", group: "GPU" },
        { name: "nvcap:7.0", title: "Tesla V100", group: "GPU" },
        { name: "nvcap:6.1", title: "GTX 1060", group: "GPU" },
        { header: "CUDA" },
        { name: "cuda:9.0", title: "9.0", group: "CUDA" },
        { name: "cuda:9.2", title: "9.2", group: "CUDA" },
        { name: "cuda:10.0", title: "10.0", group: "CUDA" },
        { header: "OS" },
        { name: "os:ubuntu18.04", title: "Ubuntu 18.04", group: "OS" },
        { name: "os:centos7", title: "CentOS 7", group: "OS" }
      ]
    },
    formatFeature(name) {
      const list = this.getCommonFeatures();
      for (var i in list) {
        const f = list[i];
        if (f.name == name) {
          return f.title;
        }
      }
      return name;
    },
  }
});
