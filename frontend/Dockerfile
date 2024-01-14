# Use the official Node.js image as the base image
FROM node:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the necessary files into the container
COPY . /app

# Install dependencies
RUN npm install

# Build the frontend application
RUN npm run build

# Expose the port that the application will run on
EXPOSE 3000

# Command to run the frontend application
CMD ["npm", "start"]
