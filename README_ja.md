[English](README.md) | [日本語](README_ja.md)

# url-anchor

[![Build Status](https://travis-ci.com/ebc-2in2crc/url-anchor.svg?branch=master)](https://travis-ci.com/ebc-2in2crc/url-anchor)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![Version](https://img.shields.io/github/release/ebc-2in2crc/url-anchor.svg?label=version)](https://img.shields.io/github/release/ebc-2in2crc/url-anchor.svg?label=version)

url-anchor は URL を HTML `<a>` タグに変換するプログラムです。

## Description

url-anchor は URL を与えると URL とページのタイトルを HTML の `<a>` タグに変換します。

オプションで Markdown のリンクと reStructuredText のリンクに変換することもできます。

## Usage

```
$ url-anchor https://google.com
<a href="https://google.com">Google</a>

$ url-anchor -m https://google.com
[Google](https://google.com)

$ url-anchor -r https://google.com
`Google <https://google.com>`_

$ url-anchor -help
# ...
```

## Installation

### Developer

```
$ go get -u github.com/ebc-2in2crc/url-anchor/...
```

### User

次の URL からダウンロードします。

- [https://github.com/ebc-2in2crc/url-anchor/releases](https://github.com/ebc-2in2crc/url-anchor/releases)

Homebrew を使うこともできます (Mac のみ)

```sh
$ brew tap ebc-2in2crc/tap
$ brew install url-anchor
```

## Contribution

1. Fork this repository
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Rebase your local changes against the master branch
5. Run test suite with the `make test` command and confirm that it passes
6. Run `make fmt` and `make lint`
7. Create new Pull Request

## Licence

[MIT](https://github.com/ebc-2in2crc/url-anchor/blob/master/LICENSE)

## Author

[ebc-2in2crc](https://github.com/ebc-2in2crc)
