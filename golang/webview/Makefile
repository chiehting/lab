# 變數定義
BINARY_NAME=browser-tool
SOURCE_DIR=./main.go
BUILD_DIR=./

.PHONY: build run openwebui

build:
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SOURCE_DIR)

run:
	go run $(SOURCE_DIR)

openwebui:
	make build
	mkdir -p webui.app/Contents/MacOS webui.app/Contents/Resources
	cp -rf open-webui/icon.png webui.app/Contents/Resources
	cp -rf open-webui/info.plist webui.app/Contents
	mv -f $(BUILD_DIR)/$(BINARY_NAME) webui.app/Contents/MacOS