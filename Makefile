run:
	go run cmd/app/main.go

image:
	docker build -t effective .

up:
	docker compose up

down:
	docker compose down

clean:
	docker system prune -a
