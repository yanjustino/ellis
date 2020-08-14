# Name of the project
# choco install make
PROJECT_DIR := ${CURDIR}/dist
DIST_FOLDER=dist
PROJECT_NAME=ellis.exe



# go build command
GB = go build -o

install: build

build:
	@echo "Creating dist folder $(DIST_FOLDER)"
	@ if not exist $(DIST_FOLDER) mkdir $(DIST_FOLDER)

	@echo "Compiling ellis"
	@ $(GB) $(DIST_FOLDER)/$(PROJECT_NAME)
	@ set PATH=$(PATH);$(PROJECT_DIR)

	@echo "Done!"

.PHONY: install