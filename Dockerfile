FROM golang:alpine

RUN apk update && apk add --no-chache git

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o binary

ENTRYPOINT [ "/app/binary" ]