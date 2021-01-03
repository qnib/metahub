# Branch `master` is archived, please use `main`

## MetaHub - Dynamic OCI Registry Proxy

Announcement: [qnib.org](http://www.qnib.org/2019/06/12/metahub/)

The MetaHub project is meta-data registry, which serves images filtered via login so that a machine gets the image that
fits the specifics of the host the image is going to run on.

That could be picking an image that not only fits the CPU Architecture (x86-64, ppcle, arm) but is optimized for the microarchitecture 
of the host (Broadwell, Skylake, ...). And it does not stop there - any host specific attribute can be use: 
Accelerators, network, configuration or the full depth of gcc options. 

![](/misc/pics/metahub-proxy.png)

## Example

A machine that logs in with the user `qnib-type1` will get a Broadwell image, `qnib-type2` serves a Skylake optimized image.

```
$ docker login -u qnib-type1 -p qnib-type1 metahub.qnib.org
$ docker run metahub.qnib.org/qnib/bench && docker run -ti --rm metahub.qnib.org/qnib/bench
Unable to find image 'metahub.qnib.org/qnib/bench:latest' locally
latest: Pulling from qnib/bench
Status: Downloaded newer image for metahub.qnib.org/qnib/bench:latest
>> This container is optimised for: cpu:broadwell
$ docker inspect  metahub.qnib.org/qnib/bench |jq '.[].RepoDigests[1]
"metahub.qnib.org/qnib/bench@sha256:f972f05f4ff0c5df22c1f4fc339b14068b8cee96d0525f4fb13dbea84a900c89"
$ docker login -u qnib-type2 -p qnib-type2 metahub.qnib.org
Login Succeeded
$ docker run metahub.qnib.org/qnib/bench && docker run -ti --rm metahub.qnib.org/qnib/bench
Unable to find image 'metahub.qnib.org/qnib/bench:latest' locally
latest: Pulling from qnib/bench
Status: Downloaded newer image for metahub.qnib.org/qnib/bench:latest
>> This container is optimised for: cpu:skylake
$ docker inspect  metahub.qnib.org/qnib/bench |jq '.[].RepoDigests[1]
"metahub.qnib.org/qnib/bench@sha256:c62049e0707d3461b0a5b69d0ebe322f16a17b4439b5d9844e24cf97d40faa64"
``` 

It does that by filtering the manifest list and only present the image that fits the demanded features.

```
$cat manifest.yml
image: docker.io/qnib/bench
manifests:
  -
    image: docker.io/qnib/plain-featuretest:cpu-skylake
    platform:
      architecture: amd64
      os: linux
      features:
        - cpu:skylake
  -
    image: docker.io/qnib/plain-featuretest:cpu-broadwell
    platform:
      architecture: amd64
      os: linux
      features:
        - cpu:broadwell
```

![](/misc/pics/metahub-overview.png)

