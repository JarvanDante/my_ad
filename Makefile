APP := my_ad
BIN := bin
MYDOCKER := /Users/wangdante/D/mydocker

.PHONY: build dev tidy clean migrate docker-up docker-down docker-logs docker-rebuild

build:
	go build -o $(BIN)/adapi .

dev:
	gf run main.go

tidy:
	go mod tidy

migrate:
	goose -dir manifest/sql/migrations postgres "host=127.0.0.1 port=5432 user=postgres password=654321 dbname=my_ad sslmode=disable" up

clean:
	rm -rf $(BIN)

docker-up:
	cd $(MYDOCKER) && docker compose up -d --build my_ad
	@echo "探活: curl -sS http://127.0.0.1:8016/healthz"

docker-down:
	cd $(MYDOCKER) && docker compose stop my_ad

docker-logs:
	docker logs -f my_ad

docker-rebuild:
	cd $(MYDOCKER) && docker compose up -d --build --force-recreate my_ad
