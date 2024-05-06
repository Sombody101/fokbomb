# Check if Go is installed
ifeq (, $(shell which go))
        $(error Go is not installed!)
endif

PACKAGE := fokbomb

output_path ?= ../build/$(PACKAGE)

LDFLAGS := -X $(PACKAGE)/main.__DEBUG_str=false

# General build that targets current OS
build:
	go build -C ./src/ -ldflags \
		"-w -s $(LDFLAGS)" \
		-trimpath \
		-o $(output_path)
	@echo Built $(output_path)

lin:
	GOOS=windows go \
		build -C ./src/ \
		-trimpath \
		-ldflags "-w -s -X main.__DEBUG_str=false" \
		-o $(output_path)
	@echo Built $(output_path)

# win[dows]
win:
	GOOS=windows go \
		build -C ./src/ \
		-trimpath \
		-ldflags "-w -s -X main.__DEBUG_str=false" \
		-o $(output_path).exe
	@echo Built $(output_path).exe

# d[bug]win[down] (uses d[ebug]_LDFLAGS)
dwin:
	GOOS=windows go build -C ./src/ -o $(output_path)_debug.exe
	@echo Built $(output_path)_debug.exe

garbled:
	bash ./gargle.make.sh

all: build win dwin garbled

.PHONY: build