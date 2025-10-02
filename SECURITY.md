# Security Policy

## Supported Versions

We release patches for security vulnerabilities. Which versions are eligible for receiving such patches depends on the CVSS v3.0 Rating:

| Version | Supported          |
| ------- | ------------------ |
| Latest  | :white_check_mark: |
| < Latest| :x:                |

## Reporting a Vulnerability

Please report (suspected) security vulnerabilities to the repository owner. You will receive a response within 48 hours. If the issue is confirmed, we will release a patch as soon as possible depending on complexity.

**Please do not report security vulnerabilities through public GitHub issues.**

### What to Include

When reporting a vulnerability, please include:

- Type of issue (e.g. buffer overflow, SQL injection, cross-site scripting, etc.)
- Full paths of source file(s) related to the manifestation of the issue
- The location of the affected source code (tag/branch/commit or direct URL)
- Any special configuration required to reproduce the issue
- Step-by-step instructions to reproduce the issue
- Proof-of-concept or exploit code (if possible)
- Impact of the issue, including how an attacker might exploit it

### Preferred Languages

We prefer all communications to be in English.

## Security Update Policy

- Security updates are released as soon as possible after a vulnerability is confirmed
- Users are notified via GitHub Security Advisories
- Critical vulnerabilities are addressed with emergency releases
- All security patches are documented in the CHANGELOG

## Security Best Practices for Users

When deploying this application:

1. **Always use HTTPS** - Never expose the service over unencrypted HTTP in production
2. **Strong tokens** - Use cryptographically secure random tokens (32+ characters)
3. **Least privilege** - Run the container as a non-root user (already configured)
4. **Network isolation** - Deploy behind a reverse proxy with rate limiting
5. **Regular updates** - Keep the application and dependencies up to date
6. **Monitor logs** - Watch for unauthorized access attempts
7. **Secrets management** - Use Docker secrets or environment variables for sensitive data
8. **Resource limits** - Set appropriate memory and CPU limits for the container

## Known Security Considerations

### Token-based Authentication

This application uses token-based authentication via URL query parameters. While convenient, this has some security implications:

- Tokens may appear in server logs
- Tokens may be cached in browser history
- Consider using reverse proxy with header-based authentication for additional security

### Upstream Feed Security

The application fetches iCal feeds from external sources. Ensure:

- Feed URLs are from trusted sources
- Network egress is properly controlled
- Consider using a proxy for external requests

## Automated Security Scanning

This repository includes:

- Trivy vulnerability scanning for Docker images and dependencies
- Gosec for Go code security analysis
- Dependabot/Renovate for automated dependency updates
- CodeQL analysis (if enabled)

## Contact

For security-related questions or concerns, please contact the repository maintainers.
