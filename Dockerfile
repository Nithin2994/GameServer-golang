FROM golang:1.13

RUN wget -O /usr/local/bin/dumb-init https://github.com/Yelp/dumb-init/releases/download/v1.2.2/dumb-init_1.2.2_amd64
RUN chmod +x /usr/local/bin/dumb-init

RUN mkdir app

ADD . /app

WORKDIR /app

RUN go build -o main *.go

CMD ["/bin/bash", "-c", "./main"]