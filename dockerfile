# docker build -t mailgo_hw1 .
FROM golang:1.19
COPY . .
RUN go test -v