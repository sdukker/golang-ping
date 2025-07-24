FROM golang:1.22.5-alpine3.20

RUN mkdir -p /usr/src/app
RUN mkdir -p /usr/src/build

COPY ./run.sh /usr/src/build
RUN chmod +x /usr/src/build/run.sh

WORKDIR /usr/src/app
CMD ["/usr/src/build/run.sh"]

FROM golang:1.22.5-alpine3.20 AS build
RUN mkdir -p /usr/src/app
RUN mkdir -p /usr/src/build

COPY . /usr/src/app
WORKDIR /usr/src/app

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o /usr/src/build/build

FROM alpine
COPY --from=build /usr/src/build/build /

CMD ["/usr/src/build/build"]
