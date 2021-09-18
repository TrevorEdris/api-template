# api-template

Template repo used for the ease of creation of API's utilizing Golang, with Kubernetes for deployment.

## Local Development

For local development, use the `make (up|restart|down|logs)` commands provided by the Makefile.

```md
❯ make help
. . .
up            Run the API locally and print logs to stdout
down          Stop all containers
restart       Restart all containers
logs          Print logs in stdout
. . .
```

The above commands use `deployments/local/docker-compose.dev.yaml` to run the API. The binary for the API will be rebuilt automatically
when a change to one of the source `.go` files is detected (configurable in `.air.toml`).

```bash
api    | app/api/v1/endpoints.go has changed
api    | building...
api    | running...
api    | {"lvl":"info","module":"api-template","msg":"Initializing application","serverConfig":{"Port":8080,"ReadTimeout":5,"WriteTimeout":10,"IdleTimeout":120,"JwtIssuer":"/opt/tedris/jwt/priv.pem"},"t":"2021-09-12T23:02:16.5528522Z"}
```

See [https://github.com/cosmtrek/air](https://github.com/cosmtrek/air) for more details.

**Note:** The live-reload feature of `cosmtrek/air` only works for go files.

### Configuration

To configure parameters for the local instance of the API, copy the `sample.env` file into `.env`. The API uses [`joho/godotenv`](https://github.com/joho/godotenv) to read environment variables from this file and apply them to the container at runtime.
Once the environment variables are set, the API will then parse the environment variables using [`kelseyhightower/envconfig`](https://github.com/kelseyhightower/envconfig).

## SLO

| Endpoint | Requests/s | 95%  |
|---|---|---|
| `/v1/health` | ??? | ???ms |
| `/v1/generalkenobi` | ??? | ???ms |

## Data Model

N/A

## Authentication

Authentication of requests is performed by validating the JWT provided in the HTTP Request headers.

```json
{"Authorization": "Bearer <JWT_HERE>"}
```

## Endpoints

### GET `/v1/health`

Perform a health-check, ensuring that all necessary components (3rd party dependencies, databases, etc.) are available and operational.

#### Request Body

N/A

#### Response Body

```json
{
    "msg": "Healthy"
}
```

#### Error Codes

* `401 Unauthorized`: Invalid credentials were provided in the request
* `500 Internal Error`: The health check failed for some reason

### GET `/v1/generalkenobi`

A sample endpoint that simply responds as if General Kenobi met General Grevious.

#### Request Body

N/A

#### Response Body

```json
{
    "msg": "Hello there."
}
```

#### Error Codes

* `401 Unauthorized`: Invalid credentials were provided in the request
* `500 Internal Error`: General Kenobi was killed by Order 66

## Deployment Procedure

Deployment of this API to kubernetes is made simpler with the following Makefile targets.

```md
❯ make help
. . .
build         Build and tag the docker container for the API
test          Run unit tests
finalize      Build, test, and tag the docker container with the finalized tag (typically, the full docker registery will be tagged here)
publish_only  Push the tagged docker image to the docker registry
publish       Finalize and publish the docker container
deploy_only   Fill out the .yaml.tmpl files and apply them to the specified namespace
deploy        Build, test, finalize, publish, and then deploy the docker container to kube
```

**Notes:**

* The default namespace is `tedris`.
* The default region is `dev`.

### Deploying to dev

1. `git checkout vX.Y.Z`
2. `make deploy NAMESPACE=<some_namespace_here>`

### Deploying to a non-dev environment

1. `git checkout vX.Y.Z`
2. `make deploy REGION=<some_non-dev_region>`

### Deploying _without_ running any tests

1. `git checkout vX.Y.Z`
2. `make deploy_only REGION=<some_region> NAMESPACE=<some_namespace>`

## Common Maintenance

_What common, repeated actions are necessary to ensure this API continues to run?_

_Do any API keys need rotated frequently? Does any data need to be deleted at some interval? etc._

## List of 3rd Party Libraries

The following is a list of all 3rd party libraries in use by this API

* _TODO_

## TODO

- [ ] Implement JWT auth
- [ ] Fix live-reload crash on API crash
- [ ] Ensure kube deployment works properly
- [ ] List of 3rd party libraries
