#!/bin/bash

tmp_dir=$(mktemp -d -t build-XXXXXXXXXX)
echo ">> TEMPDIR: $tmp_dir"
yq d yml/template.yml manifests >${tmp_dir}/mh-test.yml 
yq w -i ${tmp_dir}/mh-test.yml image docker.io/qnib/metahub-demo:$(date +%F).2 
INSTS="c5 c5n c5a c6g"
SIZES="9xl 12xl 16xl 18xl 24xl"
HTS="on off"
INSTS="c5"
SIZES="18xl"
HTS="on off"
for i in $(echo $INSTS);do
  for s in $(echo ${SIZES});do
    for ht in ${HTS};do
      IMG_TAG=${i}_${s}_${ht}
      IMG_NAME=docker.io/qnib/metahub-single:${IMG_TAG}
      echo ">>> Build ${IMG_NAME}"
      docker build --quiet -t ${IMG_NAME} \
                   --build-arg=INST=${i} \
                   --build-arg=SIZE=${s} \
                   --build-arg=HT=${ht} \
                   .
      docker push ${IMG_NAME} >/dev/null
      yq w -i ${tmp_dir}/mh-test.yml 'manifests[+].image' ${IMG_NAME}
      echo ">> ADD {image: 'IMG_NAME', platform: {architecture: 'amd64',os: 'linux', features: ['instance:$i','size:$s','ht:$ht']}}"
      yq r -j ${tmp_dir}/mh-test.yml \
          |jq -r --arg IMGNAME "${IMG_NAME}" --arg INSTR "instance:$i" --arg SIZESTR "size:$s" --arg HTSTR "ht:$ht" \
              '.manifests |= map(select(.image==$IMGNAME)+= {platform: {architecture: "amd64",os: "linux", features: [$INSTR,$SIZESTR, $HTSTR]}})' \
          |yq r --prettyPrint - >${tmp_dir}/mh-test_${IMG_TAG}.yml
      sleep 0.2
    done
  done
done
echo ">> Merge temp files"
yq merge ${tmp_dir}/mh-test_* > ${tmp_dir}/mh-test.yml
echo ">> Push manifest: ${tmp_dir}/mh-test.yml"
manifest-tool push from-spec ${tmp_dir}/mh-test.yml
#rm -rf ${tmp_dir}
