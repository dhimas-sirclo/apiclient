# apiclient

## Publishing a module

visit [Publishing a module](https://go.dev/doc/modules/publishing) and follow the guide

## Extract go API client template

```sh
docker run --rm -v $PWD:/local openapitools/openapi-generator-cli author template -g go -o /local/template/go
```

## Generate Tokopedia API Client Library

```sh
docker run --rm -v $PWD:/local openapitools/openapi-generator-cli generate -i /local/oas/tokopedia.yml -g go -o /local/tokopedia -t /local/template/go-tokopedia --git-host github.com --git-user-id dhimas-sirclo --git-repo-id apiclient/tokopedia --package-name tokopedia --additional-properties=packageName=tokopedia,generateInterfaces=true,enumClassPrefix=true && cd tokopedia && go mod tidy && cd ..
```

## Run Tokopedia Mock API

```sh
docker compose --file ./docker/tokopedia.mock.docker-compose.yml up -d
```
