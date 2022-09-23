FROM golang:1.18-alpine as build
WORKDIR /build

COPY server/routes/ server/routes/
COPY server/db/ server/db/
COPY server/migrations/ server/migrations/
COPY server/util/ server/util/
COPY server/main.go server/main.go
COPY go.sum .
COPY go.mod .

RUN go mod download
RUN go build -o ./bin/shawty-server ./server/main.go


FROM alpine:3 AS run
WORKDIR /app

COPY --from=build /build/bin .

RUN apk add ca-certificates

EXPOSE 8080
ENTRYPOINT ["/app/shawty-server"]