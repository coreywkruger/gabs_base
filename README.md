# Gabs Number Base Test

See https://github.com/Jeffail/gabs for the lib used.

## Output
The test creates and parses some json with two fields. `smaller` has a smaller integer while `bigger` has a larger integer (difference of x10 between them).
```go
jsonParsed, _ := gabs.ParseJSON([]byte(`{
    "smaller":100000,
    "bigger":1000000
}`))
```
But when the contents are printed, the larger integer is converted to e+ format:
```go
log.Println("=>", jsonParsed.String())
```
... produces `=> {"bigger":1e+06,"smaller":100000}` (notice the `bigger` field)
