FROM golang:1.8

WORKDIR $GOPATH/src/awesomeProject
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["awesomeProject"]