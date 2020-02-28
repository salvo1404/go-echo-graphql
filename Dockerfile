FROM golang:1.11.1-stretch

COPY . /go/src/app/
WORKDIR /go/src/app/

# Golang package
RUN go get -u \
    github.com/golang/dep/cmd/dep \
    github.com/tockins/realize \
    github.com/pressly/goose/cmd/goose

RUN dep ensure

# bashrc
RUN echo "alias ll='ls -la'" >> ~/.bashrc


EXPOSE 1323
