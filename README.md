# Get Moe

[![GoDoc](https://godoc.org/github.com/leonidboykov/getmoe?status.svg)](https://godoc.org/github.com/leonidboykov/getmoe)
[![Build Status](https://travis-ci.org/leonidboykov/getmoe.svg?branch=master)](https://travis-ci.org/leonidboykov/getmoe)
[![Go Report Card](https://goreportcard.com/badge/github.com/leonidboykov/getmoe)](https://goreportcard.com/report/github.com/leonidboykov/getmoe)

Get Moe &ndash; is a REST client for image boards, such as Moebooru and
Danbooru. The goal of the project is to provide APIs for the most well-known
image boards (boorus). This project started for the purpose of researching of
various characters popularity, rather than image grabbing, however save feature
is also available.

## Usage

The only implemented command for now is `get`. Here is the usage example.

    getmoe get --tags "tag1 tag2 rating:s" --from booru_name --to save/directory --as {image|json} -l login -p password

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

getmoe is free software licensed under the [MIT](LICENSE) license.
