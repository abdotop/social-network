#!/bin/bash

# Create the custom network if it doesn't exist
docker network inspect mynetwork >/dev/null 2>&1 || docker network create mynetwork

# Build the backend service
docker build -t backend-image ./backend

# Build the frontend service
docker build -t frontend-image ./client

# Create and start the backend container
docker run -d --name backend-container -p 8081:8081 --network mynetwork backend-image

# Create and start the frontend container
docker run -d --name frontend-container -p 3000:3000 --network mynetwork -e BACKEND_URL=http://backend-container:8081 frontend-image
