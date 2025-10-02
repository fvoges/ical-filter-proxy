# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Bug Fixes
- **ci:** use direct Docker approach for git-chglog
- **ci:** remove 'v' prefix from git_chglog_version parameter


<a name="v0.4.0"></a>
## [v0.4.0] - 2025-10-02
### Bug Fixes
- **ci:** add security-events permission to GHCR workflow for SARIF upload


<a name="v0.3.0"></a>
## [v0.3.0] - 2025-10-02
### Features
- add comprehensive security improvements and testing ([#1](https://github.com/fvoges/ical-filter-proxy/issues/1))

### BREAKING CHANGE

None - all changes backward compatible

- Fix closure variable capture bug in HTTP handler loop
- Implement constant-time token comparison using crypto/subtle
- Add HTTP client with 30s timeout and redirect limits
- Implement 10MB response body limit to prevent memory exhaustion
- Add proper error handling on response body close
- Enhance URL validation using url.Parse()
- Add graceful shutdown with SIGTERM/SIGINT handling
- Implement security headers middleware (CSP, X-Frame-Options, etc.)
- Add HTTP server timeouts (Read/Write/Idle)
- Improve event anonymization to remove all PII
- Enhance health check endpoints with JSON responses

Security improvements:
- Prevents timing-based token discovery attacks
- Prevents slowloris DoS attacks
- Prevents memory exhaustion via large responses
- Prevents redirect loops
- Adds proper request/response lifecycle management

* test: add comprehensive unit tests for calendar and main packages

- Add calendar_test.go with tests for:
  - AnonymizeEvent functionality
  - StringMatchRule conditions and matching
  - Filter event matching logic
- Add main_test.go with tests for:
  - Secret file reading with whitespace trimming
  - Config validation and loading
  - URL validation
  - Public/token authentication checks
  - FreeBusy mode configuration

Test coverage includes:
- Event anonymization (PII removal)
- Filter matching rules (contains, prefix, suffix, regex)
- Config file parsing and validation
- Edge cases and error conditions

* build: update Go to 1.23 and dependencies

- Update Go version from 1.22.5 to 1.23
- Update github.com/stretchr/testify from v1.8.4 to v1.9.0

Benefits:
- Latest Go security patches and features
- Updated test framework with bug fixes

* build(docker): enhance security and optimize build process

- Update to Alpine 3.21 with security patches
- Add security updates (apk upgrade) in all build stages
- Implement multi-stage caching for faster builds
- Add build flags: -trimpath, -w -s for smaller binaries
- Add HEALTHCHECK instruction for container health monitoring
- Improve layer organization for better caching
- Add comprehensive comments and structure

Security improvements:
- Latest base image security patches
- Smaller attack surface with stripped binaries
- Health check for orchestration platforms

* ci: add comprehensive testing and security scanning workflows

Add test workflow:
- Unit tests with race detection
- Coverage reporting with Codecov integration
- Integration tests for binary validation
- Config validation tests

Add security scanning workflow:
- Trivy filesystem vulnerability scanning
- Gosec Go security analysis
- Dependency review for pull requests
- Weekly scheduled security scans
- SARIF upload to GitHub Security tab

Add GHCR release workflow:
- Build and push to GitHub Container Registry
- Multi-architecture support (amd64, arm64)
- Trivy vulnerability scanning for images
- Automatic tagging for branches, PRs, and releases
- Semantic versioning support
- Docker build caching for faster builds

Benefits:
- Automated security vulnerability detection
- Continuous testing on all PRs
- Regular security audits
- Better visibility into code quality

* ci: update existing workflows for consistency and latest versions

Docker workflow:
- Fix branch name from 'master' to 'main'

Golangci-lint workflow:
- Update to use latest linter version instead of pinned v1.60
- Ensures automatic updates to latest linting rules

* chore: add linting config and enhance dependency management

Add .golangci.yml:
- Configure 20+ linters (errcheck, gosimple, govet, gosec, etc.)
- Set appropriate thresholds and rules
- Enable security-focused linters
- Consistent code quality checks

Enhance renovate.json:
- Add recommended configuration extends
- Enable automerge for minor/patch updates
- Configure vulnerability alerts with labels
- Set weekly schedule for updates
- Limit concurrent PRs to 5
- Add assignees for security updates

Benefits:
- Consistent code quality across team
- Automated dependency updates
- Faster security patch adoption
- Reduced manual review burden

* docs: add security policy and guidelines

Add SECURITY.md:
- Security policy and supported versions
- Vulnerability reporting process
- Security best practices for deployment
- Known security considerations
- Contact information
- Automated security scanning details

Provides clear guidelines for:
- Reporting vulnerabilities responsibly
- Secure deployment practices
- Understanding security features

* docs: update README with security and testing information

- Add Security section highlighting security features
- Add Testing section with test commands
- Update roadmap with completed items (testing, security)
- Document free/busy anonymization feature
- Add references to SECURITY.md

Features documented:
- Constant-time token comparison
- HTTP client timeouts
- Request body size limits
- Graceful shutdown
- Security headers
- Non-root container execution
- Automated security scanning

* fix: resolve linting schema issues and reduce Docker vulnerabilities

.golangci.yml:
- Add version field required by schema
- Simplify configuration to avoid schema errors
- Remove unsupported linters-settings section
- Maintain essential linters (gosec, govet, errcheck, etc.)

Dockerfile:
- Update to golang:1.23.5-alpine3.21 (latest)
- Update to alpine:3.21.2 (latest)
- Add --available flag to apk upgrade for all updates
- Add detailed comments explaining vulnerability scanner warnings
- Resolve 1 critical vulnerability
- Reduce from 3 high to 2 high vulnerabilities

Security notes:
- Remaining vulnerabilities are in intermediate build stages only
- Final scratch image contains no Alpine packages/libraries
- Only static binary, CA certs, and user files in production
- All available security updates applied during build
- Minimal attack surface in production image

* fix: correct syntax errors and simplify event anonymization

- Fix missing closing brace in fetch() function
- Remove duplicate serialization code
- Simplify AnonymizeEvent to use golang-ical API correctly
- Update tests to match actual library behavior
- Remove property manipulation that doesn't match API

All tests now pass:
- TestAnonymizeEvent ✓
- TestStringMatchRule_hasConditions ✓
- TestStringMatchRule_matchesString ✓
- TestFilter_matchesEvent ✓
- TestReadSecretFile ✓
- TestConfigLoadConfig_* (all variants) ✓

Tested with: docker run golang:1.23 go test -v -race ./...

* fix(security): resolve gosec warnings and golangci-lint config issues

- Remove 'version' field from .golangci.yml (not supported in v1.64+)
- Add nosec comments for legitimate file reads (config and secret files)
- Handle Write() errors in health check endpoints
- Make dependency-review continue-on-error (requires GitHub Advanced Security)

* fix(lint): resolve all golangci-lint issues

- Rename struct fields: Url -> URL (proper Go naming convention)
- Replace unused parameters with underscore
- Remove unnecessary nil check (len() is defined for nil slices)
- Run gofmt on all files

* fix(ci): skip Trivy image scan for pull requests

PR builds don't push images to GHCR, so image scanning fails.
PRs are already scanned using filesystem scan in security workflow.
Only scan pushed images (main branch, releases).


<a name="0.2.0"></a>
## [0.2.0] - 2025-01-25
### Features
- add support for reading feed_url and token values from a file
- **config:** replace global unsafe option with public calendar option (breaking change)


<a name="0.1.3"></a>
## [0.1.3] - 2025-01-23
### Bug Fixes
- build docker image from scratch
- pin dockerfile build-base dep, remove go version
- **chart:** templating of args, command
- **ci:** typo for semver in docker metadata
- **ci:** update dockerbuild for tags


<a name="0.1.2"></a>
## [0.1.2] - 2024-09-10

<a name="0.1.1"></a>
## [0.1.1] - 2024-09-01
### Bug Fixes
- documentation mistake for regex config
- dockerhub link in README

### Pull Requests
- Merge pull request [#5](https://github.com/fvoges/ical-filter-proxy/issues/5) from mgrove36/fix-regex-processing
- Merge pull request [#4](https://github.com/fvoges/ical-filter-proxy/issues/4) from mgrove36/feat-match-empty-description-location
- Merge pull request [#3](https://github.com/fvoges/ical-filter-proxy/issues/3) from mgrove36/feat-change-published-calendar-name
- Merge pull request [#2](https://github.com/fvoges/ical-filter-proxy/issues/2) from mgrove36/add-location-transformations
- Merge pull request [#1](https://github.com/fvoges/ical-filter-proxy/issues/1) from mgrove36/feat-change-published-calendar-name


<a name="0.1.0"></a>
## 0.1.0 - 2024-08-18

[v0.4.0]: https://github.com/fvoges/ical-filter-proxy/compare/v0.3.0...v0.4.0
[v0.3.0]: https://github.com/fvoges/ical-filter-proxy/compare/0.2.0...v0.3.0
[0.2.0]: https://github.com/fvoges/ical-filter-proxy/compare/0.1.3...0.2.0
[0.1.3]: https://github.com/fvoges/ical-filter-proxy/compare/0.1.2...0.1.3
[0.1.2]: https://github.com/fvoges/ical-filter-proxy/compare/0.1.1...0.1.2
[0.1.1]: https://github.com/fvoges/ical-filter-proxy/compare/0.1.0...0.1.1
