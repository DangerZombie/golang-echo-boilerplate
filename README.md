# Golang Echo Boilerplate
This is the boilerplate of Golang using Echo Framework

## 1. Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### 1.1. Prerequisites

These are the prerequisite library to run the project:

- [Go](https://golang.org/doc/install)
- [mockgen](https://github.com/uber-go/mock)
<!-- - [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [GNU Make](https://www.gnu.org/software/make/)
- [oapi-codegen](https://github.com/deepmap/oapi-codegen)
- [wire](https://github.com/google/wire)
- [migrate](https://github.com/golang-migrate/migrate)
- [gcloud CLI](https://cloud.google.com/sdk/docs/install)
- [gentool](https://gorm.io/gen/gen_tool.html) -->

### 1.2. Installing

A step-by-step series of examples that tell you how to get a development env running

1. Clone the repository

```bash
git clone
```

2. Intialize the project

```bash
make init
```

3. Open the cloned repository. We recommend you to use Visual Studio Code because it's free and light but powerfull. [Visual Studio Code](https://code.visualstudio.com/download)

4. Create lauch.json, to make you eaiser to run and debug the program. This step if you use Visual Studio Code.

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch file",
            "type": "go",
            "request": "launch",
            "env": {
                "localhost": "9000" // use your prefer port
            },
            "mode": "debug",
            "program": "main.go"
        }
    ]
}
```

5. Ask the Backend team on the latest `config-*.yaml` for running the project. Please note this config will not be uploaded to the repository.

## 2. Development

### 2.1. Working With Tests

This section describes how to work with tests in this repository.

### 2.1.1. Writing the tests

#### 2.1.1.1. Writing the unit tests

For every `.go` source code, write a corresponding `_test.go` file in the same directory. For example, if you have a `foo.go` file in `service/foo`, then you should have a `foo_test.go` file in the same directory. Just a quick reminder that unit tests do not have any dependencies to external services such as database, cache, etc. We use `gomock` to mock the dependencies.

#### 2.1.1.2. Writing the integration tests

This is similar with the unit tests. For every `.go` source code, write a corresponding `_integration_test.go` file in the same directory. For example, if you have a `foo.go` file in `repository/foo`, then you should have a `foo_integration_test.go` file in the same directory.

The test function has to use prefix `TestIntegration_` so that we can run the integration tests separately from the unit tests. After the prefix, you can name the test function according to the test case or function that you want to test. Right below the function name, you have to put the following line:

```go
if testing.Short() {
    t.Skip("skipping integration test")
}
```

This is to skip the integration test when we run the unit tests.

### 2.1.2. Writing the integration tests

Run the following to run the tests

```bash
make generate_mocks
```

Excute this to run the unit test

```bash
make test_unit
```

Execute this to run the integration test
```bash
make test_integration
```

### 2.2. Working Directory

This project also have many folder inside, then to minimize the complexity the directory will as follow:

```
project/
├── .github/
│   ├── workflow/
│   └── ...
├── endpoint/
│   └── ...
├── helper/
│   ├── auth/
│   ├── database/
│   ├── http_helper/
│   ├── message/
│   ├── static/
│   ├── util/
│   └── ...
├── http/
│   └── ...
├── initialization/
│   └── ...
├── model/
│   ├── base/
│   ├── entity/
│   ├── request/
│   ├── response/
│   └── ...
├── repository/
│   ├── repository_user/
│   └── ...
├── service/
│   ├── service_user/
│   └── ...
```

Explanations:

1. `.github` folder contains whole files related github.
2. `endpoint` folder contains whole files for related the endpoint of the API.
3. `helper` folder contains whole files for help code easier such as auth logic, conversion logic, etc.
4. `http` folder contains files conversion result after processing and convert according with the response.
5. `model` folder contains model entity represent of the database column.
6. `repository` folder contains files to querying the data to database.
7. `service` folder contains files with business logic or other logic.