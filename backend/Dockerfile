FROM golang:1.22

WORKDIR /backend-app

COPY . .

RUN go get -d -v ./...

RUN go build -o backend-app .

EXPOSE 8080

CMD [ "./backend-app" ]