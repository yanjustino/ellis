# Name of the project
# choco install make
PROJECT_DIR := ${CURDIR}/dist
DIST_FOLDER=dist

# go build command
GB = go build -o

install-windows:
	@echo "Creating dist folder $(DIST_FOLDER)"
	@ if not exist $(DIST_FOLDER) mkdir $(DIST_FOLDER)
	@echo "Compiling ellis"
	@ $(GB) $(DIST_FOLDER)/ellis.exe
	@ set PATH=$(PATH);$(PROJECT_DIR)
	@echo "Done!"

install:
	@ echo "Creating dist folder $(DIST_FOLDER)"
	@ mkdir -p $(DIST_FOLDER)

	@ echo "Compiling ellis"
	@ $(GB) $(DIST_FOLDER)/ellis
	@ echo $0
	@ export PATH=$(PATH):$(PROJECT_DIR)/ellis
	@ echo "Done!"

.PHONY: install