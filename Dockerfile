FROM golang:1.22-bookworm

WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN apt update && apt install -y ca-certificates
RUN go build -o main .

EXPOSE 8085

CMD [ "./main" ]
