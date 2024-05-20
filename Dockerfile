FROM golang:1.22-alpine AS build_base

# Set the Current Working Directory inside the container
WORKDIR /app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod go.sum ./
RUN go mod download

COPY . .



#RUN go build -o /out/bot .
RUN go build -o main .

FROM alpine:latest 

WORKDIR /app

RUN apk add ca-certificates
#COPY --from=build_base /out/bot ./bot
COPY --from=build_base /app/main /app/main

#CMD ["/app/bot"]

EXPOSE 8085

CMD [ "./main" ]
