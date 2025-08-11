FROM golang:tip-alpine3.22 AS build

WORKDIR /app

COPY service/go.mod service/go.sum ./
RUN go mod download

COPY service/. .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o /service .


FROM gcr.io/distroless/static:nonroot

COPY --from=build /service /service
USER nonroot:nonroot

EXPOSE 8080

ENTRYPOINT ["/service"]

