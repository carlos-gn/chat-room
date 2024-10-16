dev:
	air

migrate:
	sqlite < migrations/init.sql

build: 
	go build .

docker-build:
	docker build --tag go-api . 

docker-run:
	docker run --publish 3000:3000 go-api

test:
	ginkgo -vv ./...
