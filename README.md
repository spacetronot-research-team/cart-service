# cart-service

## Project Structure
```
├── cmd
│   ├──  application
│   │   ├──  app.go
├── deployment
├── external
├── internal
│   ├──  api
│   │   ├──  v{version}
│   │   │   ├──  {domain}
│   │   │   │   ├──  controller
│   │   │   │   ├──  repository
│   │   │   │   ├──  service
│   ├──  config
│   ├──  infrastructure
│   ├──  middleware
│   ├──  pkg
│   ├──  router
├── pkg
└── main.go
```

- `/cmd`: Initial stage of the application will run.
- `/deployment`: IaaS, PaaS, system and container orchestration deployment configurations and templates (docker-compose, kubernetes/helm, mesos, terraform, bosh).
- `/external`: This module is the place to use third-party libraries or APIs.
    - Is your code contain logic to interact with external? Put here.
- `/internal`: This module is the core module of the application and contains the implementation of various business logic.
- `/internal/api/v{version}/{domain}/controller`: This module is only to gather input (REST/gRPC/etc) and pass request to service.
    - Is your code contain logic to get input? Put here.
    - Is your code contain business logic? Do not put here.
    - Is your code contain saving data to database? Do not put here.
- `/internal/api/v{version}/{domain}/repository`: This module is only contain logic related to persistence (CRUD database/redis/etc).
    - Is your code contain logic to get input? Do not put here.
    - Is your code contain business logic? Do not put here.
    - Is your code contain saving data to database? Put here.
- `/internal/api/v{version}/{domain}/service`: This module contain business logic, this module get input from controller and use repository for things related to persistence.
    - Is your code contain logic to get input? Do not put here.
    - Is your code contain business logic? Put here.
    - Is your code contain saving data to database? Do not put here.
- `/internal/config`: This module contains the config.go file which serves to provide config variables for various environments, production or development.
- `/internal/infrastructure`: This module is to define the infrastructure that will be used in the application (database, redis, etc).
- `/internal/middleware`: This module contains middleware used in the application (auth, error handling, etc.).
- `/internal/pkg`: This module contains library code to use by only this project.
    - Is your code contain package that can be used by other project? Do not put here.
- `/internal/router`: This module contains a router file where endpoints are defined as well as connecting the middleware and controller.
- `/pkg`: This module contains library code that is ok to use by external applications. Other projects will import these libraries expecting them to work, so think twice before you put something here.
    - Is your code contain package that can be used by other project? Put here.

## How to Run

1. Rename `.env-example` to `.env`.

```shell
mv .env-example .env
```

2. Run database using docker compose in folder `deployment`.

```shell
cd deployment && sudo docker compose up
```

3. Run `main.go` in other terminal.

```shell
go run main.go
```
