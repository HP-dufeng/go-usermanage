# golang image where workspace (GOPATH) configured at /go.
FROM golang

# Install bcrypt package
RUN mkdir -p $GOPATH/src/golang.org/x
RUN cd $GOPATH/src/golang.org/x
RUN git clone https://github.com/golang/crypto.git

# Copy the local package files to the container's workspace
ADD . /go/src/github.com/dufeng/usermanager

# Setting up working directory
WORKDIR /go/src/github.com/dufeng/usermanager

# Get godeps for managing and restoring dependenices
RUN go get github.com/tools/godep

# Restore godep dependencies
RUN godep restore

# Build the usermanager command inside the container
RUN go install github.com/dufeng/usermanager

# Run the usermanager command when the container starts
# ENTRYPOINT /go/bin/usermanager

# Service listens on port 8080
EXPOSE 8080
