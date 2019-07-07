FROM golang:1.8

RUN mkdir -p /go/src/AvitoPro
WORKDIR /go/src/AvitoPro

COPY . /go/src/AvitoPro

RUN go-wrapper download
RUN go-wrapper install
CMD ["go-wrapper", "run", "-web"]
EXPOSE 80