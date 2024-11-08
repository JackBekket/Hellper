#FROM golang:1.22-bookworm AS build_base

# Set the Current Working Directory inside the container
#WORKDIR /app

# We want to populate the module cache based on the go.{mod,sum} files.
#COPY go.mod go.sum ./
#RUN go mod download

#COPY . .


FROM golang:1.22-bookworm

WORKDIR /app

COPY ./tmp ./tmp
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN apt update && apt install -y ca-certificates
RUN go build -o main .
#RUN go build -o /out/bot .



#RUN apk add ca-certificates
#COPY --from=build_base /out/bot ./bot
#COPY --from=build_base /app/main /app/main
#COPY --from=build_base /app/.env /app/.env

#CMD ["/app/bot"]

EXPOSE 8085

CMD [ "./main" ]
