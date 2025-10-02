# Release Process

This document describes how to create a new release of ical-filter-proxy.

## Prerequisites

- [ ] All tests passing on `main` branch
- [ ] All CI/CD workflows passing
- [ ] CHANGELOG.md updated
- [ ] Security scans completed without critical issues
- [ ] Documentation updated

## Versioning

This project follows [Semantic Versioning](https://semver.org/):

- **MAJOR** version (x.0.0) - Incompatible API changes
- **MINOR** version (0.x.0) - New functionality (backward compatible)
- **PATCH** version (0.0.x) - Bug fixes (backward compatible)

### Version Guidelines

Based on the changes in the current PR, this would be a **MINOR** version bump (e.g., `v0.1.0` → `v0.2.0`) because:
- ✅ New features added (free/busy mode, security improvements)
- ✅ No breaking changes
- ✅ Multiple enhancements and bug fixes

## Release Steps

### 1. Prepare the Release

```bash
# Ensure you're on main and up to date
git checkout main
git pull origin main

# Verify all tests pass
go test -v -race ./...

# Run linting
golangci-lint run

# Build locally to verify
go build -v .
```

### 2. Update CHANGELOG.md

Move changes from `[Unreleased]` section to a new version section:

```markdown
## [0.2.0] - 2025-10-02

### Added
- (list of additions)

### Changed
- (list of changes)

### Fixed
- (list of fixes)

### Security
- (security improvements)

## [Unreleased]
(empty for now)
```

### 3. Commit CHANGELOG

```bash
git add CHANGELOG.md
git commit -m "docs: update CHANGELOG for v0.2.0 release"
git push origin main
```

### 4. Create and Push Git Tag

```bash
# Create annotated tag
git tag -a v0.2.0 -m "Release v0.2.0

Major security improvements and testing enhancements.

See CHANGELOG.md for full details."

# Push tag to trigger release
git push origin v0.2.0
```

### 5. Automated Release Process

Once the tag is pushed, GitHub Actions will automatically:

1. **GoReleaser Workflow** triggers:
   - Builds binaries for multiple platforms (Linux, macOS, Windows)
   - Creates GitHub Release with changelog
   - Uploads release artifacts
   - Generates checksums

2. **Docker Workflows** trigger:
   - Builds multi-arch Docker images
   - Pushes to Docker Hub with version tags
   - Pushes to GitHub Container Registry
   - Runs security scans

### 6. Verify Release

After workflows complete:

- [ ] Check GitHub Releases page: https://github.com/fvoges/ical-filter-proxy/releases
- [ ] Verify binaries are uploaded
- [ ] Test Docker image: `docker pull ghcr.io/fvoges/ical-filter-proxy:v0.2.0`
- [ ] Review changelog in release notes
- [ ] Check security scan results

### 7. Announce Release

Consider announcing the release:
- [ ] Update README if needed
- [ ] Notify users of security improvements
- [ ] Update documentation site (if applicable)
- [ ] Post in relevant communities

## Docker Image Tags

After release, the following tags will be available:

- `ghcr.io/fvoges/ical-filter-proxy:latest` - Latest stable release
- `ghcr.io/fvoges/ical-filter-proxy:v0.2.0` - Specific version
- `ghcr.io/fvoges/ical-filter-proxy:v0.2` - Minor version
- `ghcr.io/fvoges/ical-filter-proxy:v0` - Major version
- `ghcr.io/fvoges/ical-filter-proxy:main` - Latest main branch build

## Rollback Process

If issues are discovered after release:

### Quick Rollback

```bash
# Delete the tag locally
git tag -d v0.2.0

# Delete the tag remotely
git push origin :refs/tags/v0.2.0

# Delete the release on GitHub (via UI or CLI)
gh release delete v0.2.0
```

### Fix and Re-release

1. Fix the issue on main
2. Create a new patch release (e.g., v0.2.1)
3. Follow the release process again

## Release Checklist Template

Copy this checklist for each release:

```markdown
## Release vX.Y.Z Checklist

### Pre-release
- [ ] All tests passing
- [ ] All CI/CD workflows green
- [ ] CHANGELOG.md updated
- [ ] Security scans reviewed
- [ ] Documentation updated
- [ ] Version number determined

### Release
- [ ] On main branch
- [ ] CHANGELOG committed
- [ ] Tag created: `git tag -a vX.Y.Z -m "Release vX.Y.Z"`
- [ ] Tag pushed: `git push origin vX.Y.Z`

### Post-release
- [ ] GitHub Release created (automatic)
- [ ] Binaries uploaded (automatic)
- [ ] Docker images built (automatic)
- [ ] Release verified and tested
- [ ] Announcement sent (if applicable)

### If Issues Found
- [ ] Issues documented
- [ ] Rollback performed (if needed)
- [ ] Patch release planned
```

## Troubleshooting

### GoReleaser Fails

```bash
# Test GoReleaser locally without publishing
goreleaser release --snapshot --clean

# Check configuration
goreleaser check
```

### Docker Build Fails

```bash
# Build locally
docker build -t ical-filter-proxy:test .

# Check Dockerfile syntax
docker build --check -t ical-filter-proxy:test .
```

### Tag Already Exists

```bash
# Force update tag (use with caution)
git tag -fa v0.2.0 -m "Release v0.2.0"
git push origin v0.2.0 --force
```

## Additional Notes

- Always test the release process in a fork first
- Keep security patches confidential until released
- Major version bumps may require migration guide
- Document breaking changes clearly

## Questions?

For questions about the release process, see:
- [CONTRIBUTING.md](./CONTRIBUTING.md) (if exists)
- Open an issue on GitHub
- Contact maintainers
