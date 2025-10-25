# List all available commands
default:
    @just --list

# Build the Docker image
build:
    docker build -t waittimewizard-api .

# Run the Docker container in production mode
run: build
    docker run -p 8080:8080 waittimewizard-api

# Start the development environment using docker-compose
dev:
    docker-compose up --build

# Stop all running containers
stop:
    docker-compose down

# Clean up Docker resources
clean:
    docker-compose down -v
    docker system prune -f

# Run tests in Docker
test:
    docker-compose run --rm api npm test

# Generate a production Docker image with a specific tag
release tag:
    docker buildx build --platform linux/amd64,linux/arm64 --load -t medwards2009/waittimewizard-api:{{tag}} .
    docker tag medwards2009/waittimewizard-api:{{tag}} medwards2009/waittimewizard-api:latest

push tag:
    docker push medwards2009/waittimewizard-api:{{tag}}
    docker push medwards2009/waittimewizard-api:latest
    

# Open a shell in the running container
shell:
    docker-compose exec api sh

# View logs from the running container
logs:
    docker-compose logs -f api

# Check the health of the running container
health:
    docker-compose ps api