# Get Moe

[![GoDoc](https://godoc.org/github.com/leonidboykov/getmoe?status.svg)](https://godoc.org/github.com/leonidboykov/getmoe)
[![Build Status](https://travis-ci.org/leonidboykov/getmoe.svg?branch=master)](https://travis-ci.org/leonidboykov/getmoe)
[![codecov](https://codecov.io/gh/leonidboykov/getmoe/branch/master/graph/badge.svg)](https://codecov.io/gh/leonidboykov/getmoe)
[![Go Report Card](https://goreportcard.com/badge/github.com/leonidboykov/getmoe)](https://goreportcard.com/report/github.com/leonidboykov/getmoe)
[![Dependabot Status](https://api.dependabot.com/badges/status?host=github&repo=leonidboykov/getmoe)](https://dependabot.com)

Get Moe &ndash; is a REST client for image boards, such as Moebooru and
Danbooru. The goal of the project is to provide APIs for the most well-known
image boards (boorus). This project started for the purpose of researching of
various characters popularity, rather than image grabbing, however save feature
is also available.

## Usage

The only implemented command for now is `get`. Here is the usage example.

```
USAGE:
   getmoe [global options] command [command options] [arguments...]

COMMANDS:
     get      get data from booru
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --quiet, -q    disable progress bar
   --help, -h     show help
   --version, -v  print the version
```

Login and password are optional.

## Supported Boards

  * yande.re
  * konachan.com
  * gelbooru.com
  * danbooru.donmai.us
  * chan.sankakucomplex.com
  * idol.sankakucomplex.com

Custom boorus are not available yet.

## License

getmoe is a free software licensed under the [MIT](LICENSE) license.
