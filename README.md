# apiclient

## Requirements

1. Docker Engine or Desktop
2. OpenAPI Generator CLI Docker Image v7
3. Go v18

### Get OpenAPI Generator CLI Docker Image v7

```sh
docker pull openapitools/openapi-generator-cli:v7.0.0
```

## Publishing a module

visit [Publishing a module](https://go.dev/doc/modules/publishing) and follow the guide

## Extract go API client template

```sh
docker run --rm -v $PWD:/local openapitools/openapi-generator-cli:v7.0.0 author template -g go -o /local/template/go
```

## Generate Tokopedia API Client Library

```sh
docker run --rm -v $PWD:/local openapitools/openapi-generator-cli:v7.0.0 generate -i /local/oas/tokopedia.yml -g go -o /local/tokopedia -t /local/template/go-tokopedia --git-host github.com --git-user-id dhimas-sirclo --git-repo-id apiclient/tokopedia --package-name tokopedia --additional-properties=packageName=tokopedia,generateInterfaces=true,enumClassPrefix=true && cd tokopedia && go mod tidy && cd ..
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
