FROM golang:1.19.1-alpine as build
RUN apk add make
RUN mkdir -p /app/src
WORKDIR /app/src

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY ./data_collection_service ./data_collection_service
COPY ./pkg ./pkg
RUN go build -o main ./data_collection_service/cmd/data_collection_service

FROM alpine:3.7
COPY --from=build /app/src/main /root/main

WORKDIR /root/

CMD ["./main"]
