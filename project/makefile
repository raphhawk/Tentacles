FRONT_END_BINARY=frontApp
BROKER_BINARY=brokerApp

up:
	@echo "Starting Docker images..."
	sudo docker-compose up -d
	@echo "Docker images started!"

up_build: build_broker
	@echo "Stopping docker images (if running)..."
	sudo docker-compose down
	@echo "Building (when required) and starting docker images..."
	sudo docker-compose up --build -d
	@echo "Docker images built and started!"

down:
	@echo "Stopping docker compose..."
	sudo docker-compose down
	@echo "Done!"

build_broker:
	@echo "Building broker binary..."
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"

build_front:
	@echo "Building front end binary..."
	cd ../front-end && env CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} ./cmd/web
	@echo "Done!"

start: build_front
	@echo "Starting front end"
	cd ../front-end && go build -o ${FRONT_END_BINARY} ./cmd/web
	cd ../front-end && sudo ./${FRONT_END_BINARY} 