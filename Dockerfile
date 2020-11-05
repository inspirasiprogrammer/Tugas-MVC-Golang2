FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/inspirasiprogrammer/Tugas-MVC-Golang2

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["go-sample-app"]

