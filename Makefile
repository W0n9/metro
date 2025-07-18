# Makefile for Metro Go Project

.PHONY: build run clean test install deps help fmt vet check profile info dev \
         build-linux build-linux-arm64 build-windows build-windows-arm64 \
         build-macos build-macos-arm64 build-all build-common build-arm64 \
         build-linux-all build-windows-all build-macos-all

# 默认目标
.DEFAULT_GOAL := help

# 变量定义
BINARY_NAME=metro
BUILD_DIR=build
SOURCE_DIR=.
MAIN_FILE=main.go

# 帮助信息
help: ## 显示帮助信息
	@echo "北京地铁查询工具 (Go 版本)"
	@echo ""
	@echo "可用命令:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "  %-20s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# 安装依赖
deps: ## 安装项目依赖
	@echo "正在安装依赖..."
	go mod download
	go mod tidy

# 构建项目
build: deps ## 构建可执行文件
	@echo "正在构建项目..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)
	@echo "构建完成: $(BUILD_DIR)/$(BINARY_NAME)"

# 运行项目 (示例参数)
run: ## 运行项目 (使用示例参数: 西直门 4)
	@echo "运行示例 (西直门 -> 4元预算):"
	go run $(MAIN_FILE) 西直门 4

# 运行测试
test: deps ## 运行测试
	@echo "正在运行测试..."
	go test -v ./...

# 格式化代码
fmt: ## 格式化代码
	@echo "正在格式化代码..."
	go fmt ./...

# 代码检查
vet: ## 运行代码静态检查
	@echo "正在进行代码检查..."
	go vet ./...

# 完整检查 (格式化 + 检查 + 测试)
check: fmt vet test ## 运行完整检查 (格式化 + 静态检查 + 测试)

# 清理构建文件
clean: ## 清理构建文件
	@echo "正在清理构建文件..."
	rm -rf $(BUILD_DIR)
	go clean

# 安装到系统路径
install: build ## 安装到系统路径
	@echo "正在安装到系统路径..."
	cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/
	@echo "安装完成，现在可以直接使用 '$(BINARY_NAME)' 命令"

# 交叉编译 - Linux AMD64
build-linux: deps ## 交叉编译 Linux AMD64 版本
	@echo "正在编译 Linux AMD64 版本..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_FILE)

# 交叉编译 - Linux ARM64
build-linux-arm64: deps ## 交叉编译 Linux ARM64 版本
	@echo "正在编译 Linux ARM64 版本..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 $(MAIN_FILE)

# 交叉编译 - Windows AMD64
build-windows: deps ## 交叉编译 Windows AMD64 版本
	@echo "正在编译 Windows AMD64 版本..."
	@mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_FILE)

# 交叉编译 - Windows ARM64
build-windows-arm64: deps ## 交叉编译 Windows ARM64 版本
	@echo "正在编译 Windows ARM64 版本..."
	@mkdir -p $(BUILD_DIR)
	GOOS=windows GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-arm64.exe $(MAIN_FILE)

# 交叉编译 - macOS AMD64 (Intel)
build-macos: deps ## 交叉编译 macOS AMD64 版本 (Intel Mac)
	@echo "正在编译 macOS AMD64 版本 (Intel Mac)..."
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-macos-amd64 $(MAIN_FILE)

# 交叉编译 - macOS ARM64 (Apple Silicon)
build-macos-arm64: deps ## 交叉编译 macOS ARM64 版本 (Apple Silicon)
	@echo "正在编译 macOS ARM64 版本 (Apple Silicon)..."
	@mkdir -p $(BUILD_DIR)
	GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-macos-arm64 $(MAIN_FILE)

# 编译所有平台
build-all: build-linux build-linux-arm64 build-windows build-windows-arm64 build-macos build-macos-arm64 ## 编译所有平台版本
	@echo "所有平台编译完成:"
	@ls -la $(BUILD_DIR)/

# 编译常用平台 (AMD64)
build-common: build-linux build-windows build-macos ## 编译常用平台 (AMD64)
	@echo "常用平台 (AMD64) 编译完成:"
	@ls -la $(BUILD_DIR)/

# 编译 ARM64 平台
build-arm64: build-linux-arm64 build-windows-arm64 build-macos-arm64 ## 编译所有 ARM64 版本
	@echo "ARM64 平台编译完成:"
	@ls -la $(BUILD_DIR)/

# 编译 Linux 所有架构
build-linux-all: build-linux build-linux-arm64 ## 编译 Linux 所有架构
	@echo "Linux 所有架构编译完成:"
	@ls -la $(BUILD_DIR)/*linux*

# 编译 Windows 所有架构
build-windows-all: build-windows build-windows-arm64 ## 编译 Windows 所有架构
	@echo "Windows 所有架构编译完成:"
	@ls -la $(BUILD_DIR)/*windows*

# 编译 macOS 所有架构
build-macos-all: build-macos build-macos-arm64 ## 编译 macOS 所有架构
	@echo "macOS 所有架构编译完成:"
	@ls -la $(BUILD_DIR)/*macos*

# 开发模式 - 监听文件变化并重新运行 (需要安装 air)
dev: ## 开发模式 (需要先安装 air: go install github.com/cosmtrek/air@latest)
	@which air > /dev/null || (echo "请先安装 air: go install github.com/cosmtrek/air@latest" && exit 1)
	air

# 生成性能分析
profile: ## 生成性能分析报告
	@echo "生成性能分析报告..."
	go build -o $(BUILD_DIR)/$(BINARY_NAME)-profile $(MAIN_FILE)
	$(BUILD_DIR)/$(BINARY_NAME)-profile 西直门 4 2>/dev/null >/dev/null

# 显示项目信息
info: ## 显示项目信息
	@echo "项目信息:"
	@echo "  名称: 北京地铁查询工具 (Go 版本)"
	@echo "  版本: $(shell git describe --tags --always 2>/dev/null || echo 'dev')"
	@echo "  Go 版本: $(shell go version)"
	@echo "  构建目录: $(BUILD_DIR)"
	@echo "  可执行文件: $(BINARY_NAME)"
