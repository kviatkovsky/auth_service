ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

run:
	go run cmd/app/main.go --config=./config/local.yaml

migrate-db:
	docker run -v ./db/migrations:/migrations --network host migrate/migrate -path=/migrations -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)?query" up
																													
migrate-db-down:
	docker run -v ./db/migrations:/migrations --network host migrate/migrate -path=/migrations -database "mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)?query" down


.PHONY: run migrate-db migrate-db-down