# Project Structure

## Golang Clean Architect

---

### Group by types

```
/project
|-- /cmd
|   |-- main.go
|
|-- /internal
|   |-- /adapters
|   |   |-- /kafka
|   |   |   |-- producer.go
|   |   |   |-- consumer.go
|   |   |
|   |   |-- /postgres
|   |       |-- client.go
|   |
|   |-- /shared
|   |   |-- date.go
|   |   |-- utils.go
|   |   
|   |-- /models
|   |   |-- user.go
|   |
|   |-- /controllers
|   |   |-- user_controller.go
|   |
|   |-- /usecases
|   |   |-- user_usecase.go
|   |   |-- user_usecase_test.go
|   |
|   |-- /repositories
|       |-- user_repository.go
|
|-- /tests
|   |-- user_test.go
|
|-- go.mod
|-- go.sum
```

#### cmd

contains the entry point of the project.
can be extended if the project has multiple entry points.

```
|-- /cmd
    |-- /server
    |   |-- main.go
    |-- /cli
        |-- main.go
```

#### adapters

contains the external dependencies or other services.

#### shared

contains the shared code aka helper or utility.

#### models

contains the entities and value objects.

#### controllers

implement the http transport based on framework ex. echo gin.

#### usecases

implement the business logic and the unit test.

#### repositories

implement the data store with sql or orm.

#### tests

implement the integration testing.

---

### Group by features/domains

```
/project
|-- /cmd
|   |-- main.go
|
|-- /internal
|   |-- /adapters
|   |   |-- /kafka
|   |   |   |-- producer.go
|   |   |   |-- consumer.go
|   |   |
|   |   |-- /postgres
|   |       |-- client.go
|   |
|   |-- /shared
|   |   |-- date.go
|   |   |-- utils.go
|   |   
|   |-- /domains
|       |-- /user
|       |   |-- /models
|       |   |   |-- payload.go
|       |   |   |-- user.go
|       |   |
|       |   |-- controller.go
|       |   |-- usecases.go
|       |   |-- usecases_test.go
|       |   |-- repository.go
|       |   |-- service.go (main/initialize)
|       |
|       |-- /order
|           |-- ...
|
|-- /tests
|   |-- user_test.go
|   |-- order_test.go
|
|-- go.mod
|-- go.sum
```

#### domains

move all clean architect related directory into domains instead.

#### models

models can split into more sub directories

```
/models
|-- /entities
|   |-- user.go
|-- /payload (request, response)
    |-- get_user.go
    |-- create_user.go
```

#### hexagonal architecture (port & adapter)

use interfaces to decoupling each layers from each others

```
/domains
|-- /controller
|   |-- http.go
|   |-- cli.go
|
|-- /usecases
|   |-- validation.go
|   |-- validation_test.go
|   |-- usecases.go
|   |-- usecases_test.go
|
|-- /repository
|   |-- repository.go (interface)
|   |-- pgstore.go
|   |-- memstore.go
|
|-- service.go (main/initialize)
```

### Infrastructure

```
/project
|-- /scripts
|   |-- initdb.sql
|   |-- user.sql
|   |-- seed.sql
|
|-- /deployment
|   |-- /docker
|   |-- /k8s
|
|-- /internal
|   |-- /logging
|   |   |-- healthcheck.go
|   |
|   |-- /middleware
|       |-- auth.go
|       |-- ratelimit.go
|       |-- circuitbreaker.go
|
|-- /docs
|   |-- openapi.yaml
|   |-- example.http
|
|-- .env.example
```

#### scripts

contains all scripts related files ex. sql, shell.

#### deployment

containerize the application and automate deployment.

#### logging

contains all logging related ex. health check, log level.

#### middleware

contains middlewares that check the request before access the controllers.

### docs

documentation, mostly only openapi.yaml