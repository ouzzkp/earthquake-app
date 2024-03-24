FROM golang:1.22

WORKDIR /script-input

COPY . .

RUN go get -d -v ./...

RUN go build -o script-input .

RUN chmod +x ./script-input

CMD [ "./script-input" ]