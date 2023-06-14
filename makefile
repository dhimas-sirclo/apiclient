tokopedia:
	openapi-generator-cli generate -g go -i ./oas/tokopedia.yml -o ./tokopedia --package-name tokopedia -t ./template/go --additional-properties=generateInterfaces=true --additional-properties=isGoSubmodule=true --git-repo-id apiclient --git-host github.com --git-user-id dhimas-sirclo && openapi-generator-cli generate -g go -i ./oas/tokopedia.auth.yml -o ./tokopediaauth --package-name tokopediaauth -t ./template/go --additional-properties=generateInterfaces=true --additional-properties=isGoSubmodule=true --git-repo-id apiclient --git-host github.com --git-user-id dhimas-sirclo

generate:
	make tokopedia

template-go:
	openapi-generator-cli author template -g go -o ./template/go

template:
	make template-go