# Use the official Node.js image as a base
# FROM node:latest

# FROM node:20.7.0-alpine
FROM node:latest

# Set the working directory inside the container
WORKDIR /app
# Copy package.json and package-lock.json to install dependencies
# COPY package*.json ./
COPY . .
# Install dependencies
# RUN yarn install
RUN npm install
# RUN npm install execa
# Copy the rest of the application source code
# COPY ./ .
# Build the Nuxt application
RUN npm run build
# Expose port 3000 to the outside world
EXPOSE 3000
# Command to run the application
CMD ["npm", "start"]

