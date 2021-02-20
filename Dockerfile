FROM golang:1.12 as base
WORKDIR /src
ENV GO111MODULE=on
COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

FROM base as build-env
COPY . .
RUN CGO_ENABLED=0 go build -o reitapp

FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build-env /src/config.yml /app/config.yml
COPY --from=build-env /src/reitapp /app/
CMD /app/reitapp