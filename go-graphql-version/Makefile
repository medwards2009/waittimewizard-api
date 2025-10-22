docker-build:
	docker buildx build --platform linux/amd64,linux/arm64 -t medwards2009/waittimewizard-api:latest --push .

docker-run:
	docker run -d -p 8080:8080 medwards2009/waittimewizard-api

docker-push:
	docker push medwards2009/waittimewizard-api

docker-pull:
	docker pull medwards2009/waittimewizard-api:latest

gql-gen:
	go run github.com/99designs/gqlgen generate

gql-init:
	go run github.com/99designs/gqlgen init

dev:
	go run server.go

test:
	go test ./... -v

test-coverage:
	go test -cover ./...

test-report:
	go test -coverprofile=coverage.out ./...

view-coverage:
	go tool cover -html=coverage.out