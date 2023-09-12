# apiclient

## Extract go API client template

```sh
docker run --rm -v $PWD:/local openapitools/openapi-generator-cli author template -g go -o /local/template/go
```

## Generate Tokopedia API Client Library

```sh
docker run --rm -v $PWD:/local openapitools/openapi-generator-cli generate -i /local/oas/tokopedia.yml -g go -o /local/tokopedia -t /local/template/go-tokopedia --git-host github.com --git-user-id dhimas-sirclo --git-repo-id apiclient --package-name tokopedia --additional-properties=packageName=tokopedia,generateInterfaces=true,enumClassPrefix=true && cd tokopedia && go mod tidy && cd ..
```
