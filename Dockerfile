FROM golang:1.19

WORKDIR /usr/src/botson

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/botson ./...

CMD [ "botson" ]