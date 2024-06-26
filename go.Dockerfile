FROM golang:1.22

WORKDIR /app 
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Live Hot Reloading using Air (https://github.com/air-verse/air)
# binary will be $(go env GOPATH)/bin/air
RUN curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

CMD ["air"]
