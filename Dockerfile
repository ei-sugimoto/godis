FROM golang:latest AS base
WORKDIR /app



FROM base AS dev
RUN go install github.com/air-verse/air@latest

COPY . .
CMD ["air", "-c", ".air.toml"]

FROM base AS build

RUN --mount=source=go.mod,target=go.mod \
    --mount=source=go.sum,target=go.sum \
    go mod download

RUN --mount=source=.,target=. \
    go build -o /go/bin/main .


FROM gcr.io/distroless/cc:latest AS prod

COPY --from=build /go/bin/main /go/bin/main

CMD ["/go/bin/main", "run"]