# Multi Stage container
# Note : All latest images are used for simplicity

####
# STAGE 1 : BUILD APPLICATION BINARY
FROM golang:latest as builder
WORKDIR /lifesaver

# CERTIFIACTES : COPY FROM STAGE 1
# To run this container from the GoLang container
# COPY --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# DEPENDENCIES
# Copy go mod and sum files (To avoid duplicate requests for dependencies)
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# COPY AND BUILD APPLICATION
# Copy the source from the current directory to the Working Directory inside the container
COPY . .
# Build the Go Application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=readonly -o ./lifesaver ./cmd/lifesaver/



####
# STAGE 1 : ROOT CERTIFICATES
FROM alpine:latest as final

# CERTIFICATES
RUN mkdir -p /etc/ssl/certs/
RUN apk update && apk add -U --no-cache ca-certificates

# APPLICATION USER
RUN addgroup -g 1001 app
RUN adduser app -u 1001 -D -G app /home/app

# Copy Root Certificates | Set Owner and Group to 1001
RUN chown -R 1001:1001 /etc/ssl/certs/

# Copy binary that is built in Stage 2 | Set owner to group and user 1001
COPY --chown=1001:1001 --from=builder /lifesaver/lifesaver /lifesaver

RUN ls
# RUN cd lifesaver

# Container User : App
USER app

EXPOSE 8081

ENV PATH="/usr/local/go/bin:${PATH}"

# Run Library (that is built from previous commands)
ENTRYPOINT ["./lifesaver"]


















# # Multi Stage container
# # Note : All latest images are used for simplicity

# ####
# # STAGE 1 : ROOT CERTIFICATES
# FROM alpine:latest as root-certs

# # CERTIFICATES
# RUN mkdir -p /etc/ssl/certs/ && apk update && apk add -U --no-cache ca-certificates

# # APPLICATION USER
# RUN addgroup -g 1001 app
# RUN adduser app -u 1001 -D -G app /home/app


# ####
# # STAGE 2 : BUILD APPLICATION BINARY
# FROM golang:latest as builder
# WORKDIR /lifesaver

# # CERTIFIACTES : COPY FROM STAGE 1
# # To run this container from the GoLang container
# COPY --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# # DEPENDENCIES
# # Copy go mod and sum files (To avoid duplicate requests for dependencies)
# COPY go.mod go.sum ./
# # Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
# RUN go mod download

# # COPY AND BUILD APPLICATION
# # Copy the source from the current directory to the Working Directory inside the container
# COPY . .
# # Build the Go Application
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=readonly -o ./lifesaver ./cmd/lifesaver/


# ####
# # STAGE 3 : BUILD FROM SCRATCH
# FROM scratch as final

# # Copy Groups and Users
# COPY --from=root-certs /etc/passwd /etc/passwd
# COPY --from=root-certs /etc/group /etc/group

# # Copy Root Certificates | Set Owner and Group to 1001
# COPY --chown=1001:1001 --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# # Copy binary that is built in Stage 2 | Set owner to group and user 1001
# COPY --chown=1001:1001 --from=builder /lifesaver/lifesaver /lifesaver

# # Container User : App
# USER app

# # ENVIRONMENT PATH
# ENV PATH="/usr/local/go/bin:${PATH}"

# # Expose Port
# EXPOSE 8081

# # Run Library (that is built from previous commands)
# ENTRYPOINT ["./lifesaver"]