# TOC MACHINE TRADING

## Layers

![Example](docs/img/layers.png)

## Tools

### Conventional Commit

- install git cz tool global

```sh
sudo npm install -g commitizen
sudo npm install -g cz-conventional-changelog
sudo npm install -g conventional-changelog-cli
echo '{ "path": "cz-conventional-changelog" }' > ~/.czrc
```

### Pre-commit

- install git pre-commit tool global

```sh
brew install pre-commit
```

- install/modify from config

```sh
pre-commit autoupdate
pre-commit install
pre-commit run --all-files
```

### Modify CHANGELOG

- First Time

```sh
conventional-changelog -p angular -i CHANGELOG.md -s -r 0
```

- From Last semver tag

```sh
conventional-changelog -p angular -i CHANGELOG.md -s
```

### Find ignored files

```sh
find . -type f  | git check-ignore --stdin
```

### Migrate Tool

- install

```sh
version=v4.15.2
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@$version
```

### env file example

```sh
echo 'DEPLOYMENT=dev
LOG_FORMAT=console
LOG_LEVEL=trace
DISABLE_SWAGGER_HTTP_HANDLER=
GIN_MODE=release
SINOPAC_URL=127.0.0.1:56666
PG_URL=postgres://postgres:asdf0000@127.0.0.1:5432/
RABBITMQ_URL=amqp://admin:password@127.0.0.1:5672/%2f
RABBITMQ_EXCHANGE=toc
DB_NAME=machine_trade
' > .env
```

### VSCode Debug Setting

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Debug",
            "type": "go",
            "request": "attach",
            "debugAdapter": "dlv-dap",
            "processId": "toc-machine-trading",
        }
    ]
}
```

### Config

```sh
cp ./configs/default.config.yml ./configs/config.yml
```

## Authors

- [**Tim Hsu**](https://gitlab.tocraw.com/root)
