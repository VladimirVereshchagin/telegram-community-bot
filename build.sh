#!/bin/bash

# Сборка всех компонентов
echo "Сборка analytics..."
go build -o bin/analytics cmd/analytics/main.go

echo "Сборка automation..."
go build -o bin/automation cmd/automation/main.go

echo "Сборка moderation..."
go build -o bin/moderation cmd/moderation/main.go

echo "Сборка user_management..."
go build -o bin/user_management cmd/user_management/main.go

echo "Сборка завершена!"