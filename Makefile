start_web:
	docker compose run --service-ports web bash

run:
	docker compose up

build:
	docker compose build