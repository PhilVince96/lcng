FROM golang:1.13-alpine
WORKDIR /lcng

RUN apk update && apk add --no-cache gcc musl-dev git bash

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app

CMD ["/lcng/bin/app"]
EXPOSE 8080
