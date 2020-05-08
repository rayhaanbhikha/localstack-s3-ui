FROM golang:alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o localstack-s3-server main.go server.go handlers.go

WORKDIR /dist
RUN cp -r /app/localstack-s3-server /app/build .

RUN ls -l

FROM scratch 

COPY --from=builder /dist /

ENTRYPOINT [ "/localstack-s3-server" ]