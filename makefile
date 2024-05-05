# Check if Go is installed
ifeq (, $(shell which go))
        $(error Go is not installed!)
endif

PACKAGE := fokbomb

# Local to the root of the project, not the repository
output_path ?= ../build/$(PACKAGE)

# Define ldflags with version information
LDFLAGS := -X $(PACKAGE)/main.__DEBUG_str=false
d_LDFLAGS := -X $(PACKAGE)/main.__DEBUG_str=true

build:
	go build -C ./src/ -ldflags \
		"-w -s $(LDFLAGS)" \
		-o $(output_path)
	@echo Built $(output_path)

win:
	GOOS=windows go build -C ./src/ -ldflags \
		"-w -s $(LDFLAGS)" \
		-o $(output_path).exe
	@echo Built $(output_path).exe

dwin:
	GOOS=windows go build -C ./src/ -ldflags \
		"-w -s $(d_LDFLAGS)" \
		-o $(output_path).exe
	@echo Built $(output_path).exe

.PHONY: build