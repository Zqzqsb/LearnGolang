#!/bin/bash

# 运行所有测试并生成覆盖率报告
echo "Running tests and generating coverage report..."
go test -coverprofile=coverage.out ./...

# 显示覆盖率统计信息
echo "Coverage summary:"
go tool cover -func=coverage.out

# 打开覆盖率的 HTML 报告
echo "Generating HTML coverage report..."
go tool cover -html=coverage.out -o coverage.html

echo "Done. Open 'coverage.html' in your browser to view the detailed coverage report."