FRONT_END_BINARY=frontApp
BROKER_BINARY=brokerApp
AUTH_BINARY=authApp
LOGGER_BINARY=logApp
MAILER_BINARY=mailApp
LISTENER_BINARY=listenerApp

up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

up_build: build_broker build_auth build_logger build_mailer build_listener 
	@echo "Stopping docker images (if running)..."
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

build_broker: 
	@echo "building broker binary..."
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "done!"

build_auth:
	@echo "building authentication binary..."
	cd ../authentication-service && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd/api
	@echo "done!"

build_logger:
	@echo "building logger binary"
	cd ../logger-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LOGGER_BINARY} ./cmd/api
	@echo "done!"

build_mailer:
	@echo "building mailer binary"
	cd ../mail-service && env GOOS=linux CGO_ENABLED=0 go build -o ${MAILER_BINARY} ./cmd/api
	@echo "done!"

build_listener:
	@echo "building listener binary"
	cd ../listener-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LISTENER_BINARY} .
	@echo "done!"

build_front:
	@echo "Building front end binary..."
	cd ../front-end && env CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} ./cmd/web
	@echo "Done!"

start: build_front
	@echo "Starting front end"
	cd ../front-end && go build -o ${FRONT_END_BINARY} ./cmd/web
	cd ../front-end && sudo ./${FRONT_END_BINARY} 
