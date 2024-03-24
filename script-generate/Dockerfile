# use official Golang image as base image
FROM golang:1.16.3-alpine3.13

# Set the Current Working Directory inside the container
WORKDIR /script-generate

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Download and install any required third party dependencies into the container
RUN go get -d -v ./...

# Build the Go app
RUN go build -o script-generate .

RUN chmod +x ./script-generate

# Run the executable
CMD [ "./script-generate" ]