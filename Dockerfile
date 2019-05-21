FROM golang

WORKDIR $GOPATH/src/github.com/axiomzen/beanstalks-api
ADD . .
RUN go build -o /bin/beanstalks-api
ENTRYPOINT [ "beanstalks-api" ]