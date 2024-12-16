docker-build:
	docker buildx build --platform linux/amd64,linux/arm64 -t medwards2009/waittimewizard-api:latest --push .

docker-run:
	docker run -d -p 8080:8080 medwards2009/waittimewizard-api

docker-push:
	docker push medwards2009/waittimewizard-api

docker-pull:
	docker pull medwards2009/waittimewizard-api:latest