#!/bin/bash

# Build the backend Docker image
docker build -t backend-image -f Dockerfile.backend .

# Build the frontend Docker image
docker build -t frontend-image -f Dockerfile.frontend .

# Run the backend container
docker run -d -p 8081:8081 backend-image

# Run the frontend container
docker run -d -p 3000:3000 frontend-image

