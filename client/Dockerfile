# Build stage server
FROM golang:1.18.3-alpine as build

RUN apk update

RUN mkdir /app
WORKDIR /app

COPY . /app

RUN cd /app && go build -o /app/server .


# Run stage
FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY --from=build /app /app

EXPOSE 3000
CMD [ "/app/server" ]
