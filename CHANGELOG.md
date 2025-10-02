# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Comprehensive unit tests for calendar and main packages (29.8% coverage)
- Security scanning workflows (Trivy + Gosec)
- Test workflow with race detection and coverage reporting
- GHCR (GitHub Container Registry) release workflow
- Security policy document (SECURITY.md)
- golangci-lint configuration with 20+ linters
- Enhanced renovate.json with automerge for dependencies
- Graceful shutdown with SIGTERM/SIGINT handling
- Security headers middleware (X-Frame-Options, CSP, etc.)
- HTTP server with proper timeouts (Read/Write/Idle)
- Health check endpoints with JSON responses
- Free/busy mode for complete event anonymization
- Docker HEALTHCHECK instruction

### Changed
- Updated Go from 1.22.5 to 1.23
- Updated testify from v1.8.4 to v1.9.0
- Enhanced Dockerfile with security improvements
- Updated to golang:1.23.5-alpine3.21 and alpine:3.21.2
- Improved Docker layer caching
- Updated golangci-lint to use latest version

### Fixed
- **CRITICAL:** Closure variable capture bug in HTTP handler loop
- **SECURITY:** Token comparison now uses constant-time to prevent timing attacks
- **SECURITY:** HTTP client now has 30s timeout to prevent slowloris DoS attacks
- **SECURITY:** Added 10MB response body limit to prevent memory exhaustion
- Enhanced URL validation using url.Parse()
- Improved error handling on response body close
- Fixed branch name in docker workflow (master → main)
- Resolved 1 critical and reduced Docker vulnerabilities

### Security
- Constant-time token comparison using crypto/subtle
- HTTP client with timeout and redirect limits
- Request body size limits (10MB)
- Multiple layers of DoS protection
- Proper resource management and cleanup
- Production-ready security headers
- Non-root container execution
- Automated security scanning

## [Previous Versions]

See Git history for changes prior to structured changelog.
