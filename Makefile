# Name of the project
# choco install make
PROJECT_DIR := ${CURDIR}/bin

# go build command
GB = go build -o

install:
	@echo "Creating bin folder $(PROJECT_DIR)"
	@mkdir -p bin
	@echo "Compiling ellis"
	@$(GB) bin/ellis
	@echo $0
	@export PATH=$(PATH):${CURDIR}/bin
	@echo "Done!"

install-windows:
	@echo "Creating bin folder $(PROJECT_DIR)"
	@if not exist bin mkdir bin
	@echo "Compiling ellis"
	@$(GB) bin/ellis.exe
	@echo "Done!"

.PHONY: install