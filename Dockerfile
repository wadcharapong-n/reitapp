FROM golang:1.10 as build-dev

ADD . /src
RUN cd /src && go build -o reitapp

FROM alpine

WORKDIR /app
COPY --from=build-env /src/reitapp /app/
CMD /app/reitapp