## Autologin

```bash
docker pull -q public.ecr.aws/a4y4t9s0/metahub:autologin_v0.2.6
export MH_USER=$(docker run --rm --network=host public.ecr.aws/a4y4t9s0/metahub:autologin_v0.2.6 -get-user)
echo "> MH_USER: ${MH_USER}"
export MH_PASS=$(docker run --rm --network=host public.ecr.aws/a4y4t9s0/metahub:autologin_v0.2.6 -get-pass -aws-region=eu-west-1)
echo "> MH_PASS: ${MH_PASS}"
echo ${MH_PASS} |docker login --username ${MH_USER} --password-stdin mh.qnib.org
```