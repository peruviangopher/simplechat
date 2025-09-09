#!/bin/bash
set -e

echo "🧹 Stopping and removing containers..."
docker rm -f rabbitmq simplechatapi simplechatbotapi 2>/dev/null || true

echo "🧹 Removing images..."
docker rmi -f simplechat simplechatbot 2>/dev/null || true

echo "✅ Cleanup done. Current Docker status:"
docker ps -a
docker images
