
generate:
	@echo "Generating API code..."
	oapi-codegen -generate types -package gen wallet.yaml > internal\server\gen\models.gen.go
	oapi-codegen -generate server -package gen wallet.yaml > internal\server\gen\server.gen.go

clean:
	if exist server\gen rmdir /s /q server\gen

dev: generate
	go run cmd\server\main.go