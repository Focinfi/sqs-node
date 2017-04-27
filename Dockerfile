FROM daocloud.io/golang:1.8.1-onbuild
RUN go get github.com/Focinfi/sqs-node/
RUN CGO_ENABLED=0 go install -a github.com/Focinfi/sqs-node/
WORKDIR $GOPATH/src/github.com/Focinfi/sqs-node/
EXPOSE 8080
CMD ["sqs-node"]
