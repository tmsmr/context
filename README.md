# context
*Minimal HTTP service that dumps the request context as JSON response*

## Get `context`
- `$ go install github.com/tmsmr/context@latest`, or
- `$ docker pull ghcr.io/tmsmr/context:latest`

## Use `context`
- `$ context [-e]`, or
- `$ docker run --rm -p 8080:8080 ghcr.io/tmsmr/context:latest [-e]`

*With the `-e` flag set, the response also contains environment variables*

## Example
```bash
curl "localhost:8080/a/b?a=1&b=2" | jq
{
  "host": "localhost:8080",
  "headers": {
    "Accept": [
      "*/*"
    ],
    "User-Agent": [
      "curl/8.4.0"
    ]
  },
  "path": "/a/b",
  "query": {
    "a": "1",
    "b": "2"
  },
  "env": null
}
```