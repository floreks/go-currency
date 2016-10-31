FROM golang:latest

# Prepare environment
RUN mkdir -p go/src/github.com/floreks/go-currency
ADD . go/src/github.com/floreks/go-currency
WORKDIR go/src/github.com/floreks/go-currency

# Download godeps and restore dependencies
RUN go get github.com/tools/godep
RUN go/bin/godep restore

# Go build binary
RUN go build -o go/bin/go-currency .

# Start command for our binary
CMD ["/go/bin/go-currency"]

# Expose port for a container
EXPOSE 8080
