.PHONY: build clean install run

# Build com permissões de execução
build:
	go build -o kraken -ldflags="-s -w" .
	chmod +x kraken

# Instalar no sistema (opcional)
install: build
	sudo cp kraken /usr/local/bin/

# Limpar arquivos de build
clean:
	rm -f kraken
	rm -rf docs/kraken/*.md

# Executar o kraken
run: build
	./kraken

# Build para diferentes plataformas
build-linux:
	GOOS=linux GOARCH=amd64 go build -o kraken-linux -ldflags="-s -w" .
	chmod +x kraken-linux

build-windows:
	GOOS=windows GOARCH=amd64 go build -o kraken.exe -ldflags="-s -w" .

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o kraken-mac -ldflags="-s -w" .
	chmod +x kraken-mac

# Development
dev:
	go run main.go

# Testes
test:
	go test ./...

# Formatar código
fmt:
	go fmt ./...

# Verificar código
vet:
	go vet ./...

# Tudo
all: fmt vet test build
