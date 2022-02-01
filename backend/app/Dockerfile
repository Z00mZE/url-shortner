FROM golang:alpine3.15 as builder
RUN apk add --update-cache && apk add --no-cache git ca-certificates
WORKDIR /src/app
COPY . .
RUN go mod download && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o service ./main.go

FROM scratch
COPY --from=builder /src/app/assets /assets
COPY --from=builder /src/app/template /template
COPY --from=builder /src/app/service /service
ENTRYPOINT ["/service"]
