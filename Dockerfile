FROM golang:1.12 as build-env

ENV GO111MODULE=on
ADD . /src
RUN cd /src && go build -o reitapp

FROM alpine

WORKDIR /app
COPY --from=build-env /src/reitapp /app/
CMD /app/reitapp