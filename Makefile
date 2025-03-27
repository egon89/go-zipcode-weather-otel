run:
	# the "-" sign is to ignore errors
	-make down
	docker compose up --build
	
down:
	docker compose down --remove-orphans

zipcode_sp:
	curl -X POST http://localhost:8081/weather -H "Content-Type: application/json" -d '{"cep": "01001000"}' --verbose

zipcode_poa:
	curl -X POST http://localhost:8081/weather -H "Content-Type: application/json" -d '{"cep": "90010000"}' --verbose
