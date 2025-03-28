run:
	# the "-" sign is to ignore errors
	-make down
	docker compose --profile app up --build

rund:
	-make down
	docker compose --profile app up -d --build
	docker compose ps

run-otel:
	-make down
	docker compose up

down:
	docker compose --profile app down --remove-orphans

zipcode_sp:
	curl -X POST http://localhost:8081/weather -H "Content-Type: application/json" -d '{"cep": "01001000"}' --verbose

zipcode_poa:
	curl -X POST http://localhost:8081/weather -H "Content-Type: application/json" -d '{"cep": "90010000"}' --verbose

zipcode_invalid:
	curl -X POST http://localhost:8081/weather -H "Content-Type: application/json" -d '{"cep": "9a010000"}' --verbose

zipcode_not_found:
	curl -X POST http://localhost:8081/weather -H "Content-Type: application/json" -d '{"cep": "00000000"}' --verbose