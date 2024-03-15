FROM golang:1.22 

WORKDIR /api_gateway/

COPY . .

RUN CGO_ENABLED=1 go build -o api_gateway

WORKDIR /api_gateway

