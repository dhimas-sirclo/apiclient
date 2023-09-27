# apiclient

## Requirements

1. Docker Engine or Desktop
2. OpenAPI Generator CLI Docker Image v7
3. Go >= v19
4. NPM

### Get OpenAPI Generator CLI Docker Image v7

```sh
docker pull openapitools/openapi-generator-cli:v7.0.0
```

## Publishing modules

visit [Publishing a module](https://go.dev/doc/modules/publishing) and follow the guide

For major release

```sh
npm run release:major
```

For minor release

```sh
npm run release:minor
```

For patch release

```sh
npm run release:patch
```

Make module available

```sh
GOPROXY=proxy.golang.org go list -m github.com/dhimas-sirclo/apiclient/tokopedia
GOPROXY=proxy.golang.org go list -m github.com/dhimas-sirclo/apiclient/sirclo
GOPROXY=proxy.golang.org go list -m github.com/dhimas-sirclo/apiclient/connector
```

## Extract go API client template

```sh
docker run --rm -v $PWD:/local openapitools/openapi-generator-cli:v7.0.0 author template -g go -o /local/template/go
```

## Generate Tokopedia API Client Library

```sh
docker run --rm -v $PWD:/local openapitools/openapi-generator-cli:v7.0.0 generate -i /local/oas/tokopedia.yml -g go -o /local/tokopedia -t /local/template/go-tokopedia --git-host github.com --git-user-id dhimas-sirclo --git-repo-id apiclient/tokopedia --package-name tokopedia --additional-properties=packageName=tokopedia,generateInterfaces=true,enumClassPrefix=true,structPrefix=true && cd tokopedia && go mod tidy && cd ..
```

## Run Tokopedia Mock API

```sh
docker compose --file ./docker/tokopedia.mock.docker-compose.yml up -d
```

## Generate SIRCLO API Client Library

```sh
docker run --rm -v $PWD:/local openapitools/openapi-generator-cli:v7.0.0 generate -i /local/oas/sirclo.yml -g go -o /local/sirclo -t /local/template/go --git-host github.com --git-user-id dhimas-sirclo --git-repo-id apiclient/sirclo --package-name sirclo --additional-properties=packageName=sirclo,generateInterfaces=true,enumClassPrefix=true && cd sirclo && go mod tidy && cd ..
```

## Run SIRCLO Mock API

```sh
docker compose --file ./docker/sirclo.mock.docker-compose.yml up -d
```

## Extract go gin server stubs template

```sh
docker run --rm -v $PWD:/local openapitools/openapi-generator-cli:v7.0.0 author template -g go-gin-server -o /local/template/go-gin-server
```

## Generate Connector Server Stubs

```sh
docker run --rm -v $PWD:/local openapitools/openapi-generator-cli:v7.0.0 generate -i /local/oas/connector.yml -g go-gin-server -o /local/connector -t /local/template/go-gin-server --git-host github.com --git-user-id dhimas-sirclo --git-repo-id apiclient/connector --package-name connector --additional-properties=packageName=connector,enumClassPrefix=true,isConnector=true
```
