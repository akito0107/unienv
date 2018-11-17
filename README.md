# unienv

generate .env from `.ini`

## Getting Started

### Prerequisites
- Go 1.11+

### Installing
```
$ go get -u github.com/akito0107/unienv/cmd/unienv
```

### How to use
1. Declare `.ini` format file as a below; You *must* declare `default` section.

```ini
[default]
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=unienv

[develop_local]
DB_NAME=unienv_develop
DB_PASSWORD=xxxxxxxx

[develop_docker]
DB_HOST=db
```

2. generate `.env` file.
```sh
$ unienv -env develop_local
```
You can get merged `.env` file.

```
DB_NAME=unienv_develop
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=xxxxxxxx
```

## Why?
This tool aims to improve developing with mono-repo 
by simplify distribution and upgrading of env variables.

## Options
```sh
% bin/unienv -h
NAME:
   unienv - unify environment variables

USAGE:
   generr [subcommand] [OPTIONS]

VERSION:
   0.0.1

COMMANDS:
     unify    unify env
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## License
This project is licensed under the Apache License 2.0 License - see the [LICENSE](LICENSE) file for details
