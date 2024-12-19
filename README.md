<a name="readme-top"></a>

# <p align="center">FIAP-TechChallenge-1</p>

<p align="center">
    <img src="https://img.shields.io/badge/Code-Go-informational?style=flat-square&logo=go&color=00ADD8" alt="Go" />
    <img src="https://img.shields.io/badge/Tools-Docker-informational?style=flat-square&logo=docker&color=2496ED" alt="Docker" />
</p>

## 💬 About

Repository for the FIAP Tech Challenge, focused on developing a backend system for managing orders in a fast-food restaurant.

### :open_file_folder: Project Structure

```
.
├── bin
├── cmd
│   └── http
├── docs
└── internal
    ├── adapter
    │   ├── cache
    │   │   └── redis
    │   ├── handler
    │   │   └── http
    │   ├── repository
    │   │   └── postgres
    │   │       └── migrations
    │   └── token
    │       └── paseto
    └── core
        ├── domain
        ├── port
        ├── service
        └── util

```

- `bin`: directory to store compiled executable binary.
- `docs`: directory to store project's documentation, such as swagger static files.
- `cmd`: directory for main entry points or commands of the application. The http sub-directory holds the main HTTP server entry point.
- `internal`: directory for containing application code that should not exposed to external packages.
- `core`: directory that contains the central business logic of the application. Inside it there are 4 sub-directories.
- `domain`: directory that contains domain models/entities representing core business concepts.
- `port`: directory that contains defined interfaces or contracts that adapters must follow.
- `service`: directory that contains the business logic or services of the application.
- `util`: directory that contains utility functions that reused in the service package.
- `adapters`: directory for containing external services that will interact with the core of application. There are 4 external services used in this application.
- `handler/http`: directory that contains HTTP request and response handler.
- `repository/postgres`: directory that contains database adapters for PostgreSQL.
- `cache/redis`: directory that contains cache adapters for Redis.
- `token/paseto`: directory that contains token generation and validation adapters using Paseto.

### :pushpin: Decisions

TBD


### :pushpin: Features
- [x] Dockerfile: small image with multi-stage docker build, and independent of the host environment
- [x] Makefile: to simplify the build and run commands
- [x] Hexagonal architecture
- [x] PostgreSQL database
- [x] Conventional commits
- [x] Unit tests
- [x] Code coverage
- [x] Swagger documentation
- [x] Feature branch workflow
- [x] Air to run go


## :computer: Technologies

- [Go](https://golang.org/)
- [Docker](https://www.docker.com/)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## :scroll: Requirements

### Build/Run with Docker

- [Docker](https://www.docker.com/)

### Build/Run Locally

- [Go](https://golang.org/)

> [!NOTE]
> You need to have Go (> 1.18) installed in your machine to build, run and test the application locally

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## :cd: Installation

```sh
git clone git@github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1.git
```

```sh
cd FIAP-TechChallenge-Fase1
```

```sh
cp .env.example .env
```

### :whale: Docker

```sh
make compose-build
```

### :hammer: Build (build locally)

```sh
make build
```
> The binary will be created in the `bin` folder

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## :runner: Running

### :whale: Docker

```sh
make run-compose
```

### :hammer: Build (run locally)

```sh
make run
```

> [!NOTE]
> `make run` will run the application locally, and will build and run PostgreSQL container using Docker Compose

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## :white_check_mark: Tests

```sh
make test
```
> It will run the unit tests and generate the coverage report as `cp.out` and `coverage.html`

## :clap: Acknowledgments

- [Hexagonal Architecture, Ports and Adapters in Go](https://medium.com/@kyodo-tech/hexagonal-architecture-ports-and-adapters-in-go-f1af950726b)
- [Building RESTful API with Hexagonal Architecture in Go](https://dev.to/bagashiz/building-restful-api-with-hexagonal-architecture-in-go-1mij)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

