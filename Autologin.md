## Autologin

```bash
docker pull -q public.ecr.aws/a4y4t9s0/metahub-autologin:v0.2.9
export MH_USER=$(docker run --rm --network=host public.ecr.aws/a4y4t9s0/metahub-autologin:v0.2.9 -get-user)
echo "> MH_USER: ${MH_USER}"
export MH_PASS=$(docker run --rm --network=host public.ecr.aws/a4y4t9s0/metahub-autologin:v0.2.9 -get-pass -aws-region=eu-west-1)
echo "> MH_PASS: ${MH_PASS}"
echo ${MH_PASS} |docker login --username ${MH_USER} --password-stdin mh.qnib.org
```

## Autologin on AWS instance

First we pull the latest image.

```bash
$ docker pull -q public.ecr.aws/a4y4t9s0/metahub-autologin:v0.2.9
```

As we are running on an `c5.9xlarge` instance with hyperthreading enabled, the username generated is `metahub--c59xl-ht`.

```bash
$ docker run -ti --rm --network=host public.ecr.aws/a4y4t9s0/metahub-autologin:v0.2.7_dirty -get-user
metahub-c59xl-ht
```

To login to metahub, we need to present the docker socket. As the login information are local to the user, we'll map in the `.docker` directory to persist the login.
The `AWS_REGION` defines the parameter store to fetch the password from `/metahub/password`.

```bash
$ docker run -e AWS_REGION=eu-west-1 -ti --rm --network=host -v ${HOME}/.docker:/root/.docker -v /var/run/docker.sock:/var/run/docker.sock public.ecr.aws/a4y4t9s0/metahub-autologin:v0.2.7_dirty
2021/01/05 16:30:06 Use docker login via client
2021/01/05 16:30:08 Login Succeeded
```

From now on we'll be able to fetch images from metahub.

```bash
$ cat .docker/config.json
{
	"auths": {
		"mh.qnib.org": {
			"auth": "bWV0YWh1Yi1jNTl4bC1odDpncm9tYWNzLXNhcnVzMTIz"
		}
	},
	"HttpHeaders": {
		"User-Agent": "Docker-Client/19.03.12 (linux)"
	}
}
$ docker pull -q mh.qnib.org/qnib/metahub-demo:2021-01-04.2
mh.qnib.org/qnib/metahub-demo:2021-01-04.2
$ docker run -ti --rm mh.qnib.org/qnib/metahub-demo:2021-01-04.2
>> This image is optimized for 'c5.9xl' with hyperthreading turned 'on'
$
```

Instead of autogenerate the login, the type can also be enforced.

```bash
$ docker run -ti --rm --network=host public.ecr.aws/a4y4t9s0/metahub-autologin:v0.2.9 -get-user -type c524xl
metahub-c524xl
$ docker run -e AWS_REGION=eu-west-1 -ti --rm --network=host -v ${HOME}/.docker:/root/.docker \
             -v /var/run/docker.sock:/var/run/docker.sock public.ecr.aws/a4y4t9s0/metahub-autologin:v0.2.9 -docker-login -type=c524xl
$ docker rmi mh.qnib.org/qnib/metahub-demo:2021-01-04.2
$ docker pull -q mh.qnib.org/qnib/metahub-demo:2021-01-04.2
mh.qnib.org/qnib/metahub-demo:2021-01-04.2
$ docker run -ti --rm mh.qnib.org/qnib/metahub-demo:2021-01-04.2
>> This image is optimized for 'c5.24xl' with hyperthreading turned 'off'
```