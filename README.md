# go_test


## godoc

To use godoc you must install to your system:
```
go install golang.org/x/tools/cmd/godoc@latest
```

To use godoc on localhost use these commands
```
godoc -http :8000
http://localhost:8000/pkg/gotest/
```

## Testing

To test all subdirectories run the command:
```
go test ./pkg/...
```

Or to test just the packages with name "_test" at the end:
```
go test ./*_test
```

To get the coverage of all packages while testing:
```
go test -cover ./pkg/...
```

To get detailed log while testing use -v flag after command like:
```
go test ./pkg/... -v
```

## Code coverage

Get coverage HTML report by:
```
go test -coverprofile=coverage.out ./pkg/...
```

Then use that file with:
```
go tool cover -html=coverage.out
```