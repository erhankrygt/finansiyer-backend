cover: cover-profile cover-html

check-swagger:
	which swagger || go install github.com/go-swagger/go-swagger/cmd/swagger@latest

swagger: check-swagger
	swagger generate spec --exclude-deps -o ./docs/swagger.yaml --scan-models

serve-swagger: swagger
	swagger serve -F=swagger ./docs/swagger.yaml

test:
	@go test -v ./...

cover-profile:
	@go test -v -coverprofile cover.out ./...

cover-html:
	@go tool cover -html=cover.out -o cover.html