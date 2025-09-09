#!/bin/bash
set -e  # si falla un comando, se detiene el script

echo "🚀 Setting up RabbitMQ..."
docker rm -f rabbitmq 2>/dev/null || true
docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management

echo "🐹 Building Chat API..."
docker build -t simplechat -f Dockerfile .

echo "🐹 Running Chat API..."
docker rm -f simplechatapi 2>/dev/null || true
docker run -d --name simplechatapi -p 8080:8080 simplechat

echo "🤖 Building Bot API..."
docker build -t simplechatbot -f BotAPIDockerFile .

echo "🤖 Running Bot API..."
docker rm -f simplechatbotapi 2>/dev/null || true
docker run -d --name simplechatbotapi -p 8081:8081 simplechatbot

echo "✅ All services are up and running!"
echo "   - RabbitMQ: http://localhost:15672 (user: guest / pass: guest)"
echo "   - Chat API: http://localhost:8080"
echo "   - Bot API : http://localhost:8081"
