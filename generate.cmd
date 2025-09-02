go install golang.org/x/tools/cmd/goimports@latest

docker run --rm ^
    -v "%cd%":/local ^
    openapitools/openapi-generator-cli:v7.15.0 generate ^
    -g go ^
    --git-user-id eliona-smart-building-assistant ^
    --git-repo-id go-eliona-api-client/v2 ^
    -i https://raw.githubusercontent.com/eliona-smart-building-assistant/eliona-api/main/openapi.yaml ^
    -o /local ^
    --additional-properties="packageName=api"

go mod tidy
FOR /R %%f IN (*.go) DO (goimports -w "%%f")
go vet ./...
