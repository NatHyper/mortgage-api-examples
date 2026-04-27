# Mortgage API Examples

Working examples for the Mortgage Financial Modelling Tool API.

## Prerequisites

- Python 3.x (for Python examples)
- Node.js (for JavaScript examples)
- Go 1.21+ (for Go examples)
- curl (for command line examples)

## Getting Started

1. Get your API key by running any example
2. Use the key in subsequent requests
3. See full API documentation at [https://cleargage.co.uk/api-docs.html](https://cleargage.co.uk/api-docs.html)

## Examples

| Language | File | Run command |
|----------|------|-------------|
| curl | `curl/register.bat` | `register.bat` |
| Python | `python/example.py` | `python example.py` |
| JavaScript | `javascript/example.js` | `node example.js` |
| Go | `go/main.go` | `go run main.go` |

## Notes

- All examples call the production API at `https://cleargage.co.uk/api`
- Each example registers a new API key automatically. Be advised that registration allows 5 registration attempts per hour per IP address. If you see a rate limit error, wait an hour or use a different network.
- Save your API key for future requests.