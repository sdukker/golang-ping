FROM golang:1.22.5-alpine3.20

RUN mkdir -p /usr/src/app
RUN mkdir -p /usr/src/build

COPY ./run.sh /usr/src/build
RUN chmod +x /usr/src/build/run.sh

WORKDIR /usr/src/app
CMD ["/usr/src/build/run.sh"]