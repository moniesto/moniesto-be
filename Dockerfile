FROM golang:1.18.4-alpine

LABEL maintainer="ParvinEyvazov"

# Install git (to get dependencies)
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

ENV PROJECTNAME=moniesto
ENV PROJECTPATH=/app/${PROJECTNAME}

RUN mkdir -p ${PROJECTPATH}/
WORKDIR ${PROJECTPATH}
COPY go.mod go.sum ${PROJECTPATH}/
RUN go mod download
ADD . ${PROJECTPATH}/

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Build the Go app
RUN go build -o /build -v cmd/main.go

# Copy the source from the current directory to the working Directory inside the container
COPY . .
COPY .env .

# # Build the Go app
# RUN go build -o /build

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD [ "/build" ]
