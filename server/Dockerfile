ARG BUILDER_IMAGE=golang:alpine

#####################
# BUILD environment #
#####################
FROM ${BUILDER_IMAGE} as builder

RUN apk update && apk add --no-cache ca-certificates tzdata && update-ca-certificates

# Create a user
ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735
# for more details
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR /src/

COPY go.mod .
COPY go.sum .

RUN go mod download

# this copies everything from the current directory
# to the src folder defined as WORKDIR above
COPY . ./

# Build the app
RUN GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' -a \
    -o bin/sampleserver .

###########################################
# Create a new Image that will become our #
# RUN environment                         #
###########################################
# you can change to the alpine base image if you
# need to debug and get a shell
# Example: docker run -it --rm samplemicroservice --entrypoint bin/sh 
#FROM alpine
FROM scratch

WORKDIR /app/

###
# Copy our static executable from
# the build image to this image
###
COPY --from=builder /src/bin/sampleserver /app/sampleserver

###
# Note: If your app needs anything other than what is built
# into your executable, you will need to add it here.
###
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# EXAMPLE: You need to copy any config files required at run time into the same directory
# COPY --from=builder /src/sample.yaml .
COPY --from=builder /src/templates/ /app/templates/
COPY --from=builder /src/resources/ /app/resources/

# Use an unprivileged user
USER appuser:appuser

# By default we expose port 8080,
# but this can be changed through environment variables
EXPOSE 8080

ENTRYPOINT ["/app/sampleserver"]