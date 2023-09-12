tokopedia:
	docker run --rm -v $PWD:/local openapitools/openapi-generator-cli generate -i /local/oas/tokopedia.yml -g go -o /local/tokopedia -t /local/template/go-tokopedia --git-host github.com --git-user-id dhimas-sirclo --git-repo-id apiclient --package-name tokopedia --additional-properties=packageName=tokopedia,generateInterfaces=true,enumClassPrefix=true

generate:
	make tokopedia

template-go:
	docker run --rm -v $PWD:/local openapitools/openapi-generator-cli author template -g go -o /local/template/go

template:
	make template-go