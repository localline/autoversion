############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/localline/autoversion/
# Fetch dependencies.
COPY go.mod .
COPY go.sum .
RUN go mod download
# Copy source code
COPY . .
# Build the binary.
RUN go build -o /go/bin/autoversion
############################
# STEP 2 build a small image
############################
FROM alpine
# Add bash
RUN apk update && apk add bash && apk add curl
# Set tmp as workdir
WORKDIR /tmp
# Copy our static executable.
COPY --from=builder /go/bin/autoversion /go/bin/autoversion
# Copies your code file from your action repository to the filesystem path `/` of the container
COPY entrypoint.sh /entrypoint.sh
# Executes `entrypoint.sh` when the Docker container starts up
ENTRYPOINT ["/entrypoint.sh"]