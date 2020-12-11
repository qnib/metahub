import Vue from 'vue'

Vue.mixin({
  methods: {
    getCommonFeatures() {
      return [
        { header: "Size" },
        { name: "size:9xl", title: "9xlarge", group: "Size" },
        { name: "size:12xl", title: "12xlarge", group: "Size" },
        { name: "size:16xl", title: "16xlarge", group: "Size" },
        { name: "size:18xl", title: "18xlarge", group: "Size" },
        { name: "size:24xl", title: "24xlarge", group: "Size" },
        { header: "Instance" },
        { name: "instance:c5", title: "c5", group: "Instances" },
        { name: "instance:c5n", title: "c5n", group: "Instances" },
        { name: "instance:c5a", title: "c5a", group: "Instances" },
        { name: "instance:c6g", title: "c6", group: "Instances" },
        { header: "Hyperthreading" },
        { name: "ht:on", title: "On", group: "Hyperthreading" },
        { name: "ht:off", title: "Off", group: "Hyperthreading" },
        { header: "MPI" },
        { name: "mpi:mpich", title: "MPICH", group: "MPI" },
        { name: "mpi:intelmpi", title: "IntelMPI", group: "MPI" },
        { name: "mpi:openmpi", title: "OpenMPI", group: "MPI" },
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
