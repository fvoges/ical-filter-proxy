# Stage 1: build golang binary
# Note: Base image vulnerabilities are scanner-detected CVEs in Alpine packages.
# These do not affect the final scratch-based image as no Alpine packages are included.
# We apply all available security updates and use static compilation.
FROM golang:1.23.5-alpine3.21 AS builder
ARG VERSION="unknown"
WORKDIR /go/src/app

# Install build dependencies and apply all security updates
RUN apk update && apk upgrade --available --no-cache && apk add --no-cache ca-certificates

# Copy dependency files first for better caching
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Build with security flags
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s -extldflags '-static' -X 'main.version=${VERSION}'" \
    -trimpath \
    -o /go/bin/ical-filter-proxy

# Stage 2: setup alpine base for building scratch image
# Note: Intermediate stage vulnerabilities do not affect final image security.
# Final image is scratch-based with only: binary, CA certs, and user files.
# No Alpine packages or libraries are included in production image.
FROM alpine:3.21.2 AS base
RUN apk update && apk upgrade --available --no-cache && apk add --no-cache ca-certificates && \
    adduser -s /bin/true -u 1000 -D -h /app app && \
    sed -i -r "/^(app|root)/!d" /etc/group /etc/passwd && \
    sed -i -r 's#^(.*):[^:]*$#\1:/sbin/nologin#' /etc/passwd

# Stage 3: create final image from scratch
FROM scratch
WORKDIR /app

# Copy certificates and user info
COPY --from=base /etc/passwd /etc/group /etc/shadow /etc/
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy binary
COPY --from=builder /go/bin/ical-filter-proxy /usr/bin/ical-filter-proxy

# Use non-root user
USER app

# Expose port
EXPOSE 8080/tcp

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD ["/usr/bin/ical-filter-proxy", "-version"]

# Run the application
ENTRYPOINT ["/usr/bin/ical-filter-proxy"]
CMD ["-config", "/app/config.yaml"]
