run:
	TZ=Europe/Vilnius \
	APP_NAME="local.go-fiber-test" \
	ENV="develop" \
  	APP_VERSION="0.0.1" \
  	DATABASE_URL="postgres://postgres:postgres@0.0.0.0:5432/test?connect_timeout=5" \
  	CACHE_URL="localhost:6379" \
  	JWT_SECRET="secret" \
	go run server.go

up:
	docker-compose up -d

down:
	docker-compose down

build:
	docker-compose build