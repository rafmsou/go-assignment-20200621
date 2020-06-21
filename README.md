
## How to run the server
```
go run agentero_server/main.go -ams-api-url=http://localhost:4000 -schedule-period=5m
```
Parameters:
- ams-api-url: AMS API server address where agents data will be imported from.
- schedule-period: Frequency at which data will be imported. Eg.: 5m (5 minutes).
If not provided, data will be imported only on server startup.

## How to run the sample client
```
go run agentero_client/main.go
```

## How to run tests with coverage report
```
go test -cover ./...
```
