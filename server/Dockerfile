FROM golang:1.20-alpine3.18 AS builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o /app/build/server .

FROM alpine:latest

COPY --from=builder /app/build/server /app/build/server

ENV GIN_MODE=release
ENV PORT=8080
ENV JWT_SECRET=secret
ENV JWT_ISSUER=secret
ENV DB_USER=postgres
ENV DB_PASS=admin
ENV DB_HOST=db
ENV DB_NAME=typing
ENV ORIGIN=http://localhost

EXPOSE 8080

ENTRYPOINT [ "/app/build/server" ]