#!/bin/bash
set -e

echo "🧹 Stopping and removing containers..."
docker rm -f simplechatrabbitmq simplechatapi simplechatbotapi 2>/dev/null || true

echo "🧹 Removing images..."
docker rmi -f simplechat simplechatbot simplechatrabbitmq 2>/dev/null || true

echo "✅ Cleanup done. Current Docker status:"
docker ps -a
docker images
