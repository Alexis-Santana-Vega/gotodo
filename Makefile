# Variables
BINARY_NAME = gotodo
GO = go
ARGS ?= example title

.PHONY: build run test clean

# Detectar sistema operativo
ifeq ($(OS),Windows_NT)
    RM = del /q
    BINARY_PATH = bin\$(BINARY_NAME)
else
    RM = rm -f
    BINARY_PATH = bin/$(BINARY_NAME)
endif

# Objetivo por defecto
all: build

# Compilar el proyecto
build:
	$(GO) build -o bin/$(BINARY_NAME) main.go

# Ejecutar el proyecto
run:
	$(GO) run .

# Ejecutar pruebas
test:
	$(GO) test -v ./...

# Limpiar binarios
clean:
	$(GO) clean
	$(RM) $(BINARY_PATH)

# Ayuda
help:
	@echo "Available commands:"
	@echo "  make build - Compiles the project"
	@echo "  make run   - Runs the proyect"
	@echo "  make test  - Runs the tests"
	@echo "  make clean - Remove the binary"