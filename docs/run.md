## To make the app work we need to run the app container with these simple steps:

### 1) Build and start the app container
Details about app services can be found in `docker-compose.yaml` file
```shell
docker-compose up
```
### 2) Check the running containers
```shell
docker container ls
```
### 3) Perform REST API operations

All examples:

```shell
`POST` http://127.0.0.1:8080/records
```
The request body will contain `IntValue`, `StrValue`, `BoolValue` and `TimeValue` fields

```shell
`GET` http://127.0.0.1:8080/1
```

```shell
`PUT` http://127.0.0.1:8080/records/1
```
The request body will contain `IntValue`, `StrValue`, `BoolValue` and `TimeValue` fields

```shell
`DELETE` http://127.0.0.1:8080/1
```

Detailed description of the endpoints can be found in `docs/api/openapi.spec.yml` file

## Run tests (make sure the records.bin file is generated, if not run the command twice to generate it)
```shell
go test ./pkg/model/records -v
```