## Autologin

```bash
docker pull -q public.ecr.aws/a4y4t9s0/metahub-autologin:v0.2.7
export MH_USER=$(docker run --rm --network=host public.ecr.aws/a4y4t9s0/metahub-autologin:v0.2.7 -get-user)
echo "> MH_USER: ${MH_USER}"
export MH_PASS=$(docker run --rm --network=host public.ecr.aws/a4y4t9s0/metahub-autologin:v0.2.7 -get-pass -aws-region=eu-west-1)
echo "> MH_PASS: ${MH_PASS}"
echo ${MH_PASS} |docker login --username ${MH_USER} --password-stdin mh.qnib.org
```

```bash
docker pull -q public.ecr.aws/a4y4t9s0/metahub-autologin:v0.2.7
docker run -ti --rm --network=host -v /var/run/docker.sock:/var/run/docker.sock public.ecr.aws/a4y4t9s0/metahub-autologin:v0.2.7 autologin -docker-login -aws-region=eu-west-1
docker pull -q mh.qnib.org/qnib/metahub-demo:2021-01-04.2
docker run -ti --rm mh.qnib.org/qnib/metahub-demo:2021-01-04.2
```