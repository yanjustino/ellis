# Name of the project
# choco install make
PROJECT_DIR := ${CURDIR}/bin

# go build command
GB = go build -o

install-windows:
	@echo "Creating bin folder $(PROJECT_DIR)"
	@if not exist bin mkdir bin
	@echo "Compiling ellis"
	@$(GB) bin/ellis.exe
	@echo "Done!"

install:
	@echo "Creating bin folder $(PROJECT_DIR)"
	@mkdir -p bin
	@echo "Compiling ellis"
	@$(GB) bin/ellis
	@export PATH=${CURDIR}/bin:$(PATH)
	@echo "Done!"

.PHONY: install