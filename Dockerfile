# Use a specific GoLang version as specified by the build argument
FROM golang:1.21.3-bookworm


# Set your working directory
WORKDIR /

# Copy your GoLang application source code into the container
COPY . .

RUN apt-get update && apt-get install -y build-essential pkg-config git

# Build your GoLang application
RUN go build main.go

# Specify the command to run when the container starts
CMD ["./app"]
