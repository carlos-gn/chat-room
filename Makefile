dev:
	air

setup-db:
	touch chat.sqlite && make migrate

setup-test-db:
	touch test_chat.sqlite && make migrate

migrate:
	sqlite3 test_chat.sqlite < migrations/init.sql

build:
	go build .

docker-build:
	docker build --tag go-api .

docker-run:
	docker run --publish 3000:3000 go-api

test:
	ginkgo -vv ./...
