# use official Golang image as base image
FROM node:17-alpine

# Set the Current Working Directory inside the container
WORKDIR /frontend

# Copy the source code from the current directory to the Working Directory inside the container
COPY package.json .

RUN npm install

COPY . .

EXPOSE 3000

CMD ["npm", "start"]