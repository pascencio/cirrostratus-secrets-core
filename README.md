# Cirrstratus Secrets Core - Password and keys management on multicloud

## Requirements

- GoMock: https://github.com/golang/mock


## Testing

```shell
mockgen -source=ports/repository.go -destination=mock/repository.go -package=mock
mockgen -source=ports/services.go -destination=mock/services.go -package=mock
```