.PHONY: run
run: ## run the API server
	go run cmd/server/main.go
.PHONY: db_start
db_start:
	docker run --rm --name basic_golang -v $(shell pwd)/migrations/:/docker-entrypoint-initdb.d \
        -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=go_restful -e POSTGRES_USER=postgres-dev -d -p 5432:5432 postgres
.PHONY: db_stop
db_stop:
	docker stop basic_golang