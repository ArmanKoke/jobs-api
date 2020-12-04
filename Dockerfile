FROM golang:1.15.5-alpine3.12 as builder

RUN apk update && \
    apk add curl && \
    apk add make

WORKDIR /app
COPY ./ ./

RUN make build

EXPOSE 8080 5432
ENTRYPOINT ["./main"]
