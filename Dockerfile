FROM golang

WORKDIR $GOPATH/src/github.com/axiomzen/beanstalks-api
ADD . .
RUN curl https://glide.sh/get | sh
RUN glide install
RUN go build -o /bin/beanstalks-api
ENTRYPOINT [ "beanstalks-api" ]