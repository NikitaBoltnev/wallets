build:
	docker compose build

up:
	docker compose up -d

down:
	docker compose down

clean:
	docker compose down -v

test:
	chmod +x test_api.sh
	./test_api.sh