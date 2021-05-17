## Will use alpine linux for base image
FROM alpine:3.7 AS builder
MAINTAINER Marcelo Pinheiro <mpinheir@gmail.com>

## Uses alpine package manager to install go, git and g++ packages - note sqllite needs g++
RUN apk add go git g++

## Creates an /app directory within the image to hold application source files
RUN mkdir /app

## Copies everything in the root directory into the /app directory
ADD . /app

## Specifies base /app directory
WORKDIR /app

## Downloads dependencies
RUN go get -d -v

## Builds go app
## GCO_ENABLED=1 is due to sqllite
RUN CGO_ENABLED=1 go build -o quotes

## Deletes unecessary packages
RUN apk del go git g++

# Build the second stage (small) image
FROM alpine
WORKDIR /app
COPY --from=builder /app/quotes .
COPY --from=builder /app/quotesqlite .

## Defines start command which kicks off newly created binary executable
CMD ["/app/quotes"]
