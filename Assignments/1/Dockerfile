FROM golang:1.20

WORKDIR /app

RUN export GOROOT=/usr/local/go
RUN export GOPATH=$HOME/go
RUN  PATH=$GOPATH/bin:$GOROOT/bin:$PATH

COPY . /app

RUN go build -o main .
CMD ["go", "run", "main.go","mono"]