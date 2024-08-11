.PHONY: migrate-up migrate-down seed \
	set-migration-env set-docker-env \
	build-app clean-app

migrate-up:
	go run cmd/cli/*.go migrate up

migrate-down:
	go run cmd/cli/*.go migrate down

seed:
	go run cmd/cli/*.go seed

# Detect the operating system
ifeq ($(OS),Windows_NT)
    SHELL := powershell.exe
    .SHELLFLAGS := -NoProfile -Command
    SCRIPT_EXT := .ps1
else
    SHELL := /bin/bash
    SCRIPT_EXT := .sh
endif

YAML_FILE := config/app.yaml
SCRIPT_FILE := generate_config$(SCRIPT_EXT)

set-docker-env:
	@echo "Generating YAML file for Docker environment..."
ifeq ($(OS),Windows_NT)
	@powershell -Command ".\$(SCRIPT_FILE) docker > $(YAML_FILE)"
else
	@bash $(SCRIPT_FILE) docker > $(YAML_FILE)
endif
	@echo "YAML file updated for Docker environment."

set-migration-env:
	@echo "Generating YAML file for migration environment..."
ifeq ($(OS),Windows_NT)
	@powershell -Command ".\$(SCRIPT_FILE) local > $(YAML_FILE)"
else
	@bash $(SCRIPT_FILE) local > $(YAML_FILE)
endif
	@echo "YAML file updated for migration environment."

build-app :
	docker-compose build --no-cache
	docker-compose up -d

clean-app :
	docker-compose down --rmi all --volumes
	docker system prune -f
