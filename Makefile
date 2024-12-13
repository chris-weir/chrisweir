build:
	go build -v -o ./app ./cmd
run:
	go run ./cmd/server.go
test:
	go test ./...
css: 
	cd tailwind; \
	npx tailwindcss -i ./app.css -o ../assets/app.css --minify
docker-prod:
	docker compose -f docker/docker-compose.production.yml up --build
docker:
	docker compose up -d