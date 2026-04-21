include .env
export


export PROJECT_ROOT
env-up:
	@docker compose up -d

env-down:
	@docker compose down -v

info: 
	@docker compose ps -a
	
env-cleanup:
	@read -p "Delete all volume data? This will erase database files [y/N] :" ans; \
	if [ "$$ans" = "y" ]; then\
		docker compose down crypto-postgres && \
		rm -rf out/pgdata &&\
		echo "Environment files cleaned";\
		else\
		echo "Cleanup cancelled";\
	fi

env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder
	
migrate-create:
	@if [ -z "$(seq)" ]; then \
	echo "Required parameter is missing. Example: make migrate-create seq=init";\
	exit 1; \
	fi; \
	docker compose run --rm crypto-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Required parameter is missing. Example: make migrate-action action=up/down";\
		exit 1; \
		fi; \
	docker compose run --rm crypto-postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@crypto-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"
