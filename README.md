# Cirrstratus Secrets Core - Password and Keys Management on Multicloud

- [Cirrstratus Secrets Core - Password and Keys Management on Multicloud](#cirrstratus-secrets-core---password-and-keys-management-on-multicloud)
  - [Requirements](#requirements)
  - [Setup](#setup)
  - [Testing](#testing)
  - [TODO](#todo)

## Requirements

For developers. To work in this project you will need the following tools:

- GoMock: https://github.com/golang/mock
- Go Task: https://taskfile.dev/#/installation?id=get-the-binary

## Setup

The easy setup require Go Task tool. Once this tool it's installed, you can run:

```shell
task
```

Then... ta da!! All it's ready to go!

## Testing

First of all, you will need to generate some mock files for `service` and `repository` ports.

You can use the following commands:

```shell
mockgen -source=ports/repository.go -destination=mock/repository.go -package=mock
mockgen -source=ports/service.go -destination=mock/service.go -package=mock
```

Using `task`, all mocks are generated automatically:

```shell
task test
```

> To see all tasks included, execute this command `task -l`.

## TODO

- [ ] Encrypt logic
- [ ] Decrypt logic
- [ ] Integration tests with Test Containers