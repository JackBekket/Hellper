FROM golang:1.20-alpine AS build_base

# Set the Current Working Directory inside the container
WORKDIR /app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /out/bot .

FROM alpine:latest 

WORKDIR /app

RUN apk add ca-certificates
COPY --from=build_base /out/bot ./bot

CMD ["/app/bot"]
